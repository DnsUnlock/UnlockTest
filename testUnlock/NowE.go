package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	//"log"
)

func NowE(c http.Client) result.Result {
	resp, err := url.PostJson(c, "https://webtvapi.nowe.com/16/1/getVodURL",
		`{"contentId":"202403181904703","contentType":"Vod","pin":"","deviceName":"Browser","deviceId":"w-663bcc51-913c-913c-913c-913c913c","deviceType":"WEB","secureCookie":null,"callerReferenceNo":"W17151951620081575","profileId":null,"mupId":null}`,
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	//log.Println(string(b))
	var res noweRes
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Unexpected, Err: err}
	}
	if res.ResponseCode == "SUCCESS" {
		return result.Result{Status: status.OK}
	} else if res.ResponseCode == "GEO_CHECK_FAIL" {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.Unexpected}
}

type noweRes struct {
	ResponseCode string
}
