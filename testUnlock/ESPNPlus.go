package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
)

func ESPNPlus(c http.Client) result.Result {
	resp, err := url.PostForm(c, "https://espn.api.edge.bamgrid.com/token",
		`grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Atoken-exchange&latitude=0&longitude=0&platform=browser&subject_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJjYWJmMDNkMi0xMmEyLTQ0YjYtODJjOS1lOWJkZGNhMzYwNjkiLCJhdWQiOiJ1cm46YmFtdGVjaDpzZXJ2aWNlOnRva2VuIiwibmJmIjoxNjMyMjMwMTY4LCJpc3MiOiJ1cm46YmFtdGVjaDpzZXJ2aWNlOmRldmljZSIsImV4cCI6MjQ5NjIzMDE2OCwiaWF0IjoxNjMyMjMwMTY4LCJqdGkiOiJhYTI0ZWI5Yi1kNWM4LTQ5ODctYWI4ZS1jMDdhMWVhMDgxNzAifQ.8RQ-44KqmctKgdXdQ7E1DmmWYq0gIZsQw3vRL8RvCtrM_hSEHa-CkTGIFpSLpJw8sMlmTUp5ZGwvhghX-4HXfg&subject_token_type=urn%3Abamtech%3Aparams%3Aoauth%3Atoken-type%3Adevice`,
		url.H{"authorization", "Bearer ZXNwbiZicm93c2VyJjEuMC4w.ptUt7QxsteaRruuPmGZFaJByOoqKvDP2a5YkInHrc7c"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}
	d := `{"query":"mutation registerDevice($input: RegisterDeviceInput!) {\n            registerDevice(registerDevice: $input) {\n                grant {\n                    grantType\n                    assertion\n                }\n            }\n        }","variables":{"input":{"deviceFamily":"browser","applicationRuntime":"chrome","deviceProfile":"windows","deviceLanguage":"zh-CN","attributes":{"osDeviceIds":[],"manufacturer":"microsoft","model":null,"operatingSystem":"windows","operatingSystemVersion":"10.0","browserName":"chrome","browserVersion":"96.0.4664"}}}}`
	resp2, err := url.PostJson(c, "https://espn.api.edge.bamgrid.com/graph/v1/device/graphql", d,
		url.H{"authorization", "ZXNwbiZicm93c2VyJjEuMC4w.ptUt7QxsteaRruuPmGZFaJByOoqKvDP2a5YkInHrc7c"},
	)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	b2, err := io.ReadAll(resp2.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res2 struct {
		Extensions struct {
			Sdk struct {
				Session struct {
					Location struct {
						CountryCode string `json:"countryCode"`
					}
					InSupportedLocation bool `json:"inSupportedLocation"`
				}
			}
		}
	}
	if err := json.Unmarshal(b2, &res2); err != nil {
		return result.Result{Status: status.Failed, Err: err}
	}
	if res2.Extensions.Sdk.Session.Location.CountryCode == "US" && res2.Extensions.Sdk.Session.InSupportedLocation {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}
