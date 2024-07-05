package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func FuboTV(c http.Client) result.Result {
	randNum := strconv.Itoa(rand.Intn(2))
	resp, err := url.GET(c, "https://api.fubo.tv/appconfig/v1/homepage?platform=web&client_version=R20230310."+randNum+"&nav=v0")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "Forbidden IP") {
		return result.Result{Status: status.No, Info: "IP Forbidden"}
	}
	if strings.Contains(s, "No Subscription") {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
