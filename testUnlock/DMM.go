package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func DMM(c http.Client) result.Result {
	resp, err := url.GET(c, "https://bitcoin.dmm.com")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "This page is not available in your area") {
		return result.Result{Status: status.No}
	}
	if strings.Contains(s, "暗号資産") {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No, Info: "Unsupported"}
}
func DMMTV(c http.Client) result.Result {
	resp, err := url.PostJson(c, "https://api.beacon.dmm.com/v1/streaming/start", `{"player_name":"dmmtv_browser","player_version":"0.0.0","content_type_detail":"VOD_SVOD","content_id":"11uvjcm4fw2wdu7drtd1epnvz","purchase_product_id":null}`)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "FOREIGN") {
		return result.Result{Status: status.No}
	}
	if strings.Contains(s, "UNAUTHORIZED") {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No, Info: "Unsupported"}
}
