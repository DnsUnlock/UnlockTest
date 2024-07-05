package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Philo(c http.Client) result.Result {
	resp, err := url.GET(c, "https://content-us-east-2-fastly-b.www.philo.com/geo")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	if strings.Contains(string(b), "SUCCESS") {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
