package dpanel

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/DnsUnlock/UnlockTest/lib/result"
)

// TestFunc 测试函数类型
type TestFunc func(c http.Client) result.Result

// TestRunner 测试运行器
type TestRunner struct {
	reporter    *Reporter
	tests       map[string]TestFunc
	client      http.Client
	concurrency int
	interval    time.Duration
	stopChan    chan struct{}
	wg          sync.WaitGroup
}

// NewTestRunner 创建测试运行器
func NewTestRunner(reporter *Reporter, client http.Client) *TestRunner {
	return &TestRunner{
		reporter:    reporter,
		tests:       make(map[string]TestFunc),
		client:      client,
		concurrency: 5,
		interval:    30 * time.Minute,
		stopChan:    make(chan struct{}),
	}
}

// SetConcurrency 设置并发数
func (r *TestRunner) SetConcurrency(n int) {
	r.concurrency = n
}

// SetInterval 设置测试间隔
func (r *TestRunner) SetInterval(d time.Duration) {
	r.interval = d
}

// RegisterTest 注册测试
func (r *TestRunner) RegisterTest(name string, fn TestFunc) {
	r.tests[name] = fn
}

// RegisterTests 批量注册测试
func (r *TestRunner) RegisterTests(tests map[string]TestFunc) {
	for name, fn := range tests {
		r.tests[name] = fn
	}
}

// RunOnce 运行一次所有测试
func (r *TestRunner) RunOnce() map[string]result.Result {
	results := make(map[string]result.Result)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// 使用信号量控制并发
	sem := make(chan struct{}, r.concurrency)

	for name, fn := range r.tests {
		wg.Add(1)
		go func(name string, fn TestFunc) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			res := fn(r.client)
			res.SetStatusText()

			mu.Lock()
			results[name] = res
			mu.Unlock()
		}(name, fn)
	}

	wg.Wait()
	return results
}

// RunAndReport 运行测试并上报结果
func (r *TestRunner) RunAndReport() error {
	results := r.RunOnce()
	return r.reporter.ReportMap(results)
}

// Start 启动定时测试
func (r *TestRunner) Start() {
	r.wg.Add(1)
	go r.run()
}

func (r *TestRunner) run() {
	defer r.wg.Done()

	// 立即运行一次
	if err := r.RunAndReport(); err != nil {
		log.Printf("Failed to report unlock results: %v", err)
	}

	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := r.RunAndReport(); err != nil {
				log.Printf("Failed to report unlock results: %v", err)
			}
		case <-r.stopChan:
			return
		}
	}
}

// Stop 停止定时测试
func (r *TestRunner) Stop() {
	close(r.stopChan)
	r.wg.Wait()
}

