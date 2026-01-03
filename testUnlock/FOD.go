package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func FOD(c http.Client) result.Result {
	resp, err := url.GET(c, "https://geocontrol1.stream.ne.jp/fod-geo/check.xml?time=1624504256")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "true") {
		return result.Result{Status: status.OK}
	}
	if strings.Contains(s, "false") {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.Unexpected}
}
