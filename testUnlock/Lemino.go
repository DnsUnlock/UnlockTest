package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func Lemino(c http.Client) result.Result {
	resp, err := url.PostJson(c, "https://if.lemino.docomo.ne.jp/v1/user/delivery/watch/ready",
		`{"inflow_flows":[null,"crid://plala.iptvf.jp/group/b100ce3"],"play_type":1,"key_download_only":null,"quality":null,"groupcast":null,"avail_status":"1","terminal_type":3,"test_account":0,"content_list":[{"kind":"main","service_id":null,"cid":"00lm78dz30","lid":"a0lsa6kum1","crid":"crid://plala.iptvf.jp/vod/0000000000_00lm78dymn","preview":0,"trailer":0,"auto_play":0,"stop_position":0}]}`,
		url.H{"accept", "application/json, text/plain, */*"},
		url.H{"accept-language", "en-US,en;q=0.9"},
		url.H{"content-type", "application/json"},
		url.H{"origin", "https://lemino.docomo.ne.jp"},
		url.H{"referer", "https://lemino.docomo.ne.jp/"},
		url.H{"x-service-token", "f365771afd91452fa279863f240c233d"},
		url.H{"x-trace-id", "556db33f-d739-4a82-84df-dd509a8aa179"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 403 {
		return result.Result{Status: status.No}
	}

	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
