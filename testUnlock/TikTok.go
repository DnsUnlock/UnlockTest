package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func extractTikTokRegion(body string) string {
	re := regexp.MustCompile(`"region":"(\w+)"`)
	matches := re.FindStringSubmatch(body)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func TikTok(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.tiktok.com/explore")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	if err != nil {
		return result.Result{Status: status.Failed}
	}

	if strings.Contains(bodyString, "https://www.tiktok.com/hk/notfound") {
		return result.Result{Status: status.No, Region: "hk"}
	}

	if region := extractTikTokRegion(bodyString); region != "" {

		return result.Result{Status: status.OK, Region: strings.ToLower(region)}
	}

	return result.Result{Status: status.No}
}
