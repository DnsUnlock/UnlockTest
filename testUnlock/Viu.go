package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"strings"
)

func ViuCom(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.viu.com")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if location := resp.Header.Get("location"); location != "" {
		region := strings.Split(location, "/")[4]
		if region == "no-service" {
			return result.Result{Status: status.No}
		}
		return result.Result{Status: status.OK, Region: region}
	}
	return result.Result{Status: status.No}
}
