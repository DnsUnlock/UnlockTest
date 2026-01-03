package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func Abema(c http.Client) result.Result {
	resp, err := url.GETDalvik(c, "https://api.abema.io/v1/ip/check?device=android")
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
		IsoCountryCode string
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.IsoCountryCode == "JP" {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.Restricted, Info: "Oversea Only"}
}
