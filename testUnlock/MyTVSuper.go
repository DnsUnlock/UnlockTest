package testUnlock

import (
	"context"
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/ua"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"time"
)

func MyTvSuper(c http.Client) result.Result {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r, err := http.NewRequestWithContext(ctx, "GET", "https://www.mytvsuper.com/api/auth/getSession/self/", nil)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	r.Header.Set("User-Agent", ua.PCBrowserUA())
	r.Header.Set("Content-Type", "application/json")

	resp, err := url.RunFor(c, r)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res mytvsuperRes
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.Region == 1 {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}

type mytvsuperRes struct {
	Region int
}
