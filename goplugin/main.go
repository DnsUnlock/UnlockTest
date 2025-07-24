package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"
	"syscall"

	"github.com/DnsUnlock/UnlockTest/lib/client"
	"github.com/DnsUnlock/UnlockTest/lib/dialer"
	"github.com/DnsUnlock/UnlockTest/lib/proxy"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/transport"
	"github.com/DnsUnlock/UnlockTest/testUnlock"
	"github.com/gophertool/tool/plugin"
)

// UnlockTestPlugin 实现了 ToolPluginInterface 接口
// 提供解锁测试相关的工具功能
type UnlockTestPlugin struct{}

var ToolList = map[string]string{ // name -> function
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

// GetPluginInfo 返回插件的基本信息
// 包括插件名称、版本和描述等元数据
func (p *UnlockTestPlugin) GetPluginInfo() (plugin.PluginInfo, error) {
	return plugin.PluginInfo{
		Name:        "UnlockTest",
		Version:     "1.1.2",
		Description: "提供流媒体解锁测试功能的插件",
		Author:      "DnsUnlock Team",
	}, nil
}

type Options struct {
	Mode       int
	Interface  string
	DnsServers string
	HttpProxy  string
	ClientIP   string
}

var SetSocketOptions = func(network, address string, c syscall.RawConn, interfaceName string) (err error) {
	return
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

// GetTools 返回插件提供的所有工具定义
// 定义了可用的工具及其输入参数模式
func (p *UnlockTestPlugin) GetTools() ([]plugin.Tool, error) {
	tools := []plugin.Tool{
		*plugin.NewTool(
			"batch_test",
			"批量测试多个流媒体平台",
			plugin.WithArray("platforms", plugin.Description("要测试的平台列表"), plugin.WithStringItems()),
			plugin.WithInteger("mode", plugin.Description("0: 自动, 4: IPv4, 6: IPv6"), plugin.Default(0)),
			plugin.WithString("interface", plugin.Description("网卡接口名称"), plugin.Default("")),
			plugin.WithString("dns_servers", plugin.Description("DNS 服务器地址"), plugin.Default("")),
			plugin.WithString("http_proxy", plugin.Description("HTTP 代理地址"), plugin.Default("")),
			plugin.WithString("client_ip", plugin.Description("目标 IP 地址"), plugin.Default(""))),
	}
	for name := range ToolList {
		tools = append(tools, *plugin.NewTool(name, fmt.Sprintf("测试 %s 解锁状态", name),
			plugin.WithInteger("mode", plugin.Description("0: 自动, 4: IPv4, 6: IPv6"), plugin.Default(0)),
			plugin.WithString("interface", plugin.Description("网卡接口名称"), plugin.Default("")),
			plugin.WithString("dns_servers", plugin.Description("DNS 服务器地址"), plugin.Default("")),
			plugin.WithString("http_proxy", plugin.Description("HTTP 代理地址"), plugin.Default("")),
			plugin.WithString("client_ip", plugin.Description("目标 IP 地址"), plugin.Default("")),
		))
	}
	return tools, nil
}

// OptionsInit 初始化Options结构体
// 从参数映射中解析配置选项，只有类型断言错误时才返回error
// 参数不存在时使用默认值，不会返回错误
func OptionsInit(params map[string]any) (*Options, error) {
	o := &Options{}

	// 处理mode参数
	if modeVal, exists := params["mode"]; exists {
		if mode, ok := modeVal.(int); ok {
			o.Mode = mode
		} else {
			return nil, fmt.Errorf("mode parameter must be an integer, got %T", modeVal)
		}
	} else {
		o.Mode = 0 // 默认值
	}

	// 处理interface参数
	if ifaceVal, exists := params["interface"]; exists {
		if iface, ok := ifaceVal.(string); ok {
			o.Interface = iface
		} else {
			return nil, fmt.Errorf("interface parameter must be a string, got %T", ifaceVal)
		}
	} else {
		o.Interface = "" // 默认值
	}

	// 处理dns_servers参数
	if dnsVal, exists := params["dns_servers"]; exists {
		if dns, ok := dnsVal.(string); ok {
			o.DnsServers = dns
		} else {
			return nil, fmt.Errorf("dns_servers parameter must be a string, got %T", dnsVal)
		}
	} else {
		o.DnsServers = "" // 默认值
	}

	// 处理http_proxy参数
	if proxyVal, exists := params["http_proxy"]; exists {
		if proxy, ok := proxyVal.(string); ok {
			o.HttpProxy = proxy
		} else {
			return nil, fmt.Errorf("http_proxy parameter must be a string, got %T", proxyVal)
		}
	} else {
		o.HttpProxy = "" // 默认值
	}

	// 处理client_ip参数
	if clientIPVal, exists := params["client_ip"]; exists {
		if clientIP, ok := clientIPVal.(string); ok {
			o.ClientIP = clientIP
		} else {
			return nil, fmt.Errorf("client_ip parameter must be a string, got %T", clientIPVal)
		}
	} else {
		o.ClientIP = "" // 默认值
	}

	return o, nil
}

// CallTool 调用指定的工具并返回结果
// 根据工具名称和参数执行相应的解锁测试功能
func (p *UnlockTestPlugin) CallTool(toolName string, params map[string]any) (*plugin.CallToolResult, error) {
	// 使用OptionsInit初始化配置
	o, err := OptionsInit(params)
	if err != nil {
		return plugin.NewErrorResult(err.Error()), nil
	}

	switch toolName {
	case "batch_test":
		var platforms []any
		if platform, exists := params["platforms"]; exists {
			if p, ok := platform.([]any); ok {
				platforms = p
			} else {
				return nil, fmt.Errorf("platforms parameter must be an array, got %T", platform)
			}
		}
		platformsStr := make([]string, 0, len(platforms))
		for _, p := range platforms {
			platformsStr = append(platformsStr, fmt.Sprintf("%v", p))
		}
		return p.batchTest(platformsStr, o)
	default:
		result, err := p.Call(toolName, o)
		if err != nil {
			return plugin.NewErrorResult(err.Error()), nil
		}
		result.SetStatusText()
		return plugin.NewCallToolResult().AddStructContent(result, toolName), nil
	}
}

func TS(result result.Result) string {
	return result.ToString()
}

// batchTest 批量测试多个平台，使用 5 个并发线程
func (p *UnlockTestPlugin) batchTest(platforms []string, o *Options) (*plugin.CallToolResult, error) {
	// 创建结果map和等待组
	var resultmp = make(map[string]result.Result)
	var wg sync.WaitGroup

	// 创建信号量来限制并发数为 5
	semaphore := make(chan struct{}, 5)

	// 启动并发测试
	for _, platform := range platforms {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()

			// 获取信号量
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// 调用测试函数
			testResult, err := p.Call(name, o)
			if err != nil {
				resultmp[name] = result.Result{
					Status: status.Err,
					Info:   err.Error(),
				}
				return
			}

			// 格式化结果
			testResult.SetStatusText()
			resultmp[name] = testResult
		}(platform)
	}

	// 等待所有测试完成
	wg.Wait()
	// 收集结果
	callResult := plugin.NewCallToolResult()
	callResult.AddTextContent(fmt.Sprintf("批量测试完成，共测试 %d 个平台:", len(resultmp)))
	for name, result := range resultmp {
		callResult.AddStructContent(result, name)
	}

	return callResult, nil
}

// Call 调用函数
// Call 执行指定的解锁测试函数并返回结果
// Function: 要执行的测试函数名称
// o: 包含客户端配置的选项
// 返回: result.Result 类型的测试结果
func (p *UnlockTestPlugin) Call(Function string, o *Options) (result.Result, error) {
	c := o.Client()

	switch Function {
	case "AcornTV":
		return testUnlock.AcornTV(c), nil
	case "BahamutAnime":
		return testUnlock.BahamutAnime(c), nil
	case "MoviStarPlus":
		return testUnlock.MoviStarPlus(c), nil
	case "ABCiView":
		return testUnlock.ABCiView(c), nil
	case "CBCGem":
		return testUnlock.CBCGem(c), nil
	case "Joyn":
		return testUnlock.Joyn(c), nil
	case "Stan":
		return testUnlock.Stan(c), nil
	case "Watcha":
		return testUnlock.Watcha(c), nil
	case "DAnimeStore":
		return testUnlock.DAnimeStore(c), nil
	case "PlutoTV":
		return testUnlock.PlutoTV(c), nil
	case "TVBAnywhere":
		return testUnlock.TVBAnywhere(c), nil
	case "Funimation":
		return testUnlock.Funimation(c), nil
	case "NicoNico":
		return testUnlock.Niconico(c), nil
	case "TVer":
		return testUnlock.TVer(c), nil
	case "Tving":
		return testUnlock.Tving(c), nil
	case "DirecTVStream":
		return testUnlock.DirectvStream(c), nil
	case "PrettyDerby":
		return testUnlock.PrettyDerbyJP(c), nil
	case "StarPlus":
		return testUnlock.StarPlus(c), nil
	case "RakutenTVEU":
		return testUnlock.RakutenTV_EU(c), nil
	case "RakutenTVJP":
		return testUnlock.RakutenTV_JP(c), nil
	case "TLCGO":
		return testUnlock.TlcGo(c), nil
	case "Wavve":
		return testUnlock.Wavve(c), nil
	case "CanalPlus":
		return testUnlock.CanalPlus(c), nil
	case "CoupangPlay":
		return testUnlock.CoupangPlay(c), nil
	case "HamiVideo":
		return testUnlock.HamiVideo(c), nil
	case "J_COM_ON_DEMAND":
		return testUnlock.J_COM_ON_DEMAND(c), nil
	case "KayoSports":
		return testUnlock.KayoSports(c), nil
	case "SHOWTIME":
		return testUnlock.SHOWTIME(c), nil
	case "Abema":
		return testUnlock.Abema(c), nil
	case "BiliBiliID":
		return testUnlock.BilibiliID(c), nil
	case "BiliBiliVN":
		return testUnlock.BilibiliVN(c), nil
	case "BiliBiliHKMO":
		return testUnlock.BilibiliHKMO(c), nil
	case "BiliBiliTW":
		return testUnlock.BilibiliTW(c), nil
	case "BiliBiliSEA":
		return testUnlock.BilibiliSEA(c), nil
	case "BiliBiliTH":
		return testUnlock.BilibiliTH(c), nil
	case "HBOGoAsia":
		return testUnlock.HboGoAsia(c), nil
	case "MeWatch":
		return testUnlock.MeWatch(c), nil
	case "MyTVSuper":
		return testUnlock.MyTvSuper(c), nil
	case "FOD":
		return testUnlock.FOD(c), nil
	case "NowE":
		return testUnlock.NowE(c), nil
	case "ZDF":
		return testUnlock.ZDF(c), nil
	case "NLZIET":
		return testUnlock.NLZIET(c), nil
	case "NaverTV":
		return testUnlock.NaverTV(c), nil
	case "PJSK":
		return testUnlock.PJSK(c), nil
	case "UNext":
		return testUnlock.U_NEXT(c), nil
	case "VideoMarket":
		return testUnlock.VideoMarket(c), nil
	case "Bing":
		return testUnlock.Bing(c), nil
	case "Molotov":
		return testUnlock.Molotov(c), nil
	case "TikTok":
		return testUnlock.TikTok(c), nil
	case "TubiTV":
		return testUnlock.TubiTV(c), nil
	case "Wikipedia":
		return testUnlock.WikipediaEditable(c), nil
	case "Epix":
		return testUnlock.Epix(c), nil
	case "HBOMax":
		return testUnlock.HBOMax(c), nil
	case "KKTV":
		return testUnlock.KKTV(c), nil
	case "Catchplay":
		return testUnlock.Catchplay(c), nil
	case "SkyShowTime":
		return testUnlock.SkyShowTime(c), nil
	case "SlingTV":
		return testUnlock.SlingTV(c), nil
	case "StarZ":
		return testUnlock.Starz(c), nil
	case "Viu":
		return testUnlock.ViuCom(c), nil
	case "Afreeca":
		return testUnlock.Afreeca(c), nil
	case "Channel10":
		return testUnlock.Channel10(c), nil
	case "MaoriTV":
		return testUnlock.MaoriTV(c), nil
	case "Radiko":
		return testUnlock.Radiko(c), nil
	case "iQiYi":
		return testUnlock.IQiYi(c), nil
	case "ThreeNow":
		return testUnlock.ThreeNow(c), nil
	case "BritBox":
		return testUnlock.BritBox(c), nil
	case "CWTV":
		return testUnlock.CW_TV(c), nil
	case "DiscoveryPlus":
		return testUnlock.DiscoveryPlus(c), nil
	case "DocPlay":
		return testUnlock.DocPlay(c), nil
	case "NFLPlus":
		return testUnlock.NFLPlus(c), nil
	case "DisneyPlus":
		return testUnlock.DisneyPlus(c), nil
	case "Lemino":
		return testUnlock.Lemino(c), nil
	case "PrimeVideo":
		return testUnlock.PrimeVideo(c), nil
	case "Shudder":
		return testUnlock.Shudder(c), nil
	case "Spotify":
		return testUnlock.Spotify(c), nil
	case "EurosportRO":
		return testUnlock.EurosportRO(c), nil
	case "ITVX":
		return testUnlock.ITVX(c), nil
	case "MyVideo":
		return testUnlock.MyVideo(c), nil
	case "SkyGo":
		return testUnlock.SkyGo(c), nil
	case "SkyGoNZ":
		return testUnlock.SkyGo_NZ(c), nil
	case "VideoLand":
		return testUnlock.VideoLand(c), nil
	case "AISPlay":
		return testUnlock.AISPlay(c), nil
	case "Popcornflix":
		return testUnlock.Popcornflix(c), nil
	case "Reddit":
		return testUnlock.Reddit(c), nil
	case "SetantaSports":
		return testUnlock.SetantaSports(c), nil
	case "Crunchyroll":
		return testUnlock.Crunchyroll(c), nil
	case "FXNOW":
		return testUnlock.FXNOW(c), nil
	case "Fox":
		return testUnlock.Fox(c), nil
	case "TrueID":
		return testUnlock.TrueID(c), nil
	case "ViuTV":
		return testUnlock.ViuTV(c), nil
	case "YouTubeRegion":
		return testUnlock.YoutubeRegion(c), nil
	case "YouTubeCDN":
		return testUnlock.YoutubeCDN(c), nil
	case "Amediateka":
		return testUnlock.Amediateka(c), nil
	case "Binge":
		return testUnlock.Binge(c), nil
	case "DirecTVGO":
		return testUnlock.DirecTVGO(c), nil
	case "LineTV":
		return testUnlock.LineTV(c), nil
	case "RaiPlay":
		return testUnlock.RaiPlay(c), nil
	case "FuboTV":
		return testUnlock.FuboTV(c), nil
	case "HotStar":
		return testUnlock.Hotstar(c), nil
	case "MusicJP":
		return testUnlock.MusicJP(c), nil
	case "KonosubaFD":
		return testUnlock.KonosubaFD(c), nil
	case "NetflixRegion":
		return testUnlock.NetflixRegion(c), nil
	case "NetflixCDN":
		return testUnlock.NetflixCDN(c), nil
	case "Wowow":
		return testUnlock.Wowow(c), nil
	case "Crave":
		return testUnlock.Crave(c), nil
	case "Kancolle":
		return testUnlock.Kancolle(c), nil
	case "Showmax":
		return testUnlock.Showmax(c), nil
	case "Telasa":
		return testUnlock.Telasa(c), nil
	case "BBCiPlayer":
		return testUnlock.BBCiPlayer(c), nil
	case "ESPNPlus":
		return testUnlock.ESPNPlus(c), nil
	case "GYAO":
		return testUnlock.GYAO(c), nil
	case "KBS":
		return testUnlock.KBS(c), nil
	case "SpotvNow":
		return testUnlock.SpotvNow(c), nil
	case "Channel5":
		return testUnlock.Channel5(c), nil
	case "PeacockTV":
		return testUnlock.PeacockTV(c), nil
	case "Philo":
		return testUnlock.Philo(c), nil
	case "WFJP":
		return testUnlock.WFJP(c), nil
	case "OptusSports":
		return testUnlock.OptusSports(c), nil
	case "PCRJP":
		return testUnlock.PCRJP(c), nil
	case "DMM":
		return testUnlock.DMM(c), nil
	case "DMMTV":
		return testUnlock.DMMTV(c), nil
	case "DSTV":
		return testUnlock.DSTV(c), nil
	case "HuluJP":
		return testUnlock.HuluJP(c), nil
	case "Hulu":
		return testUnlock.Hulu(c), nil
	case "Karaoke":
		return testUnlock.Karaoke(c), nil
	case "NBATV":
		return testUnlock.NBA_TV(c), nil
	case "Dazn":
		return testUnlock.Dazn(c), nil
	case "Steam":
		return testUnlock.Steam(c), nil
	case "Paramount+":
		return testUnlock.ParamountPlus(c), nil
	case "Mora":
		return testUnlock.Mora(c), nil
	case "SBSonDemand":
		return testUnlock.SBSonDemand(c), nil
	case "Channel4":
		return testUnlock.Channel4(c), nil
	case "ChatGPT":
		return testUnlock.ChatGPT(c), nil
	case "EncoreTVB":
		return testUnlock.EncoreTVB(c), nil
	case "Paravi":
		return testUnlock.Paravi(c), nil
	case "4GTV":
		return testUnlock.TW4GTV(c), nil
	case "7Plus":
		return testUnlock.SevenPlus(c), nil
	case "Instagram":
		return testUnlock.Instagram(c), nil
	case "LiTV":
		return testUnlock.LiTV(c), nil
	case "NeonTV":
		return testUnlock.NeonTV(c), nil
	case "Channel9":
		return testUnlock.Channel9(c), nil
	case "NPOStartPlus":
		return testUnlock.NPOStartPlus(c), nil
	case "SonyLiv":
		return testUnlock.SonyLiv(c), nil

	default:
		// 返回一个表示未找到函数的错误结果
		return result.Result{}, errors.New("function not found")
	}
}

// main 函数是插件的入口点
// 创建插件实例并启动插件服务器
func main() {
	// 创建插件实例
	pluginImpl := &UnlockTestPlugin{}

	// 启动插件服务器
	// 这会阻塞当前 goroutine 直到插件被关闭
	plugin.ServePlugin(pluginImpl)
}
