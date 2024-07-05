package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"regexp"
)

func extractDirecTVGORegion(url string) string {
	re := regexp.MustCompile(`https?://www\.directvgo\.com/([^/]+)/`)

	matches := re.FindStringSubmatch(url)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func DirecTVGO(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.directvgo.com/registrarse")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if err != nil {
		return result.Result{Status: status.Failed}
	}

	if resp.StatusCode == 403 {
		return result.Result{Status: status.No}
	}

	if resp.StatusCode == 301 {
		if region := extractDirecTVGORegion(resp.Header.Get("Location")); region != "" {
			return result.Result{Status: status.OK, Region: region}
		}
		return result.Result{Status: status.Unexpected}
	}

	return result.Result{Status: status.Unexpected}
}
