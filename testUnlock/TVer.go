package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io/ioutil"
	"net/http"
	"regexp"
)

func extractTVerPolicyKey(body string) string {
	re := regexp.MustCompile(`policyKey:"([^"]+)"`)
	matches := re.FindStringSubmatch(body)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func extractTVerDeliveryConfigID(body string) string {
	re := regexp.MustCompile(`deliveryConfigId:"([^"]+)"`)
	matches := re.FindStringSubmatch(body)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func TVer(c http.Client) result.Result {
	resp1, err := url.PostForm(c, "https://platform-api.tver.jp/v2/api/platform_users/browser/create", "device_type=pc",
		url.H{"origin", "https://s.tver.jp"},
		url.H{"referer", "https://s.tver.jp/"},
		url.H{"accept-language", "en-US,en;q=0.9"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp1.Body.Close()

	body1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res1 struct {
		Result struct {
			PlatformUid   string `json:"platform_uid"`
			PlatformToken string `json:"platform_token"`
		} `json:"result"`
	}

	if err := json.Unmarshal(body1, &res1); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}

	resp2, err := url.GET(c, "https://platform-api.tver.jp/service/api/v1/callHome?platform_uid="+res1.Result.PlatformUid+"&platform_token="+res1.Result.PlatformToken+"&require_data=mylist%2Cresume%2Clater",
		url.H{"origin", "https://tver.jp"},
		url.H{"referer", "https://tver.jp/"},
		url.H{"accept-language", "en-US,en;q=0.9"},
		url.H{"x-tver-platform-type", "web"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()

	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res2 struct {
		Result struct {
			Components []struct {
				ComponentID string `json:"componentID"`
				Contents    []struct {
					Content struct {
						EpisodeID string `json:"id"`
					} `json:"content"`
				} `json:"contents"`
			} `json:"components"`
		} `json:"result"`
	}
	if err := json.Unmarshal(body2, &res2); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}

	EpisodeID := ""
	for _, component := range res2.Result.Components {
		if component.ComponentID == "newer." {
			if len(component.Contents) > 0 {
				EpisodeID = component.Contents[0].Content.EpisodeID
			}
			break
		}
	}

	resp3, err := url.GET(c, "https://statics.tver.jp/content/episode/"+EpisodeID+".json",
		url.H{"origin", "https://tver.jp"},
		url.H{"referer", "https://tver.jp/"},
		url.H{"accept-language", "en-US,en;q=0.9"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp3.Body.Close()

	body3, err := ioutil.ReadAll(resp3.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res3 struct {
		Video struct {
			AccountID  string `json:"accountID"`
			PlayerID   string `json:"playerID"`
			VideoID    string `json:"videoID"`
			VideoRefID string `json:"videoRefID"`
		} `json:"video"`
	}
	if err := json.Unmarshal(body3, &res3); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}

	AccountID := res3.Video.AccountID
	PlayerID := res3.Video.PlayerID
	VideoID := res3.Video.VideoID
	VideoRefID := res3.Video.VideoRefID

	resp4, err := url.GET(c, "https://players.brightcove.net/"+AccountID+"/"+PlayerID+"_default/index.min.js",
		url.H{"Referer", "https://tver.jp/"},
		url.H{"accept-language", "en-US,en;q=0.9"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp4.Body.Close()

	body4, err := ioutil.ReadAll(resp4.Body)
	if err != nil || len(body4) == 0 {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	PolicyKey := extractTVerPolicyKey(string(body4))
	DeliveryConfigID := extractTVerDeliveryConfigID(string(body4))

	var resp5 *http.Response
	if VideoRefID == "" {
		resp5, err = url.GET(c, "https://edge.api.brightcove.com/playback/v1/accounts/"+AccountID+"/videos/"+VideoID+"?config_id="+DeliveryConfigID,
			url.H{"accept", "application/json;pk=" + PolicyKey},
			url.H{"origin", "https://tver.jp"},
			url.H{"referer", "https://tver.jp/"},
			url.H{"accept-language", "en-US,en;q=0.9"},
		)
	} else {
		resp5, err = url.GET(c, "https://edge.api.brightcove.com/playback/v1/accounts/"+AccountID+"/videos/ref%3A"+VideoRefID,
			url.H{"accept", "application/json;pk=" + PolicyKey},
			url.H{"origin", "https://tver.jp"},
			url.H{"referer", "https://tver.jp/"},
			url.H{"accept-language", "en-US,en;q=0.9"},
		)
	}
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp5.Body.Close()

	body5, err := ioutil.ReadAll(resp5.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res4a []struct {
		ErrorSubcode string `json:"error_subcode"`
		//ClientGeo    string `json:"client_geo"`
	}
	var res4b struct {
		AccountID string `json:"account_id"`
	}
	if err := json.Unmarshal(body5, &res4a); err != nil {
		if err := json.Unmarshal(body5, &res4b); err != nil {
			return result.Result{Status: status.Unexpected}
		}
		if res4b.AccountID != "" {
			return result.Result{Status: status.OK}
		}
	}

	if res4a[0].ErrorSubcode == "CLIENT_GEO" {
		return result.Result{Status: status.No}
	}

	return result.Result{Status: status.Unexpected}
}
