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

func SupportStarPlus(loc string) bool {
	var STARPLUS_SUPPORT_COUNTRY = []string{
		"BR", "MX", "AR", "CL", "CO", "PE", "UY", "EC", "PA", "CR", "PY", "BO", "GT", "NI", "DO", "SV", "HN", "VE",
	}
	for _, s := range STARPLUS_SUPPORT_COUNTRY {
		if loc == s {
			return true
		}
	}
	return false
}

func StarPlus(c http.Client) result.Result {

	resp, err := url.GET(c, "https://www.starplus.com/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return result.Result{Status: status.Failed}
	}

	if resp.StatusCode == 403 {
		return result.Result{Status: status.Banned}
	}

	if resp.StatusCode == 302 && (resp.Header.Get("Location") == "https://www.preview.starplus.com/unavailable" || resp.Header.Get("Location") == "https://www.starplus.com/welcome/unavailable") {
		return result.Result{Status: status.No}
	}

	if resp.StatusCode == 200 {
		re := regexp.MustCompile(`Region:\s+([A-Za-z]{2})`)
		matches := re.FindStringSubmatch(string(body))
		if len(matches) >= 2 {
			if SupportStarPlus(matches[1]) {
				return result.Result{Status: status.OK, Region: strings.ToLower(matches[1])}
			}
			return result.Result{Status: status.No}
		}
		return result.Result{Status: status.Unexpected}
	}

	return result.Result{Status: status.Unexpected}
}
