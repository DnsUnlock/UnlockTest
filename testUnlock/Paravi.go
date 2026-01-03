package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func Paravi(c http.Client) result.Result {
	resp, err := url.PostJson(c, "https://api.paravi.jp/api/v1/playback/auth",
		`{"meta_id":17414,"vuid":"3b64a775a4e38d90cc43ea4c7214702b","device_code":1,"app_id":1}`,
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if resp.StatusCode == 403 {
		return result.Result{Status: status.No}
	}

	var res struct {
		Error struct {
			Type string
		}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}

	if res.Error.Type == "Unauthorized" {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.Unexpected}
}
