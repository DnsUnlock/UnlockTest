package client

import (
	"github.com/DnsUnlock/UnlockTest/lib/dialer"
	"github.com/DnsUnlock/UnlockTest/lib/transport"
	"net/http"
	"time"
)

var Auto = NewAuto()

func NewAuto() http.Client {
	return http.Client{
		Timeout:       30 * time.Second,
		CheckRedirect: dialer.UseLastResponse,
		Transport:     transport.Auto,
	}
}
