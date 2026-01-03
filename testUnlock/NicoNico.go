package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Niconico(c http.Client) result.Result {
	resp, err := urls.GET(c, "https://www.nicovideo.jp/watch/so40278367")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if strings.Contains(string(b), "同じ地域") {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.OK}
}
