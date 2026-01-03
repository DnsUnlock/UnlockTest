package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/md5"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/ua"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func extractWowowContentURL(body string) string {
	re := regexp.MustCompile(`https://wod.wowow.co.jp/content/\d+`)
	match := re.FindString(body)
	return match
}

func extractWowowMetaID(body string) string {
	re := regexp.MustCompile(`https://wod.wowow.co.jp/watch/(\d+)`)
	matches := re.FindStringSubmatch(body)

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func Wowow(c http.Client) result.Result {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	//resp1, err := GET(c, "https://www.wowow.co.jp/drama/original/json/lineup.jsonf?_=" + strconv.FormatInt(timestamp, 10),
	//    H{"Accept", "application/json, text/javascript, */*; q=0.01"},
	//    H{"Referer", "https://www.wowow.co.jp/drama/original/"},
	//    H{"X-Requested-With", "XMLHttpRequest"},
	//)
	//if err != nil {
	//    return result.Result{Status: status.NetworkErr, Err: err}
	//}
	//defer resp1.Body.Close()

	//body1, err := ioutil.ReadAll(resp1.Body)
	//if err != nil {
	//    return result.Result{Status: status.NetworkErr, Err: err}
	//}

	//fmt.Printf(string(body1))

	//var res1 []struct {
	//    DramaLink string `json:"link"`
	//}
	//if err := json.Unmarshal(body1, &res1); err != nil {
	//	return result.Result{Status: status.Failed, Err: err}
	//}

	//resp2, err := GET(c, res1[1].DramaLink)
	resp2, err := url.GET(c, "https://www.wowow.co.jp/drama/original/yukai/")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()

	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	wodUrl := extractWowowContentURL(string(body2))
	if wodUrl == "" {
		return result.Result{Status: status.Failed, Err: err}
	}

	resp3, err := url.GET(c, wodUrl)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp3.Body.Close()

	body3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	metaID := extractWowowMetaID(string(body3))

	vUid := md5.Sum(strconv.FormatInt(timestamp, 10))

	resp4, err := url.PostJson(c, "https://mapi.wowow.co.jp/api/v1/playback/auth",
		`{"meta_id":`+metaID+`,"vuid":"`+vUid+`","device_code":1,"app_id":1,"ua":"`+ua.PCBrowserUA()+`"}`,
		url.H{"accept", "application/json, text/plain, */*"},
		url.H{"content-type", "application/json;charset=UTF-8"},
		url.H{"origin", "https://wod.wowow.co.jp"},
		url.H{"referer", "https://wod.wowow.co.jp/"},
		url.H{"x-requested-with", "XMLHttpRequest"},
	)

	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp4.Body.Close()

	body4, err := ioutil.ReadAll(resp4.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if strings.Contains(string(body4), "VPN") {
		return result.Result{Status: status.No}
	}

	if strings.Contains(string(body4), "playback_session_id") {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
