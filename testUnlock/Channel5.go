package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strconv"
	"time"
)

func Channel5(c http.Client) result.Result {
	timestamp := time.Now().Unix()
	resp, err := url.GET(c, "https://cassie.channel5.com/api/v2/live_media/my5desktopng/C5.json?timestamp="+strconv.FormatInt(timestamp, 10)+"&auth=0_rZDiY0hp_TNcDyk2uD-Kl40HqDbXs7hOawxyqPnbI")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}

	var res struct {
		Code string `json:"code"`
	}

	if err := json.Unmarshal(b, &res); err != nil {
		//log.Println(err)
		return result.Result{Status: status.Failed, Err: err}
	}
	if res.Code == "3000" {
		return result.Result{Status: status.No}
	}

	if res.Code == "3001" {
		return result.Result{Status: status.No, Info: "Proxy Detected"}
	}

	if res.Code == "4003" {
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
