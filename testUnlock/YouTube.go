package testUnlock

import (
	"bufio"
	"github.com/DnsUnlock/UnlockTest/lib/result"
	"github.com/DnsUnlock/UnlockTest/lib/status"
	"github.com/DnsUnlock/UnlockTest/lib/url"
	"io"
	"log"
	"net/http"
	"strings"
)

func YoutubeRegion(c http.Client) result.Result {
	resp, err := url.GET(c, "https://www.youtube.com/premium", url.H{"Cookie", "YSC=BiCUU3-5Gdk; CONSENT=YES+cb.20220301-11-p0.en+FX+700; GPS=1; VISITOR_INFO1_LIVE=4VwPMkB7W5A; SOCS=CAISOAgDEitib3FfaWRlbnRpdHlmcm9udGVuZHVpc2VydmVyXzIwMjQwNTIxLjA3X3AxGgV6aC1DTiACGgYIgNTEsgY; PREF=f7=4000&tz=Asia.Shanghai&f4=4000000; _gcl_au=1.1.1809531354.1646633279"})
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "www.google.cn") {
		return result.Result{Status: status.No, Region: "cn"}
	}
	if strings.Contains(s, "Premium is not available in your country") {
		return result.Result{Status: status.No}
	}
	if EndLocation := strings.Index(s, `"countryCode":`); EndLocation != -1 {
		return result.Result{
			Status: status.OK,
			Region: strings.ToLower(s[EndLocation+15 : EndLocation+17]),
		}
	}
	if strings.Contains(s, "premiumPurchaseButton") || strings.Contains(s, "manageSubscriptionButton") || strings.Contains(s, "/月") || strings.Contains(s, "/month") {
		return result.Result{Status: status.OK}
	}
	return result.Result{Status: status.No}
}

func YoutubeCDN(c http.Client) result.Result {
	resp, err := url.GET(c, "https://redirector.googlevideo.com/report_mapping")
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	r := bufio.NewReader(resp.Body)
	b, _, err := r.ReadLine()
	if err != nil {
		return result.Result{Status: status.NetworkErr, Err: err}
	}
	s := string(b)
	i := strings.Index(s, "=> ")
	if i == -1 {
		return result.Result{Status: status.Unexpected}
	}
	s = s[i+3:]
	i = strings.Index(s, " ")
	if i == -1 {
		return result.Result{Status: status.Unexpected}
	}
	s = s[:i]
	i = strings.Index(s, "-")

	if i == -1 {
		i = strings.Index(s, ".")
		return result.Result{
			Status: status.OK,
			Region: findAirCode(s[i+1:]),
			Info:   "Youtube Video Server",
		}
	} else {
		isp := s[:i]
		return result.Result{
			Status: status.OK,
			Region: isp + " - " + findAirCode(s[i+1:]),
			Info:   "Google Global CacheCDN (ISP Cooperation)",
		}
	}
}

