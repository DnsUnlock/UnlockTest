package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func DocPlay(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.docplay.com/subscribe")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 303 {
		return result.Result{Status: status.OK}
	}

	if resp.StatusCode == 307 {
		return result.Result{Status: status.No}
	}

	return result.Result{Status: status.Unexpected}
}
