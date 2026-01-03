package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Afreeca(c http.Client) result.Result {
	resp, err := url.GET(c, "https://vod.afreecatv.com/player/97464151")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if strings.Contains(bodyString, "document.location.href='https://vod.afreecatv.com'") {
		return result.Result{Status: status.No}
	}

	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
