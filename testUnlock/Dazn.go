package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func Dazn(c http.Client) result.Result {
	resp, err := url.PostJson(c, "https://startup.core.indazn.com/misl/v5/Startup",
		`{"LandingPageKey":"generic","Languages":"zh-CN,zh,en","Platform":"web","PlatformAttributes":{},"Manufacturer":"","PromoCode":"","Version":"2"}`,
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	// log.Println(string(b))
	var res daznRes
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	if res.Region.IsAllowed {
		return result.Result{
			Status: status.OK,
			Region: res.Region.GeolocatedCountry,
		}
	}
	return result.Result{
		Status: status.No,
		Info:   res.Region.DisallowedReason,
	}
}

type daznRegion struct {
	IsAllowed             bool
	DisallowedReason      string
	GeolocatedCountry     string
	GeolocatedCountryName string
}

type daznRes struct {
	Region daznRegion
}
