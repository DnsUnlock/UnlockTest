package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func MusicJP(c http.Client) result.Result {
	resp, err := url.GET(c, "https://overseaauth.music-book.jp/globalIpcheck.js")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	if string(b) == "" {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.OK}
}
