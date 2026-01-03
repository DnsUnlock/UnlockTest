package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func CanalPlus(c http.Client) result.Result {
	resp, err := url.GET(c, "https://boutique-tunnel.canalplus.com/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 {
		return result.Result{Status: status.No}
	}

	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}

	if resp.StatusCode == 403 {
		return result.Result{Status: status.Banned}
	}

	return result.Result{Status: status.Unexpected}
}
