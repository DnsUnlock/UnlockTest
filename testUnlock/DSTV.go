package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func DSTV(c http.Client) result.Result {
	resp, err := url.GET(c, "https://authentication.dstv.com/favicon.ico")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 403 || resp.StatusCode == 451 {
		return result.Result{Status: status.No}
	}

	if resp.StatusCode == 404 {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
