package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func MyVideo(c http.Client) result.Result {
	c.CheckRedirect = nil
	resp, err := url.GET(c, "https://www.myvideo.net.tw/login.do")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	if strings.Contains(string(b), "serviceAreaBlock") {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.OK}
}
