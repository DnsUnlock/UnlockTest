package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Stan(c http.Client) result.Result {
	resp, err := url.PostJson(c, "https://api.stan.com.au/login/v1/sessions/web/account", `{}`)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		return result.Result{Status: status.Failed}
	}

	if strings.Contains(bodyString, "Access Denied") {
		return result.Result{Status: status.No}
	}

	if strings.Contains(bodyString, "VPNDetected") {
		return result.Result{Status: status.No, Info: "VPN Detected"}
	}

	if resp.StatusCode == 400 {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
