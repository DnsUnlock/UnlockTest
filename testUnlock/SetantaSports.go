package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"net/http"
	"strings"
)

func SetantaSports(c http.Client) result.Result {
	req, err := http.NewRequest("GET", "https://dce-frontoffice.imggaming.com/api/v2/consent-prompt", nil)
	if err != nil {
		return result.Result{Status: status.Failed}
	}
	req.Header.Set("Realm", "dce.adjara")
	req.Header.Set("x-api-key", "857a1e5d-e35e-4fdf-805b-a87b6f8364bf")

	resp, err := url.RunFor(c, req)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return result.Result{Status: status.Unexpected}
	}

	flag, ok := data["outsideAllowedTerritories"].(bool)
	if !ok {
		return result.Result{Status: status.Unexpected}
	}

	if strings.HasPrefix(resp.Status, "200") {
		if flag {
			return result.Result{Status: status.No}
		}
		return result.Result{Status: status.OK}
	}

	return result.Result{Status: status.Unexpected}
}
