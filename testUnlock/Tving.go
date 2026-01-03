package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func Tving(c http.Client) result.Result {
	resp, err := url.GET(c, "https://api.tving.com/v2a/media/stream/info?apiKey=1e7952d0917d6aab1f0293a063697610&mediaCode=RV60891248")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res struct {
		Body struct {
			Result struct {
				Code string `json:"code"`
			} `json:"result"`
		} `json:"body"`
	}

	if err := json.Unmarshal(b, &res); err != nil {
		//log.Println(err)
		return result.Result{Status: status.Failed, Err: err}
	}
	if res.Body.Result.Code == "001" {
		return result.Result{Status: status.No}
	}

	if res.Body.Result.Code == "000" {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
