package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func FXNOW(c http.Client) result.Result {
	resp, err := url.GET(c, "https://fxnow.fxnetworks.com")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	if strings.Contains(string(b), "is not accessible") {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.OK}
}
