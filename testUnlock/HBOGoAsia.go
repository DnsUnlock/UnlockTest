package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func HboGoAsia(c http.Client) result.Result {
	resp, err := url.GET(c, "https://api2.hbogoasia.com/v1/geog?lang=undefined&version=0&bundleId=www.hbogoasia.com")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res hboRes
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.Territory == "" {
		return result.Result{Status: status.No}
	}
	return result.Result{Status: status.OK, Region: strings.ToLower(res.Country)}
}

type hboRes struct {
	Country   string
	Territory string
}
