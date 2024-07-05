package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func DirectvStream(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.atttvnow.com/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return result.Result{Status: status.No}
	}
	// if resp.StatusCode == 200 || resp.StatusCode == 301 {
	// 	return result.Result{Status: status.OK}
	// }
	return result.Result{Status: status.OK}
}
