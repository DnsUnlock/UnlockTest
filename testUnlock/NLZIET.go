package testUnlock

import (
	errs "github.com/DnsUnlock/UnlockTest/lib/err"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"net/http"
	"strings"
)

func SupportNLZIET(loc string) bool {
	var NLZIET_SUPPORT_COUNTRY = []string{
		"BE", "BG", "CZ", "DK", "DE", "EE", "IE", "EL", "ES", "FR", "HR", "IT", "CY", "LV", "LT", "LU", "HU", "MT", "NL", "AT", "PL", "PT", "RO", "SI", "SK", "FI", "SE",
	}
	for _, s := range NLZIET_SUPPORT_COUNTRY {
		if loc == s {
			return true
		}
	}
	return false
}

func NLZIET(c http.Client) result.Result {
	resp, err := urls.GET(c, "https://nlziet.nl/cdn-cgi/trace")
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

	if SupportNLZIET(loc) {
		return result.Result{Status: status.OK, Region: strings.ToLower(loc)}
	}
	return result.Result{Status: status.No}
}
