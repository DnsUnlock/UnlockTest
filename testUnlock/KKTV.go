package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func KKTV(c http.Client) result.Result {
	resp, err := url.GET(c, "https://api.kktv.me/v3/ipcheck")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	// log.Println(string(b))
	var res struct {
		Data struct {
			Country   string
			IsAllowed bool `json:"is_allowed"`
		}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.Data.Country == "TW" && res.Data.IsAllowed {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
