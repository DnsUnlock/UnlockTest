package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/ua"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func Watcha(c http.Client) result.Result {
	resp, err := url.GET(c, "https://watcha.com/",
		url.H{"User-Agent", ua.PCBrowserUA()},
		url.H{"Host", "watcha.com"},
		url.H{"Connection", "keep-alive"},
		url.H{"Upgrade-Insecure-Requests", "1"},
		url.H{"Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 451 {
		return result.Result{Status: status.No}
	}
	if resp.StatusCode == 403 {
		return result.Result{Status: status.Banned}
	}

	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK, Region: "kr"}
	}

	if resp.StatusCode == 302 && resp.Header.Get("Location") == "/ja-JP/" {
		return result.Result{Status: status.OK, Region: "jp"}
	}

	return result.Result{Status: status.Unexpected}
}
