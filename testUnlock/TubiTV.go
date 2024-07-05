package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func TubiTV(c http.Client) result.Result {
	resp, err := url.GET(c, "https://tubitv.com/home")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 {
		resp2, err := url.GET(c, "https://gdpr.tubi.tv")
		if err != nil {
			return result.Result{Status: status.NetworkErr, Err: err}
		}
		defer resp2.Body.Close()
		b, err := io.ReadAll(resp2.Body)
		if err != nil {
			return result.Result{Status: status.NetworkErr, Err: err}
		}
		if strings.Contains(string(b), "Unfortunately") {
			return result.Result{Status: status.No}
		}
		return result.Result{Status: status.OK}
	}
	if resp.StatusCode == 403 {
		return result.Result{Status: status.No}
	}
	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.Unexpected}
}
