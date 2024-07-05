package testUnlock

import (
	"context"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"time"
)

// Project Sekai: Colorful Stage
func PJSK(c http.Client) result.Result {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://game-version.sekai.colorfulpalette.org/1.8.1/3ed70b6a-8352-4532-b819-108837926ff5", nil)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	req.Header.Set("User-Agent", "pjsekai/48 CFNetwork/1240.0.4 Darwin/20.6.0")

	resp, err := url.RunFor(c, req)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return result.Result{Status: status.OK}
	case 403:
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.Unexpected}
}
