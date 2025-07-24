package result

import (
	"encoding/json"

	"github.com/DnsUnlock/UnlockTest/lib/status"
)

type Result struct {
	Status     int
	Region     string
	Info       string
	Err        error
	StatusText string
}

func (r *Result) SetStatusText() {
	var statusText string
	switch r.Status {
	case status.OK:
		statusText = "✅ 解锁成功"
	case status.No:
		statusText = "❌ 不支持"
	case status.Restricted:
		statusText = "🔒 受限制"
	case status.Banned:
		statusText = "🚫 被封禁"
	case status.NetworkErr:
		statusText = "🌐 网络错误"
	case status.Failed:
		statusText = "💥 测试失败"
	default:
		statusText = "❓ 未知状态"
	}
	r.StatusText = statusText
}

func (r *Result) ToString() string {
	body, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(body)
}
