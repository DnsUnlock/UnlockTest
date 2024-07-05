package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func VideoLand(c http.Client) result.Result {
	resp, err := url.PostJson(c, "https://api.videoland.com/subscribe/videoland-account/graphql", `{"operationName":"IsOnboardingGeoBlocked","variables":{},"query":"query IsOnboardingGeoBlocked {\n  isOnboardingGeoBlocked\n}\n"}`,
		url.H{"connection", "keep-alive"},
		url.H{"apollographql-client-name", "apollo_accounts_base"},
		url.H{"traceparent", "00-cab2dbd109bf1e003903ec43eb4c067d-623ef8e56174b85a-01"},
		url.H{"origin", "https://www.videoland.com"},
		url.H{"referer", "https://www.videoland.com/"},
		url.H{"accept", "application/json, text/plain, */*"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res struct {
		Data struct {
			Blocked bool `json:"isOnboardingGeoBlocked"`
		} `json:"data"`
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.Data.Blocked {
		return result.Result{Status: status.No}
	}

	return result.Result{Status: status.OK}
}
