package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func Epix(c http.Client) result.Result {
	url := "https://api.epix.com/v2/sessions"
	resp, err := urls.PostJson(c, url,
		`{"device":{"guid":"7a0baaaf-384c-45cd-a21d-310ca5d3002a","format":"console","os":"web","display_width":1865,"display_height":942,"app_version":"1.0.2","model":"browser","manufacturer":"google"},"apikey":"53e208a9bbaee479903f43b39d7301f7"}`,
		urls.H{"connection", "keep-alive"},
		urls.H{"traceparent", "00-000000000000000015b7efdb572b7bf2-4aefaea90903bd1f-01"},
		urls.H{"x-datadog-sampling-priority", "1"},
		urls.H{"x-datadog-trace-id", "1564983120873880562"},
		urls.H{"x-datadog-parent-id", "5399726519264460063"},
		urls.H{"origin", "https://www.mgmplus.com"},
		urls.H{"referer", "https://www.mgmplus.com/"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "error code") {
		return result.Result{Status: status.No}
	}
	if strings.Contains(s, "blocked") {
		return result.Result{Status: status.Banned}
	}
	var res struct {
		DeviceSession struct {
			SessionToken string `json:"session_token"`
		} `json:"device_session"`
	}
	if err := json.Unmarshal(b, &res); err != nil {
		// log.Println(err)
		return result.Result{Status: status.Failed, Err: err}
	}
	url2 := "https://api.epix.com/v2/movies/16921/play"
	resp2, err := urls.PostJson(c, url2, `{}`, urls.H{"X-Session-Token", res.DeviceSession.SessionToken})
	if err != nil {
		// log.Println(err)
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()
	b2, err := io.ReadAll(resp2.Body)
	if err != nil {
		// log.Println(err)
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res2 struct {
		Movie struct {
			Entitlements struct {
				Status string
			}
		}
	}
	if err := json.Unmarshal(b2, &res2); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}
	switch res2.Movie.Entitlements.Status {
	case "PROXY_DETECTED":
		return result.Result{Status: status.No, Info: "Proxy Detected"}
	case "GEO_BLOCKED":
		return result.Result{Status: status.No}
	case "NOT_SUBSCRIBED":
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.Failed}
}
