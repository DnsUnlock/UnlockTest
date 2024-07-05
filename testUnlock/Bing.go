package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func extractBingRegion(responseBody string) string {
	re := regexp.MustCompile(`Region:"([^"]*)"`)
	match := re.FindStringSubmatch(responseBody)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func Bing(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.bing.com/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		return result.Result{Status: status.Failed}
	}

	if resp.StatusCode == 200 {
		region := extractBingRegion(bodyString)
		if region == "CN" {
			return result.Result{Status: status.No, Region: "cn"}
		}
		if region != "" {
			return result.Result{Status: status.OK, Region: strings.ToLower(region)}
		}
	}

	if strings.Contains(bodyString, "cn.bing.com") {
		return result.Result{Status: status.No, Region: "cn"}
	}

	if resp.StatusCode == 403 {
		return result.Result{Status: status.Banned}
	}

	return result.Result{Status: status.Unexpected}
}
