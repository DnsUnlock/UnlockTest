package testUnlock

import (
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/ua"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
)

func PrettyDerbyJP(c http.Client) result.Result {
	for i := 0; i < 3; i++ {
		//resp, err := GET_Dalvik(c, "https://api-umamusume.cygames.jp/")
		resp, err := url.GET(c, "https://api-umamusume.cygames.jp/",
			url.H{"user-agent", ua.MobileBrowserUA()},
			url.H{"connection", "keep-alive"},
		)
		if err != nil {
			return result.Result{Status: status.NetworkErr, Err: err}
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 404:
			return result.Result{Status: status.OK}
		}
	}
	return result.Result{Status: status.No}
}
