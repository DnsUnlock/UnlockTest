package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/md5"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"github.com/DnsUnlock/UnlockTest/lib/uuid"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AISPlay(c http.Client) result.Result {
	userId := "09e8b25510"
	fakeApiKey := md5.Sum(uuid.New())
	fakeUdid := md5.Sum(uuid.New())
	timestamp := time.Now().Unix()

	resp1, err := url.PostJson(c, "https://web-tls.ais-vidnt.com/device/login/?d=gstweb&gst=1&user="+userId+"&pass=e49e9f9e7f", `------WebKitFormBoundaryBj2RhUIW7BtRvfK0--\r\n`,
		url.H{"accept-language", "th"},
		url.H{"api-version", "2.8.2"},
		url.H{"api_key", fakeApiKey},
		url.H{"content-type", "multipart/form-data; boundary=----WebKitFormBoundaryBj2RhUIW7BtRvfK0"},
		url.H{"device-info", "com.vimmi.ais.portal, Windows + Chrome, AppVersion: 4.9.97, 10, language: tha"},
		url.H{"origin", "https://aisplay.ais.co.th"},
		url.H{"privateid", userId},
		url.H{"referer", "https://aisplay.ais.co.th/"},
		url.H{"time", strconv.FormatInt(timestamp, 10)},
		url.H{"udid", fakeUdid},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp1.Body.Close()

	body1, err := io.ReadAll(resp1.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res1 struct {
		Info struct {
			Sid string `json:"sid"`
			Dat string `json:"dat"`
		} `json:"info"`
	}
	if err := json.Unmarshal(body1, &res1); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}
	sId := res1.Info.Sid
	datAuth := res1.Info.Dat

	timestamp = time.Now().Unix()

	resp2, err := url.GET(c, "https://web-sila.ais-vidnt.com/playtemplate/?d=gstweb",
		url.H{"accept-language", "en-US,en;q=0.9"},
		url.H{"api-version", "2.8.2"},
		url.H{"api_key", fakeApiKey},
		url.H{"dat", datAuth},
		url.H{"device-info", "com.vimmi.ais.portal, Windows + Chrome, AppVersion: 0.0.0, 10, Language: unknown"},
		url.H{"origin", "https://web-player.ais-vidnt.com"},
		url.H{"privateid", userId},
		url.H{"referer", "https://web-player.ais-vidnt.com/"},
		url.H{"sid", sId},
		url.H{"time", strconv.FormatInt(timestamp, 10)},
		url.H{"udid", fakeUdid},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res2 struct {
		Info struct {
			Live string `json:"live"`
		} `json:"Info"`
	}
	if err := json.Unmarshal(body2, &res2); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}

	mediaId := "B0006"
	realLiveUrl := strings.ReplaceAll(res2.Info.Live, "{MID}", mediaId)
	realLiveUrl = strings.ReplaceAll(realLiveUrl, "metadata.xml", "metadata.json")

	resp3, err := url.GET(c, realLiveUrl+"-https&tuid="+userId+"&tdid="+fakeUdid+"&chunkHttps=true&origin=anevia",
		url.H{"Accept-Language", "en-US,en;q=0.9"},
		url.H{"Origin", "https://web-player.ais-vidnt.com"},
		url.H{"Referer", "https://web-player.ais-vidnt.com/"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp3.Body.Close()

	body3, err := io.ReadAll(resp3.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res3 struct {
		PlaybackUrls []struct {
			PlayUrl string `json:"url"`
		} `json:"playbackUrls"`
	}
	if err := json.Unmarshal(body3, &res3); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}

	resp4, err := url.GET(c, res3.PlaybackUrls[0].PlayUrl,
		url.H{"Accept-Language", "en-US,en;q=0.9"},
		url.H{"Origin", "https://web-player.ais-vidnt.com"},
		url.H{"Referer", "https://web-player.ais-vidnt.com/"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp4.Body.Close()

	if resp4.StatusCode == 403 {
		return result.Result{Status: status.No}
	}

	if resp4.StatusCode == 200 {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
