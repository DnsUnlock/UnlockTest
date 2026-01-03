package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"net/http/cookiejar"
)

func BahamutAnime(c http.Client) result.Result {
	c.Jar, _ = cookiejar.New(nil)
	resp, err := url.GET(c, "https://ani.gamer.com.tw/ajax/getdeviceid.php")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res struct {
		AnimeSn  int
		Deviceid string
	}
	if err := json.Unmarshal(b, &res); err != nil {
		if err.Error() == "invalid character '<' looking for beginning of value" {
			return result.Result{Status: status.No}
		}
		return result.Result{Status: status.Err, Err: err}
	}
	resp, err = url.GET(c, "https://ani.gamer.com.tw/ajax/token.php?adID=89422&sn=14667&device="+res.Deviceid)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.AnimeSn != 0 {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.Unexpected}
}
