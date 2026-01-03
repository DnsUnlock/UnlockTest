package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Crave(c http.Client) result.Result {
	resp, err := url.GET(c, "https://capi.9c9media.com/destinations/se_atexace/platforms/desktop/bond/contents/2205173/contentpackages/4279732/manifest.mpd")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, `Geo Constraint Restrictions`) {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.OK}
}
