package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"strings"
)

func Hotstar(c http.Client) result.Result {
	resp, err := url.GET(c, "https://api.hotstar.com/o/v1/page/1557?offset=0&size=20&tao=0&tas=20")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 475:
		return result.Result{Status: status.No}
	case 401:
		resp, err := url.GET(c, "https://www.hotstar.com")
		if err != nil {
			return result.Result{Status: status.NetworkErr, Err: err}
		}
		if resp.StatusCode == 301 {
			return result.Result{Status: status.No}
		}
		u := resp.Header.Get("Location")
		if u == "" {
			return result.Result{Status: status.No}
		}
		t := strings.SplitN(u, "/", 4)
		if len(t) < 4 {
			return result.Result{Status: status.No}
		}
		return result.Result{Status: status.OK, Region: t[3]}
	}
	return result.Result{Status: status.Unexpected}
}
