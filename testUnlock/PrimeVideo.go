package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func PrimeVideo(c http.Client) result.Result {
	c.CheckRedirect = nil
	resp, err := url.GET(c, "https://www.primevideo.com")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if i := strings.Index(s, `"currentTerritory":`); i != -1 {
		return result.Result{
			Status: status.OK,
			Region: strings.ToLower(s[i+20 : i+22]),
		}
	}
	return result.Result{Status: status.No}
}
