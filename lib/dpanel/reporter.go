// Package dpanel 提供与 Dpanel 服务器通信的功能
package dpanel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
)

// UnlockResult 解锁测试结果
type UnlockResult struct {
	NodeID      uint   `json:"node_id"`
	ServiceName string `json:"service_name"`
	Status      string `json:"status"`
	Region      string `json:"region,omitempty"`
	Info        string `json:"info,omitempty"`
	TestedAt    int64  `json:"tested_at"`
}

// Reporter Dpanel 上报器
type Reporter struct {
	apiURL   string
	apiKey   string
	nodeID   uint
	client   *http.Client
}

// NewReporter 创建上报器
func NewReporter(apiURL, apiKey string, nodeID uint) *Reporter {
	return &Reporter{
		apiURL: apiURL,
		apiKey: apiKey,
		nodeID: nodeID,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// ConvertStatus 将状态码转换为字符串
func ConvertStatus(statusCode int) string {
	switch statusCode {
	case status.OK:
		return "ok"
	case status.Restricted:
		return "restricted"
	case status.Banned:
		return "banned"
	case status.NetworkErr, status.Err, status.Failed:
		return "error"
	default:
		return "unknown"
	}
}

// ConvertResult 将测试结果转换为上报格式
func (r *Reporter) ConvertResult(serviceName string, res result.Result) *UnlockResult {
	return &UnlockResult{
		NodeID:      r.nodeID,
		ServiceName: serviceName,
		Status:      ConvertStatus(res.Status),
		Region:      res.Region,
		Info:        res.Info,
		TestedAt:    time.Now().Unix(),
	}
}

// Report 上报单个测试结果
func (r *Reporter) Report(serviceName string, res result.Result) error {
	results := []*UnlockResult{r.ConvertResult(serviceName, res)}
	return r.ReportBatch(results)
}

// ReportBatch 批量上报测试结果
func (r *Reporter) ReportBatch(results []*UnlockResult) error {
	data, err := json.Marshal(results)
	if err != nil {
		return fmt.Errorf("failed to marshal results: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/unlock/report", r.apiURL)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.apiKey)

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// ReportMap 上报测试结果映射
func (r *Reporter) ReportMap(results map[string]result.Result) error {
	unlockResults := make([]*UnlockResult, 0, len(results))
	for serviceName, res := range results {
		unlockResults = append(unlockResults, r.ConvertResult(serviceName, res))
	}
	return r.ReportBatch(unlockResults)
}

