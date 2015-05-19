package net

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	tcp = js.Global.Get("chrome").Get("sockets").Get("tcp")
)

// chrome.sockets.tcp.create( SocketProperties properties, function callback)
// Creates a TCP socket.
//
// Parameters
// SocketProperties	(optional) properties
// The socket properties (optional).
//
// function	callback
// Called when the socket has been created.
//
// The callback parameter should be a function that looks like this:
//
// function(object createInfo) {...};
//    object	createInfo
//       The result of the socket creation.
//
//    integer	socketId
//        The ID of the newly created socket. Note that socket IDs created from this API are not compatible with socket IDs created from other APIs, such as the deprecated socket API.
func Create(p SocketProperties, callback func(*CreateInfo)) {
	tcp.Call("create", p, callback)
}
