package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Channel10(c http.Client) result.Result {
	resp, err := url.GET(c, "https://10play.com.au/geo-web")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res struct {
		allow bool
	}
	if err := json.Unmarshal(b, &res); err != nil {
		if strings.Contains(string(b), "not available") {
			return result.Result{Status: status.No}
		}
		return result.Result{Status: status.Err, Err: err}
	}
	if res.allow {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
