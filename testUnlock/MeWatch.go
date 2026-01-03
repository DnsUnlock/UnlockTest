package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func MeWatch(c http.Client) result.Result {
	resp, err := url.GET(c, "https://cdn.mewatch.sg/api/items/97098/videos?delivery=stream%2Cprogressive&ff=idp%2Cldp%2Crpt%2Ccd&lang=en&resolution=External&segments=all")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}

	if resp.StatusCode == 403 {
		return result.Result{Status: status.No}
	}

	return result.Result{Status: status.Unexpected}
}
