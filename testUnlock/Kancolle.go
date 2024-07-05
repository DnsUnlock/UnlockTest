package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func Kancolle(c http.Client) result.Result {
	resp, err := url.GETDalvik(c, "http://203.104.209.7/kcscontents/news/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}

	if resp.StatusCode == 403 || resp.StatusCode == 302 {
		return result.Result{Status: status.No}
	}

	return result.Result{Status: status.Unexpected}
}
