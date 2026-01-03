package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"regexp"
)

func extractSkyShowTimeRegion(url string) string {
	re := regexp.MustCompile(`https://www.skyshowtime.com/([a-z]{2})\?`)

	matches := re.FindStringSubmatch(url)

	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

func SkyShowTime(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.skyshowtime.com/",
		url.H{"Cookie", "sat_track=true; AMCVS_99B971AC61C1E36F0A495FC6@AdobeOrg=1; AMCV_99B971AC61C1E36F0A495FC6@AdobeOrg=179643557|MCIDTS|19874|MCMID|36802229575946481753961418923958457479|MCOPTOUT-1717079521s|NONE|vVersion|5.5.0"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if resp.StatusCode == 307 {
		if resp.Header.Get("Location") == "https://www.skyshowtime.com/where-can-i-stream" {
			return result.Result{Status: status.No}
		}
		region := extractSkyShowTimeRegion(resp.Header.Get("Location"))
		if region != "" {
			return result.Result{Status: status.OK, Region: region}
		}
		return result.Result{Status: status.Failed}
	}

	return result.Result{Status: status.Unexpected}
}
