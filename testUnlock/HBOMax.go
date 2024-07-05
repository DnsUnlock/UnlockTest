package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"strings"
)

func HBOMax(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.hbomax.com/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if strings.Contains(resp.Header.Get("location"), "geo-availability") {
		return result.Result{Status: status.No}
	}
	t := strings.Split(resp.Header.Get("location"), "/")
	region := ""
	if len(t) >= 4 {
		region = strings.Split(resp.Header.Get("location"), "/")[3]
	}
	return result.Result{Status: status.OK, Region: strings.ToLower(region)}
}
