package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"strings"
)

func Steam(c http.Client) result.Result {
	resp, err := url.GET(c, "https://store.steampowered.com")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	for _, c := range resp.Cookies() {
		if c.Name == "steamCountry" {
			i := strings.Index(c.Value, "%")
			if i == -1 {
				return result.Result{Status: status.No}
			}
			return result.Result{Status: status.OK, Region: strings.ToLower(c.Value[:i])}
		}
	}
	return result.Result{Status: status.No}
}
