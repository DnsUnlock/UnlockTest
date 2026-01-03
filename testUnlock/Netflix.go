package testUnlock

import (
	"encoding/json"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	urls "github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"log"
	"net/http"
	"strings"
)

func NetflixRegion(c http.Client) result.Result {
	// 70143836 绝命毒师
	// 80018499 test
	// 81280792 乐高
	resp, err := urls.GET(c, "https://www.netflix.com/title/81280792")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp2, err := urls.GET(c, "https://www.netflix.com/title/70143836")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp2.Body.Close()
	_, err = io.ReadAll(resp2.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 404 && resp2.StatusCode == 404 {
		return result.Result{Status: status.Restricted, Info: "Originals Only"}
	}
	if resp.StatusCode == 403 && resp2.StatusCode == 403 {
		return result.Result{Status: status.Banned}
	}
	if (resp.StatusCode == 200 || resp.StatusCode == 301) || (resp2.StatusCode == 200 || resp2.StatusCode == 301) {
		resp3, err := urls.GET(c, "https://www.netflix.com/title/80018499")
		if err != nil {
			return result.Result{Status: status.NetworkErr, Err: err}
		}
		defer resp3.Body.Close()
		_, err = io.ReadAll(resp3.Body)
		if err != nil {
			log.Fatal(err)
		}
		u := resp3.Header.Get("location")
		if u == "" {
			return result.Result{Status: status.OK, Region: "us"}
		}
		// log.Println("nf", u)
		t := strings.SplitN(u, "/", 5)
		if len(t) < 5 {
			return result.Result{Status: status.Unexpected}
		}
		return result.Result{Status: status.OK, Region: strings.SplitN(t[3], "-", 2)[0]}
	}
	return result.Result{Status: status.Unexpected}
}

func NetflixCDN(c http.Client) result.Result {
	resp, err := urls.GET(c, "https://api.fast.com/netflix/speedtest/v2?https=true&token=YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm&urlCount=5")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	if resp.StatusCode == 403 {
		return result.Result{
			Status: status.Banned,
			Info:   "IP Banned By Netflix",
		}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	var res netflixCdnresult
	if err := json.Unmarshal(b, &res); err != nil {
		return result.Result{Status: status.Err, Err: err}
	}
	// u, err := url.Parse(res.Targets[0].Url)
	// if err!=nil{
	// 	return result.Result{Status: , Err: err}
	// }
	// ips,err:=net.LookupHost(u.Host)
	// if err!=nil{
	// 	return result.Result{Status: , Err: err}
	// }
	return result.Result{
		Status: status.OK,
		Region: res.Targets[0].Location.Country,
	}
}

type netflixLocation struct {
	City    string
	Country string
}
type netflixCdnTarget struct {
	Name     string
	Url      string
	Location netflixLocation
}
type netflixCdnresult struct {
	Targets []netflixCdnTarget
}
