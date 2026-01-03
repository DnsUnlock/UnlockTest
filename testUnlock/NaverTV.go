package testUnlock

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func urlEncode(input string) string {
	return url.QueryEscape(input)
}

func generateHMACSignature(key, data string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func NaverTV(c http.Client) result.Result {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	signature := generateHMACSignature(
		"nbxvs5nwNG9QKEWK0ADjYA4JZoujF4gHcIwvoCxFTPAeamq5eemvt5IWAYXxrbYM",
		"https://apis.naver.com/now_web2/now_web_api/v1/clips/31030608/play-info"+strconv.FormatInt(timestamp, 10),
	)

	resp, err := urls.GET(c, "https://apis.naver.com/now_web2/now_web_api/v1/clips/31030608/play-info?msgpad="+strconv.FormatInt(timestamp, 10)+"&md="+urlEncode(signature),
		urls.H{"Connection", "keep-alive"},
		urls.H{"Accept", "application/json, text/plain, */*"},
		urls.H{"Origin", "https://tv.naver.com"},
		urls.H{"Referer", "https://tv.naver.com/v/31030608"},
	)

	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res struct {
		Playable string `json:"playable"`
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return result.Result{Status: status.Err, Err: err}
	}

	if res.Playable == "NOT_COUNTRY_AVAILABLE" {
		return result.Result{Status: status.No}
	}
	if resp.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
