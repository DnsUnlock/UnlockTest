package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func Fox(c http.Client) result.Result {
	url := "https://x-live-fox-stgec.uplynk.com/ausw/slices/8d1/d8e6eec26bf544f084bad49a7fa2eac5/8d1de292bcc943a6b886d029e6c0dc87/G00000000.ts?pbs=c61e60ee63ce43359679fb9f65d21564&cloud=aws&si=0"
	resp, err := urls.GET(c, url)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return result.Result{Status: status.OK}
	case 403:
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.Unexpected}
}
