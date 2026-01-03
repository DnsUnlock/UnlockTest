package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func extractSonyLivJwtToken(body string) string {
	re := regexp.MustCompile(`resultObj:"([^"]+)`)
	matches := re.FindStringSubmatch(body)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func SonyLiv(c http.Client) result.Result {
	req, err := http.NewRequest("GET", "https://www.sonyliv.com/", nil)
	resp1, err := url.RunFor(c, req)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp1.Body.Close()

	body1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if strings.Contains(string(body1), "geolocation_notsupported") {
		return result.Result{Status: status.No}
	}

	jwtToken := extractSonyLivJwtToken(string(body1))

	resp2, err := url.GET(c, "https://apiv2.sonyliv.com/AGL/1.4/A/ENG/WEB/ALL/USER/ULD",
		url.H{"accept", "application/json, text/plain, */*"},
		url.H{"referer", "https://www.sonyliv.com/"},
		url.H{"device_id", "25a417c3b5f246a393fadb022adc82d5-1715309762699"},
		url.H{"app_version", "3.5.59"},
		url.H{"security_token", jwtToken},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()

	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res1 struct {
		ResultObj struct {
			CountryCode string `json:"country_code"`
		} `json:"resultObj"`
	}

	if err := json.Unmarshal(body2, &res1); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}

	region := res1.ResultObj.CountryCode

	if region == "" {
		return result.Result{Status: status.Failed}
	}

	resp3, err := url.GET(c, "https://apiv2.sonyliv.com/AGL/3.8/A/ENG/WEB/"+region+"/ALL/CONTENT/VIDEOURL/VOD/1000273613/prefetch",
		url.H{"upgrade-insecure-requests", "1"},
		url.H{"accept", "application/json, text/plain, */*"},
		url.H{"origin", "https://www.sonyliv.com"},
		url.H{"referer", "https://www.sonyliv.com/"},
		url.H{"device_id", "25a417c3b5f246a393fadb022adc82d5-1715309762699"},
		url.H{"security_token", jwtToken},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp3.Body.Close()

	body3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res2 struct {
		ResultCode string `json:"resultCode"`
		Message    string `json:"message"`
	}

	if err := json.Unmarshal(body3, &res2); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}

	if res2.ResultCode == "OK" {
		return result.Result{Status: status.OK, Region: strings.ToLower(region)}
	}

	if res2.ResultCode == "KO" {
		return result.Result{Status: status.No, Region: strings.ToLower(region), Info: "Proxy"}
	}

	return result.Result{Status: status.Unexpected}
}
