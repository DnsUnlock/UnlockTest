package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"strings"
)

func Mora(c http.Client) result.Result {
	resp, err := url.GET(c, "https://mora.jp/buy?__requestToken=1713764407153&returnUrl=https%3A%2F%2Fmora.jp%2Fpackage%2F43000087%2FTFDS01006B00Z%2F%3Ffmid%3DTOPRNKS%26trackMaterialNo%3D31168909&fromMoraUx=false&deleteMaterial=")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if resp.StatusCode == 403 || resp.StatusCode == 500 {
		return result.Result{Status: status.No}
	}

	if resp.StatusCode == 302 {
		if strings.Contains(resp.Header.Get("Location"), "error") {
			return result.Result{Status: status.No}
		}
		if strings.Contains(resp.Header.Get("Location"), "signin") {
			return result.Result{Status: status.OK}
		}
	}

	return result.Result{Status: status.Unexpected}
}
