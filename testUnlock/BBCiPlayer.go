package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func BBCiPlayer(c http.Client) result.Result {
	resp, err := url.GET(c, "https://open.live.bbc.co.uk/mediaselector/6/select/version/2.0/mediaset/pc/vpid/bbc_one_london/format/json/jsfunc/JS_callbacks0")

	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		return result.Result{Status: status.Failed}
	}

	if resp.StatusCode == 200 {
		if strings.Contains(bodyString, "geolocation") {
			return result.Result{Status: status.No}
		}
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Failed}
}
