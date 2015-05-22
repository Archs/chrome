package proxy

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	proxy = chrome.Get("proxy")
	// TODO full fill this
	Settings = proxy.Get("settings")
)

/*
* Types
 */

type ProxyServer struct {
	*js.Object
	Scheme string `js:"scheme"`
	Host   string `js:"host"`
	Port   int    `js:"port"`
}

type ProxyRules struct {
	*js.Object
	SingleProxy   ProxyServer `js:"singleProxy"`
	ProxyForHttp  ProxyServer `js:"proxyForHttp"`
	ProxyForHttps ProxyServer `js:"proxyForHttps"`
	ProxyForFtp   ProxyServer `js:"proxyForFtp"`
	FallbackProxy ProxyServer `js:"fallbackProxy"`
	BypassList    []string    `js:"bypassList"`
}

type PacScript struct {
	*js.Object
	Url       string `js:"url"`
	Data      string `js:"data"`
	Mandatory bool   `js:"mandatory"`
}

type ProxyConfig struct {
	*js.Object
	Rules     ProxyRules `js:"rules"`
	PacScript PacScript  `js:"pacScript"`
	Mode      string     `js:"mode"`
}

/*
* Events
 */

// OnProxyError notifies about proxy errors.
func OnProxyError(callback func(details js.M)) {
	proxy.Get("onProxyError").Call("addListener", callback)
}
