package tcp

import (
	"github.com/Archs/chrome/net/sockets"
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
func Create(p sockets.SocketProperties, callback func(*sockets.CreateInfo)) {
	tcp.Call("create", p, callback)
}

func CreateM(socketProperties js.M, callback func(*sockets.CreateInfo)) {
	tcp.Call("create", socketProperties, callback)
}

func CreateEx(callback func(*sockets.CreateInfo)) {
	tcp.Call("create", callback)
}

// chrome.sockets.tcp.update(integer socketId, SocketProperties properties, function callback)
// Updates the socket properties.
//
// Parameters
// integer	socketId
// The socket identifier.
//
// SocketProperties	properties
// The properties to update.
//
// function	(optional) callback
// Called when the properties are updated.
//
// If you specify the callback parameter, it should be a function that looks like this:
//
// function() {...};
func Update(socketId int, properties *sockets.SocketProperties, callback func()) {
	tcp.Call("update", socketId, properties, callback)
}

func UpdateEx(socketId int, properties *sockets.SocketProperties) {
	tcp.Call("update", socketId, properties)
}

// chrome.sockets.tcp.connect(integer socketId, string peerAddress, integer peerPort, function callback)
// Connects the socket to a remote machine. When the connect operation completes successfully, onReceive events are raised when data is received from the peer. If a network error occurs while the runtime is receiving packets, a onReceiveError event is raised, at which point no more onReceive event will be raised for this socket until the resume method is called.
//
// Parameters
// integer	socketId
// The socket identifier.
//
// string	peerAddress
// The address of the remote machine. DNS name, IPv4 and IPv6 formats are supported.
//
// integer	peerPort
// The port of the remote machine.
//
// function	callback
// Called when the connect attempt is complete.
//
// The callback parameter should be a function that looks like this:
//
// function(integer result) {...};
// integer	result
// The result code returned from the underlying network call. A negative value indicates an error.
func Connect(socketId int, peerAddress string, peerPort int, callback func(result int)) {
	tcp.Call("connect", socketId, peerAddress, peerPort, callback)
}

// secure
//
// chrome.sockets.tcp.secure(integer socketId, object options, function callback)
// Since Chrome 38.
//
// Start a TLS client connection over the connected TCP client socket.
//
// Parameters
// integer	socketId
// The existing, connected socket to use.
//
// object	(optional) options
// Constraints and parameters for the TLS connection.
//
// object	(optional) tlsVersion
// string	(optional) min
// The minimum and maximum acceptable versions of TLS. These will be tls1, tls1.1, or tls1.2.
//
// string	(optional) max
// function	callback
// Called when the connection attempt is complete.
//
// The callback parameter should be a function that looks like this:
//
// function(integer result) {...};
// integer	result
func Secure(socketId int, callback func(result int)) {
	tcp.Call("secure", js.M{
		"tlsVersion": js.M{
			"min": "tsl1",
			"max": "tsl.2",
		},
	}, callback)
}

type SendInfo struct {
	*js.Object
	// integer	resultCode
	// The result code returned from the underlying network call. A negative value indicates an error.
	ResultCode int `js:"resultCode"`
	// integer	(optional) bytesSent
	// The number of bytes sent (if result == 0)
	BytesSent int `js:"bytesSent"`
}

// send
//
// chrome.sockets.tcp.send(integer socketId, ArrayBuffer data, function callback)
// Sends data on the given TCP socket.
//
// Parameters
// integer	socketId
// The socket identifier.
//
// ArrayBuffer	data
// The data to send.
//
// function	callback
// Called when the send operation completes.
//
// The callback parameter should be a function that looks like this:
//
// function(object sendInfo) {...};
// object	sendInfo
// Result of the send method.
//
// integer	resultCode
// The result code returned from the underlying network call. A negative value indicates an error.
//
// integer	(optional) bytesSent
// The number of bytes sent (if result == 0)
func Send(socketId int, data []byte, callback func(*SendInfo)) {
	tcp.Call("send", socketId, js.NewArrayBuffer(data), callback)
}

type ReceiveInfo struct {
	*js.Object
	// integer	socketId
	// The socket identifier.
	SocketId int `js:"socketId"`
	// ArrayBuffer	data
	// The data received, with a maxium size of bufferSize.
	dat  *js.Object `js:"data"`
	Data []byte
}

// onReceive
//
// Event raised when data has been received for a given socket.
//
// addListener
// chrome.sockets.tcp.onReceive.addListener(function callback)
// Parameters
// function	callback
// The callback parameter should be a function that looks like this:
//
// function(object info) {...};
// object	info
// The event data.
//
// integer	socketId
// The socket identifier.
//
// ArrayBuffer	data
// The data received, with a maxium size of bufferSize.
func OnReceive(callback func(*ReceiveInfo)) {
	tcp.Get("onReceive").Call("addListener", func(ri *ReceiveInfo) {
		ri.Data = js.Global.Get("Uint8Array").New(ri.dat).Interface().([]byte)
		callback(ri)
	})
}

type ReceiveError struct {
	*js.Object
	// integer	socketId
	// The socket identifier.
	SocketId int `js:"socketId"`
	// integer	resultCode
	// The result code returned from the underlying network call.
	ResultCode int `js:"resultCode"`
}

// onReceiveError
//
// Event raised when a network error occured while the runtime was waiting for data on the socket address and port. Once this event is raised, the socket is set to paused and no more onReceive events are raised for this socket.
//
// addListener
// chrome.sockets.tcp.onReceiveError.addListener(function callback)
// Parameters
// function	callback
// The callback parameter should be a function that looks like this:
//
// function(object info) {...};
// object	info
// The event data.
//
// integer	socketId
// The socket identifier.
//
// integer	resultCode
// The result code returned from the underlying network call.
func OnReceiveError(callback func(*ReceiveError)) {
	tcp.Get("onReceiveError").Call("addListener", callback)
}

// chrome.sockets.tcp.setPaused(integer socketId, boolean paused, function callback)
// Enables or disables the application from receiving messages from its peer. The default value is "false". Pausing a socket is typically used by an application to throttle data sent by its peer. When a socket is paused, no onReceive event is raised. When a socket is connected and un-paused, onReceive events are raised again when messages are received.
//
// Parameters
// integer	socketId
// boolean	paused
// function	(optional) callback
// Callback from the setPaused method.
//
// If you specify the callback parameter, it should be a function that looks like this:
//
// function() {...};
func SetPaused(socketId int, paused bool) {
	tcp.Call("setPaused", socketId, paused)
}

// setKeepAlive
//
// chrome.sockets.tcp.setKeepAlive(integer socketId, boolean enable, integer delay, function callback)
// Enables or disables the keep-alive functionality for a TCP connection.
//
// Parameters
// integer	socketId
// The socket identifier.
//
// boolean	enable
// If true, enable keep-alive functionality.
//
// integer	(optional) delay
// Set the delay seconds between the last data packet received and the first keepalive probe. Default is 0.
//
// function	callback
// Called when the setKeepAlive attempt is complete.
//
// The callback parameter should be a function that looks like this:
//
// function(integer result) {...};
// integer	result
// The result code returned from the underlying network call. A negative value indicates an error.
func SetKeepAlive(socketId int, enable bool, delay int, callback func(result int)) {
	tcp.Call("setKeepAlive", socketId, enable, delay, callback)
}

// setNoDelay
//
// chrome.sockets.tcp.setNoDelay(integer socketId, boolean noDelay, function callback)
// Sets or clears TCP_NODELAY for a TCP connection. Nagle's algorithm will be disabled when TCP_NODELAY is set.
//
// Parameters
// integer	socketId
// The socket identifier.
//
// boolean	noDelay
// If true, disables Nagle's algorithm.
//
// function	callback
// Called when the setNoDelay attempt is complete.
//
// The callback parameter should be a function that looks like this:
//
// function(integer result) {...};
// integer	result
// The result code returned from the underlying network call. A negative value indicates an error.
func SetNoDelay(socketId int, noDelay bool, callback func(result int)) {
	tcp.Call("setNoDelay", socketId, noDelay, callback)
}

// disconnect
//
// chrome.sockets.tcp.disconnect(integer socketId, function callback)
// Disconnects the socket.
//
// Parameters
// integer	socketId
// The socket identifier.
//
// function	(optional) callback
// Called when the disconnect attempt is complete.
//
// If you specify the callback parameter, it should be a function that looks like this:
//
// function() {...};
func Disconnect(socketId int) {
	tcp.Call("disconnect", socketId)
}

// close

// chrome.sockets.tcp.close(integer socketId, function callback)
// Closes the socket and releases the address/port the socket is bound to.
// Each socket created should be closed after use.
// The socket id is no no longer valid as soon at the function is called.
// However, the socket is guaranteed to be closed only when the callback is invoked.
//
// Parameters
// integer	socketId
// The socket identifier.
//
// function	(optional) callback
// Called when the close operation completes.
//
// If you specify the callback parameter, it should be a function that looks like this:
//
// function() {...};
func Close(socketId int) {
	tcp.Call("close", socketId)
}

// getInfo
//
// chrome.sockets.tcp.getInfo(integer socketId, function callback)
// Retrieves the state of the given socket.
//
// Parameters
// integer	socketId
// The socket identifier.
//
// function	callback
// Called when the socket state is available.
//
// The callback parameter should be a function that looks like this:
//
// function( SocketInfo socketInfo) {...};
// SocketInfo	socketInfo
// Object containing the socket information.
func GetInfo(socketId int, callback func(*sockets.SocketInfo)) {
	tcp.Call("getInfo", socketId, callback)
}

// getSockets
//
// chrome.sockets.tcp.getSockets(function callback)
// Retrieves the list of currently opened sockets owned by the application.
//
// Parameters
// function	callback
// Called when the list of sockets is available.
//
// The callback parameter should be a function that looks like this:
//
// function(array of SocketInfo socketInfos) {...};
// array of SocketInfo	socketInfos
// Array of object containing socket information.
func GetSockets(callback func([]*sockets.SocketInfo)) {
	tcp.Call("getSockets", callback)
}
