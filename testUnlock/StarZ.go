package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func Starz(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.starz.com/sapi/header/v1/starz/us/09b397fc9eb64d5080687fc8a218775b", url.H{"Referer", "https://www.starz.com/us/en/"})
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	authorization := string(b)
	resp2, err := url.GET(c, "https://auth.starz.com/api/v4/User/geolocation", url.H{"AuthTokenAuthorization", authorization})
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()
	b2, err := io.ReadAll(resp2.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res struct {
		IsAllowedAccess  bool
		IsAllowedCountry bool
		IsKnownProxy     bool
		Country          string
	}
	if err := json.Unmarshal(b2, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.IsAllowedAccess && res.IsAllowedCountry && !res.IsKnownProxy {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
