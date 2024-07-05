package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"strings"
)

func IQiYi(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.iq.com")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	s := resp.Header.Get("x-custom-client-ip")
	if s == "" {
		return result.Result{Status: status.No}
	}
	i := strings.Index(s, ":")
	if i == -1 {
		return result.Result{Status: status.No}
	}
	region := s[i+1:]
	if region == "ntw" {
		region = "tw"
	}
	return result.Result{Status: status.OK, Region: region}
}
