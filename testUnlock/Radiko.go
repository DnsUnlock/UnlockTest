package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Radiko(c http.Client) result.Result {
	resp, err := url.GET(c, "https://radiko.jp/area?_=1625406539531")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, `classs="OUT"`) {
		return result.Result{Status: status.No}
	}
	if strings.Contains(s, "JAPAN") {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
