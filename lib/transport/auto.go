package transport

import (
	"github.com/DnsUnlock/UnlockTest/lib/proxy"
	"github.com/DnsUnlock/UnlockTest/lib/tls"
	"net"
	"net/http"
	"time"
)

var Auto = &http.Transport{
	Proxy:       proxy.Client,
	DialContext: (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
	// ForceAttemptHTTP2:     true,
	MaxIdleConns:           100,
	IdleConnTimeout:        90 * time.Second,
	TLSHandshakeTimeout:    30 * time.Second,
	ExpectContinueTimeout:  1 * time.Second,
	TLSClientConfig:        tls.Config,
	MaxResponseHeaderBytes: 262144,
}
