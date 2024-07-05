package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func SHOWTIME(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.showtime.com/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return result.Result{Status: status.No}
	}
	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.Failed}
}
