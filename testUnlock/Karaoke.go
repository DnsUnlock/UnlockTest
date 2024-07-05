package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func Karaoke(c http.Client) result.Result {
	resp, err := url.GET(c, "http://cds1.clubdam.com/vhls-cds1/site/xbox/sample_1.mp4.m3u8")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return result.Result{Status: status.OK}
	case 403:
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.Unexpected}
}
