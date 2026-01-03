package transport

import (
	"context"
	"github.com/DnsUnlock/UnlockTest/lib/dialer"
	"github.com/DnsUnlock/UnlockTest/lib/proxy"
	"github.com/DnsUnlock/UnlockTest/lib/tls"
	"net"
	"net/http"
	"time"
)

var Ipv4 = &http.Transport{
	Proxy: proxy.Client,
	DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		// 强制使用IPv4
		return dialer.Dialer.DialContext(ctx, "tcp4", addr)
	},
	// ForceAttemptHTTP2:     true,
	MaxIdleConns:           100,
	IdleConnTimeout:        90 * time.Second,
	TLSHandshakeTimeout:    30 * time.Second,
	ExpectContinueTimeout:  1 * time.Second,
	TLSClientConfig:        tls.Config,
	MaxResponseHeaderBytes: 262144,
}
