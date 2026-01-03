package proxy

import "net/http"

var Client = http.ProxyFromEnvironment
