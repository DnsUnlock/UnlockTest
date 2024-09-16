package main

import (
	"context"
	"flag"
	"github.com/DnsUnlock/UnlockTest/lib/client"
	"github.com/DnsUnlock/UnlockTest/lib/dialer"
	"github.com/DnsUnlock/UnlockTest/lib/proxy"
	results "github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/transport"
	"github.com/DnsUnlock/UnlockTest/testUnlock"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"sync"
	"syscall"
)

type PluginInterface interface {
	Registers() string                               // 用于存储插件系统的注册函数名
	Func() map[string]string                         // 用于存储函数名
	Call(Function string, args []interface{}) string // 用于调用函数
}

var PluginEntrance PluginEntranceModel

type PluginEntranceModel struct {
}

// Registers 注册函数
func (p *PluginEntranceModel) Registers() string {
	return "UnlockTest"
}

// Func 函数
func (p *PluginEntranceModel) Func() map[string]string {
	return map[string]string{ // name -> function
		"Fox":             "Fox",
		"MyTVSuper":       "MyTvSuper",
		"NeonTV":          "NeonTV",
		"PrimeVideo":      "PrimeVideo",
		"Channel9":        "Channel9",
		"DAnimeStore":     "DAnimeStore",
		"ESPNPlus":        "ESPNPlus",
		"Epix":            "Epix",
		"FOD":             "FOD",
		"NetflixRegion":   "NetflixRegion",
		"NetflixCDN":      "NetflixCDN",
		"AISPlay":         "AISPlay",
		"Bing":            "Bing",
		"DSTV":            "DSTV",
		"AcornTV":         "AcornTV",
		"EurosportRO":     "EurosportRO",
		"PeacockTV":       "PeacockTV",
		"Catchplay":       "Catchplay",
		"Binge":           "Binge",
		"MeWatch":         "MeWatch",
		"Radiko":          "Radiko",
		"SkyGo":           "SkyGo",
		"SkyGoNZ":         "SkyGo_NZ",
		"VideoMarket":     "VideoMarket",
		"BiliBiliHKMO":    "BilibiliHKMO",
		"BiliBiliTW":      "BilibiliTW",
		"BiliBiliSEA":     "BilibiliSEA",
		"BiliBiliTH":      "BilibiliTH",
		"BiliBiliID":      "BilibiliID",
		"BiliBiliVN":      "BilibiliVN",
		"TVBAnywhere":     "TVBAnywhere",
		"DirecTVStream":   "DirectvStream",
		"DisneyPlus":      "DisneyPlus",
		"NPOStartPlus":    "NPOStartPlus",
		"BBCiPlayer":      "BBCiPlayer",
		"Instagram":       "Instagram",
		"J_COM_ON_DEMAND": "J_COM_ON_DEMAND",
		"KonosubaFD":      "KonosubaFD",
		"NaverTV":         "NaverTV",
		"Shudder":         "Shudder",
		"Watcha":          "Watcha",
		"CWTV":            "CW_TV",
		"ZDF":             "ZDF",
		"MusicJP":         "MusicJP",
		"NBATV":           "NBA_TV",
		"Crunchyroll":     "Crunchyroll",
		"Popcornflix":     "Popcornflix",
		"SlingTV":         "SlingTV",
		"ThreeNow":        "ThreeNow",
		"MoviStarPlus":    "MoviStarPlus",
		"GYAO":            "GYAO",
		"PCRJP":           "PCRJP",
		"SHOWTIME":        "SHOWTIME",
		"BritBox":         "BritBox",
		"Channel4":        "Channel4",
		"DiscoveryPlus":   "DiscoveryPlus",
		"iQiYi":           "IQiYi",
		"CanalPlus":       "CanalPlus",
		"HamiVideo":       "HamiVideo",
		"HotStar":         "Hotstar",
		"KayoSports":      "KayoSports",
		"MyVideo":         "MyVideo",
		"PlutoTV":         "PlutoTV",
		"TikTok":          "TikTok",
		"Channel10":       "Channel10",
		"RakutenTVEU":     "RakutenTV_EU",
		"RakutenTVJP":     "RakutenTV_JP",
		"SonyLiv":         "SonyLiv",
		"BahamutAnime":    "BahamutAnime",
		"Mora":            "Mora",
		"DMMTV":           "DMMTV",
		"DMM":             "DMM",
		"OptusSports":     "OptusSports",
		"RaiPlay":         "RaiPlay",
		"Spotify":         "Spotify",
		"SpotvNow":        "SpotvNow",
		"Steam":           "Steam",
		"Wavve":           "Wavve",
		"Channel5":        "Channel5",
		"Tving":           "Tving",
		"VideoLand":       "VideoLand",
		"TVer":            "TVer",
		"Afreeca":         "Afreeca",
		"Reddit":          "Reddit",
		"4GTV":            "TW4GTV",
		"CoupangPlay":     "CoupangPlay",
		"NicoNico":        "Niconico",
		"TubiTV":          "TubiTV",
		"ChatGPT":         "ChatGPT",
		"Karaoke":         "Karaoke",
		"NFLPlus":         "NFLPlus",
		"PJSK":            "PJSK",
		"SetantaSports":   "SetantaSports",
		"ABCiView":        "ABCiView",
		"ITVX":            "ITVX",
		"StarPlus":        "StarPlus",
		"7Plus":           "SevenPlus",
		"LiTV":            "LiTV",
		"Wowow":           "Wowow",
		"HBOMax":          "HBOMax",
		"NLZIET":          "NLZIET",
		"Philo":           "Philo",
		"KBS":             "KBS",
		"NowE":            "NowE",
		"Joyn":            "Joyn",
		"Crave":           "Crave",
		"HuluJP":          "HuluJP",
		"Hulu":            "Hulu",
		"Lemino":          "Lemino",
		"Molotov":         "Molotov",
		"WFJP":            "WFJP",
		"Abema":           "Abema",
		"FXNOW":           "FXNOW",
		"EncoreTVB":       "EncoreTVB",
		"HBOGoAsia":       "HboGoAsia",
		"Paramount+":      "ParamountPlus",
		"Paravi":          "Paravi",
		"Stan":            "Stan",
		"DocPlay":         "DocPlay",
		"Telasa":          "Telasa",
		"TrueID":          "TrueID",
		"UNext":           "U_NEXT",
		"Viu":             "ViuCom",
		"YouTubeRegion":   "YoutubeRegion",
		"YouTubeCDN":      "YoutubeCDN",
		"Amediateka":      "Amediateka",
		"MaoriTV":         "MaoriTV",
		"SBSonDemand":     "SBSonDemand",
		"Showmax":         "Showmax",
		"SkyShowTime":     "SkyShowTime",
		"StarZ":           "Starz",
		"Wikipedia":       "WikipediaEditable",
		"FuboTV":          "FuboTV",
		"Dazn":            "Dazn",
		"KKTV":            "KKTV",
		"Kancolle":        "Kancolle",
		"LineTV":          "LineTV",
		"PrettyDerby":     "PrettyDerbyJP",
		"ViuTV":           "ViuTV",
		"CBCGem":          "CBCGem",
		"Funimation":      "Funimation",
		"TLCGO":           "TlcGo",
		"DirecTVGO":       "DirecTVGO",
	}
}

