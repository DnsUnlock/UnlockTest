package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func NBA_TV(c http.Client) result.Result {
	resp, err := urls.GET(c, "https://www.nba.com/watch/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "Service is not available in your region") {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.OK}
}
