package udp

import (
	"github.com/Archs/chrome/net/sockets"
)

var (
	udp = js.Global.Get("chrome").Get("sockets").Get("udp")
)
