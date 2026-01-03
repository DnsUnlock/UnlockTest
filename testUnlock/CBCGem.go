package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func CBCGem(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.cbc.ca/g/stats/js/cbc-stats-top.js")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, `country":"CA"`) {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
