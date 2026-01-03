package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func SBSonDemand(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.sbs.com.au/api/v3/network?context=odwebsite")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res struct {
		Get struct {
			Response struct {
				CountryCode string `json:"country_code"`
			} `json:"response"`
		} `json:"get"`
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}

	if res.Get.Response.CountryCode == "AU" {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
