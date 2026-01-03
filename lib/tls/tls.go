package tls

import (
	"crypto/tls"
	utls "github.com/refraction-networking/utls"
)

var defaultCipherSuites = []uint16{0xc02f, 0xc030, 0xc02b, 0xc02c, 0xcca8, 0xcca9, 0xc013, 0xc009, 0xc014, 0xc00a, 0x009c, 0x009d, 0x002f, 0x0035, 0xc012, 0x000a}

var c, _ = utls.UTLSIdToSpec(utls.HelloChrome_Auto)

var Config = &tls.Config{
	InsecureSkipVerify: true,
	MinVersion:         c.TLSVersMin,
	MaxVersion:         c.TLSVersMax,
	CipherSuites:       c.CipherSuites,
	ClientSessionCache: tls.NewLRUClientSessionCache(32),
}