// Call 调用函数
func (p *PluginEntranceModel) Call(Function string, args []interface{}) string {
	o := Flag(args)
	c := o.Client()
	switch Function {
	case "AcornTV":
		return TS(testUnlock.AcornTV(c))
	case "BahamutAnime":
		return TS(testUnlock.BahamutAnime(c))
	case "MoviStarPlus":
		return TS(testUnlock.MoviStarPlus(c))
	case "ABCiView":
		return TS(testUnlock.ABCiView(c))
	case "CBCGem":
		return TS(testUnlock.CBCGem(c))
	case "Joyn":
		return TS(testUnlock.Joyn(c))
	case "Stan":
		return TS(testUnlock.Stan(c))
	case "Watcha":
		return TS(testUnlock.Watcha(c))
	case "DAnimeStore":
		return TS(testUnlock.DAnimeStore(c))
	case "PlutoTV":
		return TS(testUnlock.PlutoTV(c))
	case "TVBAnywhere":
		return TS(testUnlock.TVBAnywhere(c))
	case "Funimation":
		return TS(testUnlock.Funimation(c))
	case "NicoNico":
		return TS(testUnlock.Niconico(c))
	case "TVer":
		return TS(testUnlock.TVer(c))
	case "Tving":
		return TS(testUnlock.Tving(c))
	case "DirecTVStream":
		return TS(testUnlock.DirectvStream(c))
	case "PrettyDerby":
		return TS(testUnlock.PrettyDerbyJP(c))
	case "StarPlus":
		return TS(testUnlock.StarPlus(c))
	case "RakutenTVEU":
		return TS(testUnlock.RakutenTV_EU(c))
	case "RakutenTVJP":
		return TS(testUnlock.RakutenTV_JP(c))
	case "TLCGO":
		return TS(testUnlock.TlcGo(c))
	case "Wavve":
		return TS(testUnlock.Wavve(c))
	case "CanalPlus":
		return TS(testUnlock.CanalPlus(c))
	case "CoupangPlay":
		return TS(testUnlock.CoupangPlay(c))
	case "HamiVideo":
		return TS(testUnlock.HamiVideo(c))
	case "J_COM_ON_DEMAND":
		return TS(testUnlock.J_COM_ON_DEMAND(c))
	case "KayoSports":
		return TS(testUnlock.KayoSports(c))
	case "SHOWTIME":
		return TS(testUnlock.SHOWTIME(c))
	case "Abema":
		return TS(testUnlock.Abema(c))
	case "BiliBiliID":
		return TS(testUnlock.BilibiliID(c))
	case "BiliBiliVN":
		return TS(testUnlock.BilibiliVN(c))
	case "BiliBiliHKMO":
		return TS(testUnlock.BilibiliHKMO(c))
	case "BiliBiliTW":
		return TS(testUnlock.BilibiliTW(c))
	case "BiliBiliSEA":
		return TS(testUnlock.BilibiliSEA(c))
	case "BiliBiliTH":
		return TS(testUnlock.BilibiliTH(c))
	case "HBOGoAsia":
		return TS(testUnlock.HboGoAsia(c))
	case "MeWatch":
		return TS(testUnlock.MeWatch(c))
	case "MyTVSuper":
		return TS(testUnlock.MyTvSuper(c))
	case "FOD":
		return TS(testUnlock.FOD(c))
	case "NowE":
		return TS(testUnlock.NowE(c))
	case "ZDF":
		return TS(testUnlock.ZDF(c))
	case "NLZIET":
		return TS(testUnlock.NLZIET(c))
	case "NaverTV":
		return TS(testUnlock.NaverTV(c))
	case "PJSK":
		return TS(testUnlock.PJSK(c))
	case "UNext":
		return TS(testUnlock.U_NEXT(c))
	case "VideoMarket":
		return TS(testUnlock.VideoMarket(c))
	case "Bing":
		return TS(testUnlock.Bing(c))
	case "Molotov":
		return TS(testUnlock.Molotov(c))
	case "TikTok":
		return TS(testUnlock.TikTok(c))
	case "TubiTV":
		return TS(testUnlock.TubiTV(c))
	case "Wikipedia":
		return TS(testUnlock.WikipediaEditable(c))
	case "Epix":
		return TS(testUnlock.Epix(c))
	case "HBOMax":
		return TS(testUnlock.HBOMax(c))
	case "KKTV":
		return TS(testUnlock.KKTV(c))
	case "Catchplay":
		return TS(testUnlock.Catchplay(c))
	case "SkyShowTime":
		return TS(testUnlock.SkyShowTime(c))
	case "SlingTV":
		return TS(testUnlock.SlingTV(c))
	case "StarZ":
		return TS(testUnlock.Starz(c))
	case "Viu":
		return TS(testUnlock.ViuCom(c))
	case "Afreeca":
		return TS(testUnlock.Afreeca(c))
	case "Channel10":
		return TS(testUnlock.Channel10(c))
	case "MaoriTV":
		return TS(testUnlock.MaoriTV(c))
	case "Radiko":
		return TS(testUnlock.Radiko(c))
	case "iQiYi":
		return TS(testUnlock.IQiYi(c))
	case "ThreeNow":
		return TS(testUnlock.ThreeNow(c))
	case "BritBox":
		return TS(testUnlock.BritBox(c))
	case "CWTV":
		return TS(testUnlock.CW_TV(c))
	case "DiscoveryPlus":
		return TS(testUnlock.DiscoveryPlus(c))
	case "DocPlay":
		return TS(testUnlock.DocPlay(c))
	case "NFLPlus":
		return TS(testUnlock.NFLPlus(c))
	case "DisneyPlus":
		return TS(testUnlock.DisneyPlus(c))
	case "Lemino":
		return TS(testUnlock.Lemino(c))
	case "PrimeVideo":
		return TS(testUnlock.PrimeVideo(c))
	case "Shudder":
		return TS(testUnlock.Shudder(c))
	case "Spotify":
		return TS(testUnlock.Spotify(c))
	case "EurosportRO":
		return TS(testUnlock.EurosportRO(c))
	case "ITVX":
		return TS(testUnlock.ITVX(c))
	case "MyVideo":
		return TS(testUnlock.MyVideo(c))
	case "SkyGo":
		return TS(testUnlock.SkyGo(c))
	case "SkyGoNZ":
		return TS(testUnlock.SkyGo_NZ(c))
	case "VideoLand":
		return TS(testUnlock.VideoLand(c))
	case "AISPlay":
		return TS(testUnlock.AISPlay(c))
	case "Popcornflix":
		return TS(testUnlock.Popcornflix(c))
	case "Reddit":
		return TS(testUnlock.Reddit(c))
	case "SetantaSports":
		return TS(testUnlock.SetantaSports(c))
	case "Crunchyroll":
		return TS(testUnlock.Crunchyroll(c))
	case "FXNOW":
		return TS(testUnlock.FXNOW(c))
	case "Fox":
		return TS(testUnlock.Fox(c))
	case "TrueID":
		return TS(testUnlock.TrueID(c))
	case "ViuTV":
		return TS(testUnlock.ViuTV(c))
	case "YouTubeRegion":
		return TS(testUnlock.YoutubeRegion(c))
	case "YouTubeCDN":
		return TS(testUnlock.YoutubeCDN(c))
	case "Amediateka":
		return TS(testUnlock.Amediateka(c))
	case "Binge":
		return TS(testUnlock.Binge(c))
	case "DirecTVGO":
		return TS(testUnlock.DirecTVGO(c))
	case "LineTV":
		return TS(testUnlock.LineTV(c))
	case "RaiPlay":
		return TS(testUnlock.RaiPlay(c))
	case "FuboTV":
		return TS(testUnlock.FuboTV(c))
	case "HotStar":
		return TS(testUnlock.Hotstar(c))
	case "MusicJP":
		return TS(testUnlock.MusicJP(c))
	case "KonosubaFD":
		return TS(testUnlock.KonosubaFD(c))
	case "NetflixRegion":
		return TS(testUnlock.NetflixRegion(c))
	case "NetflixCDN":
		return TS(testUnlock.NetflixCDN(c))
	case "Wowow":
		return TS(testUnlock.Wowow(c))
	case "Crave":
		return TS(testUnlock.Crave(c))
	case "Kancolle":
		return TS(testUnlock.Kancolle(c))
	case "Showmax":
		return TS(testUnlock.Showmax(c))
	case "Telasa":
		return TS(testUnlock.Telasa(c))
	case "BBCiPlayer":
		return TS(testUnlock.BBCiPlayer(c))
	case "ESPNPlus":
		return TS(testUnlock.ESPNPlus(c))
	case "GYAO":
		return TS(testUnlock.GYAO(c))
	case "KBS":
		return TS(testUnlock.KBS(c))
	case "SpotvNow":
		return TS(testUnlock.SpotvNow(c))
	case "Channel5":
		return TS(testUnlock.Channel5(c))
	case "PeacockTV":
		return TS(testUnlock.PeacockTV(c))
	case "Philo":
		return TS(testUnlock.Philo(c))
	case "WFJP":
		return TS(testUnlock.WFJP(c))
	case "OptusSports":
		return TS(testUnlock.OptusSports(c))
	case "PCRJP":
		return TS(testUnlock.PCRJP(c))
	case "DMM":
		return TS(testUnlock.DMM(c))
	case "DMMTV":
		return TS(testUnlock.DMMTV(c))
	case "DSTV":
		return TS(testUnlock.DSTV(c))
	case "HuluJP":
		return TS(testUnlock.HuluJP(c))
	case "Hulu":
		return TS(testUnlock.Hulu(c))
	case "Karaoke":
		return TS(testUnlock.Karaoke(c))
	case "NBATV":
		return TS(testUnlock.NBA_TV(c))
	case "Dazn":
		return TS(testUnlock.Dazn(c))
	case "Steam":
		return TS(testUnlock.Steam(c))
	case "Paramount+":
		return TS(testUnlock.ParamountPlus(c))
	case "Mora":
		return TS(testUnlock.Mora(c))
	case "SBSonDemand":
		return TS(testUnlock.SBSonDemand(c))
	case "Channel4":
		return TS(testUnlock.Channel4(c))
	case "ChatGPT":
		return TS(testUnlock.ChatGPT(c))
	case "EncoreTVB":
		return TS(testUnlock.EncoreTVB(c))
	case "Paravi":
		return TS(testUnlock.Paravi(c))
	case "4GTV":
		return TS(testUnlock.TW4GTV(c))
	case "7Plus":
		return TS(testUnlock.SevenPlus(c))
	case "Instagram":
		return TS(testUnlock.Instagram(c))
	case "LiTV":
		return TS(testUnlock.LiTV(c))
	case "NeonTV":
		return TS(testUnlock.NeonTV(c))
	case "Channel9":
		return TS(testUnlock.Channel9(c))
	case "NPOStartPlus":
		return TS(testUnlock.NPOStartPlus(c))
	case "SonyLiv":
		return TS(testUnlock.SonyLiv(c))

	default:
		return "Function Not Found"
	}
}

