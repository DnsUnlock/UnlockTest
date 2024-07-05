package testUnlock

import (
	"context"
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"time"
)

func Telasa(c http.Client) result.Result {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api-videopass-anon.kddi-video.com/v1/playback/system_status", nil)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	req.Header.Set("X-Device-ID", "d36f8e6b-e344-4f5e-9a55-90aeb3403799")

	resp, err := url.RunFor(c, req)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	// log.Println(string(b))
	var res struct {
		Status struct {
			Type    string
			Subtype string
		}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.Status.Subtype == "IPLocationNotAllowed" {
		return result.Result{Status: status.No}
	}
	if res.Status.Type != "" {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.Unexpected}
}
