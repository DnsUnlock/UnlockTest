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
	"strings"
)

func EurosportRO(c http.Client) result.Result {
	fakeUuid := md5.Sum(uuid.New())

	resp1, err := url.GET(c, "https://eu3-prod-direct.eurosport.ro/token?realm=eurosport",
		url.H{"accept", "*/*"},
		url.H{"accept-language", "en-US,en;q=0.9"},
		url.H{"origin", "https://www.eurosport.ro"},
		url.H{"referer", "https://www.eurosport.ro/"},
		url.H{"x-device-info", "escom/0.295.1 (unknown/unknown; Windows/10; " + fakeUuid + ")"},
		url.H{"x-disco-client", "WEB:UNKNOWN:escom:0.295.1"},
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
		Data struct {
			Attributes struct {
				Token string `json:"token"`
			} `json:"attributes"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body1, &res1); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}

	token := res1.Data.Attributes.Token
	sourceSystemId := "eurosport-vid2133403"

	resp2, err := url.GET(c, "https://eu3-prod-direct.eurosport.ro/playback/v2/videoPlaybackInfo/sourceSystemId/"+sourceSystemId+"?usePreAuth=true",
		url.H{"Authorization", "Bearer " + token},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	bodyString := string(body2)

	if strings.Contains(bodyString, "access.denied.geoblocked") {
		return result.Result{Status: status.No}
	}

	if strings.Contains(bodyString, "eurosport-vod") {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