var R []*result
var wg *sync.WaitGroup
var tot int64

type result struct {
	Name    string
	Divider bool
	Value   results.Result
}

func TS(result results.Result) string {
	return result.ToString()
}

type Options struct {
	Mode       int
	Interface  string
	DnsServers string
	HttpProxy  string
	ClientIP   string
}

func (o *Options) Client() http.Client {
	if o.Interface != "" {
		if IP := net.ParseIP(o.Interface); IP != nil {
			dialer.Dialer.LocalAddr = &net.TCPAddr{IP: IP}
		} else {
			dialer.Dialer.Control = func(network, address string, c syscall.RawConn) error {
				return SetSocketOptions(network, address, c, o.Interface)
			}
		}
	}
	if o.ClientIP != "" {
		//需要判断这里的IP为IPv4/IPv6 以及是否包含端口号
		host, port, err := net.SplitHostPort(o.ClientIP)
		if err != nil {
			// 如果没有端口号，可能会报错，所以直接将 IP 当作 host
			host = o.ClientIP
			port = "443"
		}
		// 判断是否为有效的 IPv4/IPv6
		parsedIP := net.ParseIP(host)
		if parsedIP != nil {
			dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
				return dialer.Dialer.DialContext(ctx, network, net.JoinHostPort(host, port))
			}
			transport.Auto.DialContext = dialContext
			transport.Ipv4.DialContext = dialContext
			transport.Ipv6.DialContext = dialContext
		}
	}
	if o.DnsServers != "" {
		dialer.Dialer.Resolver = &net.Resolver{
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return (&net.Dialer{}).DialContext(ctx, "udp", o.DnsServers)
			},
		}
	}
	if o.HttpProxy != "" {
		log.Println(o.HttpProxy)
		if u, err := url.Parse(o.HttpProxy); err == nil {
			proxy.Client = http.ProxyURL(u)
			transport.Ipv4.Proxy = proxy.Client
			client.Ipv4.Transport = transport.Ipv4
			transport.Ipv6.Proxy = proxy.Client
			client.Ipv6.Transport = transport.Ipv6
			transport.Auto.Proxy = proxy.Client
			client.Auto.Transport = transport.Auto
		}
	}
	switch o.Mode {
	case 4:
		return client.Ipv4
	case 6:
		return client.Ipv6
	default:
		return client.Auto
	}
}

