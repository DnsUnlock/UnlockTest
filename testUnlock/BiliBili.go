package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func bilibili(c http.Client, url string) result.Result {
	resp, err := urls.GET(c, url)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res struct {
		Code int
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.Code == -10403 || res.Code == 10004001 || res.Code == 10003003 {
		return result.Result{Status: status.No}
	}
	if resp.StatusCode == 412 {
		return result.Result{Status: status.Failed}
	}
	if res.Code == 0 {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.Unexpected}
}

func BilibiliHKMO(c http.Client) result.Result {
	return bilibili(c, "https://api.bilibili.com/pgc/player/web/playurl?avid=473502608&cid=845838026&qn=0&type=&otype=json&ep_id=678506&fourk=1&fnver=0&fnval=16&module=bangumi")
}

func BilibiliTW(c http.Client) result.Result {
	return bilibili(c, "https://api.bilibili.com/pgc/player/web/playurl?avid=50762638&cid=100279344&qn=0&type=&otype=json&ep_id=268176&fourk=1&fnver=0&fnval=16&module=bangumi")
}

func BilibiliSEA(c http.Client) result.Result {
	return bilibili(c, "https://api.bilibili.tv/intl/gateway/web/playurl?s_locale=en_US&platform=web&ep_id=347666")
}

func BilibiliTH(c http.Client) result.Result {
	return bilibili(c, "https://api.bilibili.tv/intl/gateway/web/playurl?s_locale=en_US&platform=web&ep_id=10077726")
}

func BilibiliID(c http.Client) result.Result {
	return bilibili(c, "https://api.bilibili.tv/intl/gateway/web/playurl?s_locale=en_US&platform=web&ep_id=11130043")
}

func BilibiliVN(c http.Client) result.Result {
	return bilibili(c, "https://api.bilibili.tv/intl/gateway/web/playurl?s_locale=en_US&platform=web&ep_id=11405745")
}
