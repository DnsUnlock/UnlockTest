package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func J_COM_ON_DEMAND(c http.Client) result.Result {
	c.CheckRedirect = nil
	resp, err := url.GET(c, "https://linkvod.myjcom.jp/auth/login")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 403:
		return result.Result{Status: status.No}
	case 502:
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.OK}
}