func findAirCode(code string) string {
	airPortCode := []string{"KIX", "NRT", "GMP", "YOW", "YMQ/YUL", "YVR", "YYC", "YEG", "YTO/YYZ", "WAS/IAD", "ABE", "ABQ", "ATL", "AUS", "AZO", "BDL", "BHM", "BNA", "BOI", "BOS", "BRO", "BTR", "BTL", "BUF", "BWI", "CAE", "CAK", "CHA", "CHI/ORD", "CHS", "CID", "CLE", "CLT", "CMH", "CRP", "CVG", "DAY", "DEN", "DFW", "DSM", "DTW", "ELP", "ERI", "EWR", "EVV", "FLL", "FNT", "FWA", "GRR", "GEG", "GSO", "GSP", "GRB", "HAR", "HOU/IAH", "HSV", "HNL", "ICT", "ILM", "IND", "JAN", "JAX", "LAS", "LAX", "LEX", "LIT", "LNK", "LRD", "MCI", "MCO", "MEM", "MFE", "MIA", "MKC", "MKE", "MSN", "MSP", "MSY", "MOB", "NYC/JFK", "OKC", "OMA", "ORF", "ORL", "PBI", "PDX", "PHL/PHA", "PHX", "PIA", "PIT", "PNS", "PVD", "RDU", "RIC", "RNO", "ROC", "SAN", "SAT", "SAV", "SBN", "SDF", "SEA", "BFI", "SFO", "SGF", "SHV", "SLC", "SMF", "STL", "TUL", "SYR", "TOL", "TPA", "TUL", "TUS", "TYS", "MEX", "GDL/MEX", "GUA", "TGU", "SAL", "MGA", "SJO", "PTY", "NAS", "HAV", "SCU", "KIN", "PAP", "SDQ", "SJU", "ROX", "GND", "BGI", "POS", "BOG", "CCS", "GEO", "PBM", "CAY", "BSB", "CWB", "POA", "MAO", "RIO", "SAO", "UIO", "GYE", "LIM", "SRE", "ASU", "MVD", "BUE", "ANF", "SCL", "PTP", "LON/LHR", "ABZ", "BHX", "BOH", "BRS", "CWL", "EDI", "EXT", "GLA", "LPL", "MAN", "NWI", "PLH", "SOU", "BRS", "CDQ", "CVT", "LBA", "PME", "NCL", "HUY", "PIK", "EMA", "BFS", "DUB", "ORK", "SNN", "BRU", "ANR", "OST", "LUX", "AMS", "RTM", "EIN", "ENS", "CPH", "ALL", "AAR", "BLL", "BER/TXL", "MUC", "BRE", "HAJ", "DUS", "FRA", "LEJ", "DUI", "STR", "HAM", "ERF", "FMO", "NUE", "DRS", "SCN", "CGN", "DTM", "BFE", "ZTZ", "ESS", "BON", "RUN", "PAR/CDG", "MRS", "LYS", "BOD", "LIL", "TLS", "NTE", "MLH", "MPL", "GNB", "URO", "NCE", "SXB", "XVE", "PPT", "XMM/GRZ", "BRN", "GVA", "ZRH", "BSL", "ALV", "MAD", "ALC", "BCN", "VLC", "SVQ", "AGP", "VLL", "LIS", "OPO", "ROM", "AHO", "AOI", "BDS", "BLQ", "BRI", "GOA", "MIL/MXP", "SWK", "NAP", "VCE", "FLR", "TRN", "TRS", "CTA", "TAR", "PSA", "QME", "VRN", "ATH", "SKG", "VIE", "LNZ", "GRZ", "SZG", "INN", "PRG", "HEL", "STO/ARN", "AGH", "GOT", "MMA/MMX", "NRK", "OSL", "TIA", "SKP", "SOF", "BEG", "BUH", "KIV", "ZAG", "LJU", "BUD", "BTS", "WAW", "KRK", "GDN", "VNO", "RIX", "TLL", "REK", "MOW", "LED", "MSQ", "IEV/KBP", "SJJ", "THR", "ABD", "KBL", "KWI", "RUH", "JED", "DMM", "SAH", "ADE", "BGW", "BEY", "BAH", "AUH", "DXB", "SHJ", "DOH", "JRD", "TLV", "DAM", "AMM", "ANK", "ADA", "BTZ", "IZM", "IST", "BAH", "NIC", "LCA", "BAK", "EVN", "TBS", "MSH", "ASB", "DYU", "KGF", "FRU", "TAS", "CAI", "KRT", "MCT", "ADD", "JIB", "NBO", "TIP", "ALG", "AAE", "TUN", "RBA", "CAS", "NDJ", "NIM", "ABV", "LOS", "PHC", "BKO", "OUA", "COO", "LFW", "ACC", "ASK", "ABJ", "HGS", "MLW", "CKF", "DKR", "BJL", "KLA", "BGF", "YAO", "SSG", "KLA", "KGL", "DAR", "BJM", "BZV", "LBV", "TMS", "MPM", "LLW", "LUN", "HRE", "LAD", "GBE", "WDH", "JNB", "DUR", "CPT", "MRU", "TNR", "YVA", "SEZ", "NKC", "HKG", "TPE", "KHH", "FNJ", "SEL/ICN", "PUS", "TYONRT", "KIX/OSA", "NGO", "FUK", "YOK", "HIJ", "OKA", "SDJ", "SPA", "MNL", "HEB", "DVO", "KUL", "PEN", "LGK", "BKI", "KCH", "IPH", "JHB", "KBR", "SBW", "SDK", "BWN", "SIN", "JKT", "MES", "SUB", "DPS", "UPG", "PNK", "DIL", "SGN", "HAN", "HPH", "VTE", "BKK", "CEI", "HDY", "HKT", "NSI", "RGN", "MDL", "PNH", "DAC", "CGP", "DEL", "BOM", "CCU", "MAA", "BLR", "SXM", "HYD", "KTM", "ISB", "KHI", "LHE", "PEW", "CMB", "MLE", "ULN", "CBR", "MEL", "ADL", "DRN", "CNS", "BNE", "PER", "SYD", "WLG", "AKL", "CHC", "POM", "SUV", "TRW", "HIR", "TBU", "APW", "FUN", "KSA", "VLI"}
	i, v := 0, ""
	for ; i < len(code); i++ {
		if code[i] >= '0' && code[i] <= '9' {
			break
		}
	}
	code = strings.ToUpper(code[:i])
	for i, v = range airPortCode {
		if strings.Contains(code, v) {
			return v
			// return airPortCode[i]
		}
	}
	return code
}
