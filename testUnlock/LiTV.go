package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func LiTV(c http.Client) result.Result {
	resp, err := url.PostJson(c, "https://www.litv.tv/api/get-urls-no-auth",
		`{"AssetId": "vod71211-000001M001_1500K","MediaType": "vod","puid": "d66267c2-9c52-4b32-91b4-3e482943fe7e"}`,
		url.H{"Cookie", "PUID=34eb9a17-8834-4f83-855c-69382fd656fa; L_PUID=34eb9a17-8834-4f83-855c-69382fd656fa; device-id=f4d7faefc54f476bb2e7e27b7482469a"},
		url.H{"Origin", "https://www.litv.tv"},
		url.H{"Referer", "https://www.litv.tv/drama/watch/VOD00331042"},
		url.H{"Priority", "u=1, i"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	if resp.StatusCode == 200 {
		if strings.Contains(bodyString, "OutsideRegionError") {
			return result.Result{Status: status.No}
		}
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
