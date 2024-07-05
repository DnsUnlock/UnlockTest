package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func Molotov(c http.Client) result.Result {
	resp, err := url.GET(c, "https://fapi.molotov.tv/v1/open-europe/is-france")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res struct {
		isFrance bool `json:"is_france"`
	}
	if err := json.Unmarshal(b, &res); err != nil {
		// log.Println(err)
		return result.Result{Status: status.Failed, Err: err}
	}

	if res.isFrance {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.No}
}
