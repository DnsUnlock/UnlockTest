package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Joyn(c http.Client) result.Result {
	url := "https://auth.joyn.de/auth/anonymous"
	resp, err := urls.PostJson(c, url,
		`{"client_id":"b74b9f27-a994-4c45-b7eb-5b81b1c856e7","client_name":"web","anon_device_id":"b74b9f27-a994-4c45-b7eb-5b81b1c856e7"}`,
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
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(b, &res); err != nil {
		// log.Println(err)
		return result.Result{Status: status.Failed, Err: err}
	}
	url2 := "https://api.joyn.de/content/entitlement-token"
	resp2, err := urls.PostJson(c, url2, `{"content_id":"daserste-de-hd","content_type":"LIVE"}`,
		urls.H{"authorization", "Bearer " + res.AccessToken},
		urls.H{"x-api-key", "36lp1t4wto5uu2i2nk57ywy9on1ns5yg"},
	)
	if err != nil {
		// log.Println(err)
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()
	b2, err := io.ReadAll(resp2.Body)
	if err != nil {
		// log.Println(err)
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res2a []struct {
		Code string `json:"code"`
	}
	var res2b struct {
		Token string `json:"entitlement_token"`
	}
	if strings.Contains(string(b2), "Unauthorized") {
		return result.Result{Status: status.Err, Err: err}
	}
	if err := json.Unmarshal(b2, &res2a); err != nil {
		if err := json.Unmarshal(b2, &res2b); err != nil {
			return result.Result{Status: status.Failed, Err: err}
		}
		if res2b.Token != "" {
			return result.Result{Status: status.OK}
		}
	}
	if res2a[0].Code == "ENT_AssetNotAvailableInCountry" {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.Unexpected}
}
