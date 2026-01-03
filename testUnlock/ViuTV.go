package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func ViuTV(c http.Client) result.Result {
	resp, err := url.PostJson(c, "https://api.viu.now.com/p8/3/getLiveURL",
		`{"callerReferenceNo":"20210726112323","contentId":"099","contentType":"Channel","channelno":"099","mode":"prod","deviceId":"29b3cb117a635d5b56","deviceType":"ANDROID_WEB"}`,
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res noweRes
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.ResponseCode == "SUCCESS" {
		return result.Result{Status: status.OK}
	} else if res.ResponseCode == "GEO_CHECK_FAIL" {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.Unexpected}
}
