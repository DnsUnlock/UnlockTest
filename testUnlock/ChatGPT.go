package testUnlock

import (
	errs "github.com/DnsUnlock/UnlockTest/lib/err"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func SupportGPT(loc string) bool {
	var GPT_SUPPORT_COUNTRY = []string{
		"AL", "DZ", "AD", "AO", "AG", "AR", "AM", "AU", "AT", "AZ", "BS", "BD", "BB", "BE", "BZ", "BJ", "BT", "BA", "BW", "BR", "BG", "BF", "CV", "CA", "CL", "CO", "KM", "CR", "HR", "CY", "DK", "DJ", "DM", "DO", "EC", "SV", "EE", "FJ", "FI", "FR", "GA", "GM", "GE", "DE", "GH", "GR", "GD", "GT", "GN", "GW", "GY", "HT", "HN", "HU", "IS", "IN", "ID", "IQ", "IE", "IL", "IT", "JM", "JP", "JO", "KZ", "KE", "KI", "KW", "KG", "LV", "LB", "LS", "LR", "LI", "LT", "LU", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MR", "MU", "MX", "MC", "MN", "ME", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NZ", "NI", "NE", "NG", "MK", "NO", "OM", "PK", "PW", "PA", "PG", "PE", "PH", "PL", "PT", "QA", "RO", "RW", "KN", "LC", "VC", "WS", "SM", "ST", "SN", "RS", "SC", "SL", "SG", "SK", "SI", "SB", "ZA", "ES", "LK", "SR", "SE", "CH", "TH", "TG", "TO", "TT", "TN", "TR", "TV", "UG", "AE", "US", "UY", "VU", "ZM", "BO", "BN", "CG", "CZ", "VA", "FM", "MD", "PS", "KR", "TW", "TZ", "TL", "GB",
	}
	for _, s := range GPT_SUPPORT_COUNTRY {
		if loc == s {
			return true
		}
	}
	return false
}

func ChatGPT(c http.Client) result.Result {
	resp, err := url.GET(c, "https://chat.openai.com/cdn-cgi/trace")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: errs.Network}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: errs.Network}
	}
	s := string(b)
	i := strings.Index(s, "loc=")
	if i == -1 {
		return result.Result{Status: status.Unexpected}
	}
	s = s[i+4:]
	i = strings.Index(s, "\n")
	if i == -1 {
		return result.Result{Status: status.Unexpected}
	}
	loc := s[:i]

	resp, err = url.GET(c, "https://chat.openai.com")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: errs.Network}
	}
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: errs.Network}
	}
	if strings.Contains(string(b), "VPN") {
		return result.Result{Status: status.Banned, Info: "VPN Blocked"}
	}
	if resp.StatusCode == 429 {
		return result.Result{Status: status.Restricted, Region: strings.ToLower(loc), Info: "429 Rate limit"}
	}

	if SupportGPT(loc) {
		return result.Result{Status: status.OK, Region: strings.ToLower(loc)}
	}
	return result.Result{Status: status.No, Region: strings.ToLower(loc)}
}
