package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Showmax(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.showmax.com/",
		url.H{"host", "www.showmax.com"},
		url.H{"connection", "keep-alive"},
		url.H{"upgrade-insecure-requests", "1"},
		url.H{"accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	regionStart := strings.Index(bodyString, "activeTerritory")
	if regionStart == -1 {
		return result.Result{Status: status.No}
	}

	regionEnd := strings.Index(bodyString[regionStart:], "\n")
	region := strings.TrimSpace(bodyString[regionStart+len("activeTerritory")+1 : regionStart+regionEnd])

	if region != "" {
		return result.Result{Status: status.OK, Region: strings.ToLower(region)}
	}

	return result.Result{Status: status.Unexpected}
}
