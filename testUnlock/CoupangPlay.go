package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func CoupangPlay(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.coupangplay.com/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if resp.StatusCode == 302 && resp.Header.Get("Location") == "https://www.coupangplay.com/not-available" {
		return result.Result{Status: status.No}
	}

	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
