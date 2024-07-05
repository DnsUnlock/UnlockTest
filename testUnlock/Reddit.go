package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Reddit(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.reddit.com/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if resp.StatusCode == 200 || resp.StatusCode == 302 {
		return result.Result{Status: status.OK}
	}

	if resp.StatusCode == 403 && strings.Contains(bodyString, "blocked") {
		return result.Result{Status: status.No}
	}

	return result.Result{Status: status.Unexpected}
}
