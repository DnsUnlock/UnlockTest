package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

// Princess Connect Re:Dive Japan
func PCRJP(c http.Client) result.Result {
	resp, err := url.GETDalvik(c, "https://api-priconne-redive.cygames.jp/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 404:
		return result.Result{Status: status.OK}
	case 403:
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.Unexpected}
}
