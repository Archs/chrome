package net

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	sockets = js.Global.Get("chrome").Get("sockets")
)

type SocketProperties struct {
	*js.Object
	// boolean	(optional) persistent
	// Flag indicating if the socket is left open when the event page of the application is unloaded (see Manage App Lifecycle). The default value is "false." When the application is loaded, any sockets previously opened with persistent=true can be fetched with getSockets.
	Persistent bool `js:"persistent"`

	// string	(optional) name
	// An application-defined string associated with the socket.
	Name string `js:"name"`

	// integer	(optional) bufferSize
	// The size of the buffer used to receive data. The default value is 4096.
	BufferSize int `js:"bufferSize"`
}

type SocketInfo struct {
	*js.Object
	// integer	socketId
	// The socket identifier.
	SocketId int `js:"socketId"`

	// boolean	persistent
	// Flag indicating whether the socket is left open when the application is suspended (see SocketProperties.persistent).
	Persistent bool `js:"persistent"`

	// string	(optional) name
	// Application-defined string associated with the socket.
	Name string `js:"name"`

	// integer	(optional) bufferSize
	// The size of the buffer used to receive data. If no buffer size has been specified explictly, the value is not provided.
	BufferSize int `js:"bufferSize"`

	// boolean	paused
	// Flag indicating whether a connected socket blocks its peer from sending more data (see setPaused).
	Paused bool `js:"paused"`

	// boolean	connected
	// Flag indicating whether the socket is connected to a remote peer.
	Connected bool `js:"connected"`

	// string	(optional) localAddress
	// If the underlying socket is connected, contains its local IPv4/6 address.
	LocalAddress string `js:"localAddress"`

	// integer	(optional) localPort
	// If the underlying socket is connected, contains its local port.
	LocalPort int `js:"localPort"`

	// string	(optional) peerAddress
	// If the underlying socket is connected, contains the peer/ IPv4/6 address.
	PeerAddress string `js:"peerAddress"`

	// integer	(optional) peerPort
	// If the underlying socket is connected, contains the peer port.
	PeerPort int `js:"peerPort"`
}

type CreateInfo struct {
	*js.Object
	// 	integer	socketId
	// The ID of the newly created socket. Note that socket IDs created from this API are not compatible with socket IDs created from other APIs, such as the deprecated socket API.
	SocketId int `js:"socketId"`
}
