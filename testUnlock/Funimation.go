package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func Funimation(c http.Client) result.Result {
	resp, err := urls.GET(c, "https://www.crunchyroll.com/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		return result.Result{Status: status.No}
	}
	for _, c := range resp.Cookies() {
		if c.Name == "region" {
			return result.Result{Status: status.OK, Region: c.Value}
		}
	}
	return result.Result{Status: status.Failed}
}