func Flag(args []interface{}) Options {
	mode := 0
	Iface := ""
	DnsServers := ""
	httpProxy := ""
	clientIP := ""
	// Initialize the flags with default values
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flagSet.IntVar(&mode, "m", 0, "mode 0(default)/4/6")
	flagSet.StringVar(&Iface, "I", "", "source ip / interface")
	flagSet.StringVar(&DnsServers, "dns-servers", "", "specify dns servers")
	flagSet.StringVar(&httpProxy, "http-proxy", "", "http proxy")
	flagSet.StringVar(&clientIP, "client-ip", "", "client ip")
	// Parse the provided arguments
	var newArgs []string
	for _, arg := range args {
		newArgs = append(newArgs, arg.(string))
	}
	flagSet.Parse(newArgs)

	// Return the parsed values in a formatted string
	var options = Options{
		Mode:       mode,
		Interface:  Iface,
		DnsServers: DnsServers,
		HttpProxy:  httpProxy,
		ClintIP:    clientIP,
	}
	return options
}

var SetSocketOptions = func(network, address string, c syscall.RawConn, interfaceName string) (err error) {
	return
}

func init() {
	SetSocketOptions = func(network, address string, c syscall.RawConn, interfaceName string) (err error) {
		switch network {
		case "tcp", "tcp4", "tcp6":
		case "udp", "udp4", "udp6":
		default:
			return
		}
		var innerErr error
		if runtime.GOOS == "linux" {
			err = c.Control(func(fd uintptr) {
				host, _, _ := net.SplitHostPort(address)
				if ip := net.ParseIP(host); ip != nil && !ip.IsGlobalUnicast() {
					return
				}
				if innerErr = unix.BindToDevice(int(fd), interfaceName); innerErr != nil {
					return
				}
			})
		}
		if innerErr != nil {
			err = innerErr
		}
		return
	}
}
