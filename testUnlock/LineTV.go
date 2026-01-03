package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"regexp"
)

func extractLineTVMainJsUrl(html string) string {
	re := regexp.MustCompile(`href="(https://web-static\.linetv\.tw/release-fargate/public/dist/main-[a-z0-9]{8}-prod\.js)"`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func extractLineTVAppId(js string) string {
	re := regexp.MustCompile(`appId:"([^"]+)"`)
	matches := re.FindStringSubmatch(js)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func LineTV(c http.Client) result.Result {
	resp1, err := url.GET(c, "https://www.linetv.tw/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp1.Body.Close()
	body1, err := io.ReadAll(resp1.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	bodyString1 := string(body1)
	mainJsUrl := extractLineTVMainJsUrl(bodyString1)

	if mainJsUrl == "" {
		return result.Result{Status: status.Failed}
	}
	resp2, err := url.GET(c, mainJsUrl)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()
	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	bodyString2 := string(body2)
	appId := extractLineTVAppId(bodyString2)
	if appId == "" {
		return result.Result{Status: status.Failed}
	}

	resp3, err := url.GET(c, "https://www.linetv.tw/api/part/11829/eps/1/part?appId="+appId+"&productType=FAST&version=10.38.0")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp3.Body.Close()
	body3, err := io.ReadAll(resp3.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res struct {
		CountryCode int `json:"countryCode"`
	}
	if err := json.Unmarshal(body3, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.CountryCode == 228 {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
