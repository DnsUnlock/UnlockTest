package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func KayoSports(c http.Client) result.Result {
	resp, err := url.GET(c, "https://kayosports.com.au",
		url.H{"Accept", "*/*"},
		url.H{"Accept-Language", "en-US,en;q=0.9"},
		url.H{"Origin", "https://kayosports.com.au"},
		url.H{"Referer", "https://kayosports.com.au/"},
	)
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

	return result.Result{Status: status.Unexpected}
}
