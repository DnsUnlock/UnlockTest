package dialer

import (
	"net"
	"net/http"
	"time"
)

// Dialer 是 HTTP 客户端使用的默认拨号器。
var Dialer = &net.Dialer{
	Timeout:   30 * time.Second,
	KeepAlive: 30 * time.Second,
	// Resolver:  &net.Resolver{},
}

// UseLastResponse 允许使用上次响应。
func UseLastResponse(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }
