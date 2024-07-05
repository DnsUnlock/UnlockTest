package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func TlcGo(c http.Client) result.Result {
	resp, err := url.GET(c, "https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location/dnsfeed")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	if strings.Contains(string(b), `"country":"US"`) {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
