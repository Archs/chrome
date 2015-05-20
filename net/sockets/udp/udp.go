package udp

import (
	"github.com/Archs/chrome/net/sockets"
	"github.com/gopherjs/gopherjs/js"
)

var (
	udp = js.Global.Get("chrome").Get("sockets").Get("udp")
)

// chrome.sockets.udp.create( SocketProperties properties, function callback)
// Creates a udp socket.
//
// Parameters
// SocketProperties (optional) properties
// The socket properties (optional).
//
// function callback
// Called when the socket has been created.
//
// The callback parameter should be a function that looks like this:
//
// function(object createInfo) {...};
//    object    createInfo
//       The result of the socket creation.
//
//    integer   socketId
//        The ID of the newly created socket. Note that socket IDs created from this API are not compatible with socket IDs created from other APIs, such as the deprecated socket API.
func Create(p sockets.SocketProperties, callback func(*sockets.CreateInfo)) {
	udp.Call("create", p, callback)
}

func CreateM(socketProperties js.M, callback func(*sockets.CreateInfo)) {
	udp.Call("create", socketProperties, callback)
}

func CreateEx(callback func(*sockets.CreateInfo)) {
	udp.Call("create", callback)
}

// chrome.sockets.udp.update(integer socketId, SocketProperties properties, function callback)
// Updates the socket properties.
//
// Parameters
// integer  socketId
// The socket identifier.
//
// SocketProperties properties
// The properties to update.
//
// function (optional) callback
// Called when the properties are updated.
//
// If you specify the callback parameter, it should be a function that looks like this:
//
// function() {...};
func Update(socketId int, properties *sockets.SocketProperties, callback func()) {
	udp.Call("update", socketId, properties, callback)
}

func UpdateEx(socketId int, properties *sockets.SocketProperties) {
	udp.Call("update", socketId, properties)
}

// chrome.sockets.udp.bind(integer socketId, string address, integer port, function callback)
// Binds the local address and port for the socket. For a client socket, it is recommended to use port 0 to let the platform pick a free port.
//
// Once the bind operation completes successfully, onReceive events are raised when UDP packets arrive on the address/port specified -- unless the socket is paused.
//
// Parameters
// integer socketId
// The socket ID.
//
// string  address
// The address of the local machine. DNS name, IPv4 and IPv6 formats are supported. Use "0.0.0.0" to accept packets from all local available network interfaces.
//
// integer port
// The port of the local machine. Use "0" to bind to a free port.
//
// function    callback
// Called when the bind operation completes.
//
// The callback parameter should be a function that looks like this:
// // function(integer result) {...};
// integer result
// The result code returned from the underlying network call.
// A negative value indicates an error.
func Bind(socketId int, address string, port int, callback func(result int)) {
	udp.Call("bind", socketId, address, port, callback)
}

type SendInfo struct {
	*js.Object
	// integer  resultCode
	// The result code returned from the underlying network call. A negative value indicates an error.
	ResultCode int `js:"resultCode"`
	// integer  (optional) bytesSent
	// The number of bytes sent (if result == 0)
	BytesSent int `js:"bytesSent"`
}

// chrome.sockets.udp.send(integer socketId, ArrayBuffer data, string address, integer port, function callback)
// Sends data on the given socket to the given address and port. The socket must be bound to a local port before calling this method.
//
// Parameters
// integer socketId
// The socket ID.
//
// ArrayBuffer data
// The data to send.
//
// string  address
// The address of the remote machine.
//
// integer port
// The port of the remote machine.
//
// function    callback
// Called when the send operation completes.
//
// The callback parameter should be a function that looks like this:
//
// function(object sendInfo) {...};
// object  sendInfo
// Result of the send method.
//
// integer resultCode
// The result code returned from the underlying network call. A negative value indicates an error.
//
// integer (optional) bytesSent
// The number of bytes sent (if result == 0)

func Send(socketId int, data []byte, address string, port int, callback func(*SendInfo)) {
	udp.Call("send", socketId, js.NewArrayBuffer(data), address, port, callback)
}

type ReceiveInfo struct {
	*js.Object
	// integer  socketId
	// The socket identifier.
	SocketId int `js:"socketId"`
	// ArrayBuffer  data
	// The data received, with a maxium size of bufferSize.
	dat           *js.Object `js:"data"`
	RemoteAddress string     `js:"remoteAddress"`
	RemotePort    string     `js:"remotePort"`
	Data          []byte
}

// onReceive
//
// Event raised when a UDP packet has been received for the given socket.
//
// addListener
// chrome.sockets.udp.onReceive.addListener(function callback)
// Parameters
// function    callback
// The callback parameter should be a function that looks like this:
//
// function(object info) {...};
// object  info
// The event data.
//
// integer socketId
// The socket ID.
//
// ArrayBuffer data
// The UDP packet content (truncated to the current buffer size).
//
// string  remoteAddress
// The address of the host the packet comes from.
//
// integer remotePort
// The port of the host the packet comes from.
func OnReceive(callback func(*ReceiveInfo)) {
	udp.Get("onReceive").Call("addListener", func(ri *ReceiveInfo) {
		ri.Data = js.Global.Get("Uint8Array").New(ri.dat).Interface().([]byte)
		callback(ri)
	})
}

type ReceiveError struct {
	*js.Object
	// integer  socketId
	// The socket identifier.
	SocketId int `js:"socketId"`
	// integer  resultCode
	// The result code returned from the underlying network call.
	ResultCode int `js:"resultCode"`
}

// onReceiveError
//
// Event raised when a network error occured while the runtime was waiting for data on the socket address and port. Once this event is raised, the socket is paused and no more onReceive events will be raised for this socket until the socket is resumed.
//
// addListener
// chrome.sockets.udp.onReceiveError.addListener(function callback)
// Parameters
// function callback
// The callback parameter should be a function that looks like this:
//
// function(object info) {...};
// object   info
// The event data.
//
// integer  socketId
// The socket identifier.
//
// integer  resultCode
// The result code returned from the underlying network call.
func OnReceiveError(callback func(*ReceiveError)) {
	udp.Get("onReceiveError").Call("addListener", callback)
}

// chrome.sockets.udp.setPaused(integer socketId, boolean paused, function callback)
// Enables or disables the application from receiving messages from its peer. The default value is "false". Pausing a socket is typically used by an application to throttle data sent by its peer. When a socket is paused, no onReceive event is raised. When a socket is connected and un-paused, onReceive events are raised again when messages are received.
//
// Parameters
// integer  socketId
// boolean  paused
// function (optional) callback
// Callback from the setPaused method.
//
// If you specify the callback parameter, it should be a function that looks like this:
//
// function() {...};
func SetPaused(socketId int, paused bool) {
	udp.Call("setPaused", socketId, paused)
}

// joinGroup
//
// chrome.sockets.udp.joinGroup(integer socketId, string address, function callback)
// Joins the multicast group and starts to receive packets from that group. The socket must be bound to a local port before calling this method.
//
// Parameters
// integer socketId
// The socket ID.
//
// string  address
// The group address to join. Domain names are not supported.
//
// function    callback
// Called when the joinGroup operation completes.
//
// The callback parameter should be a function that looks like this:
//
// function(integer result) {...};
// integer result
// The result code returned from the underlying network call.
// A negative value indicates an error.
func JoinGroup(socketId int, address string, callback func(result int)) {
	udp.Call("joinGroup", address, callback)
}

// leaveGroup
//
// chrome.sockets.udp.leaveGroup(integer socketId, string address, function callback)
// Leaves the multicast group previously joined using joinGroup. This is only necessary to call if you plan to keep using the socketafterwards, since it will be done automatically by the OS when the socket is closed.
//
// Leaving the group will prevent the router from sending multicast datagrams to the local host, presuming no other process on the host is still joined to the group.
//
// Parameters
// integer socketId
// The socket ID.
//
// string  address
// The group address to leave. Domain names are not supported.
//
// function    callback
// Called when the leaveGroup operation completes.
//
// The callback parameter should be a function that looks like this:
//
// function(integer result) {...};
// integer result
// The result code returned from the underlying network call. A negative value indicates an error.
func LeaveGroup(socketId int, address string, callback func(result int)) {
	udp.Call("leaveGroup", address, callback)
}

// setMulticastTimeToLive
//
// chrome.sockets.udp.setMulticastTimeToLive(integer socketId, integer ttl, function callback)
// Sets the time-to-live of multicast packets sent to the multicast group.
//
// Calling this method does not require multicast permissions.
//
// Parameters
// integer socketId
// The socket ID.
//
// integer ttl
// The time-to-live value.
//
// function    callback
// Called when the configuration operation completes.
//
// The callback parameter should be a function that looks like this:
//
// function(integer result) {...};
// integer result
// The result code returned from the underlying network call. A negative value indicates an error.
func SetMulticastTimeToLive(socketId int, ttl int, callback func(result int)) {
	udp.Call("setMulticastTimeToLive", socketId, ttl, callback)
}

// setMulticastLoopbackMode
//
// chrome.sockets.udp.setMulticastLoopbackMode(integer socketId, boolean enabled, function callback)
// Sets whether multicast packets sent from the host to the multicast group will be looped back to the host.
//
// Note: the behavior of setMulticastLoopbackMode is slightly different between Windows and Unix-like systems. The inconsistency happens only when there is more than one application on the same host joined to the same multicast group while having different settings on multicast loopback mode. On Windows, the applications with loopback off will not RECEIVE the loopback packets; while on Unix-like systems, the applications with loopback off will not SEND the loopback packets to other applications on the same host. See MSDN: http://goo.gl/6vqbj
//
// Calling this method does not require multicast permissions.
//
// Parameters
// integer socketId
// The socket ID.
//
// boolean enabled
// Indicate whether to enable loopback mode.
//
// function    callback
// Called when the configuration operation completes.
//
// The callback parameter should be a function that looks like this:
//
// function(integer result) {...};
// integer result
// The result code returned from the underlying network call. A negative value indicates an error.
func SetMulticastLoopbackMode(socketId int, enabled bool, callback func(result int)) {
	udp.Call("setMulticastLoopbackMode", socketId, enabled, callback)
}

// close

// chrome.sockets.udp.close(integer socketId, function callback)
// Closes the socket and releases the address/port the socket is bound to.
// Each socket created should be closed after use.
// The socket id is no no longer valid as soon at the function is called.
// However, the socket is guaranteed to be closed only when the callback is invoked.
//
// Parameters
// integer  socketId
// The socket identifier.
//
// function (optional) callback
// Called when the close operation completes.
//
// If you specify the callback parameter, it should be a function that looks like this:
//
// function() {...};
func Close(socketId int) {
	udp.Call("close", socketId)
}

// getJoinedGroups
//
// chrome.sockets.udp.getJoinedGroups(integer socketId, function callback)
// Gets the multicast group addresses the socket is currently joined to.
//
// Parameters
// integer socketId
// The socket ID.
//
// function    callback
// Called with an array of strings of the result.
//
// The callback parameter should be a function that looks like this:
//
// function(array of string groups) {...};
// array of string groups
// Array of groups the socket joined.
func GetJoinedGroups(socketId int, callback func([]string)) {
	udp.Call("getJoinedGroups", socketId, callback)
}

// getInfo
//
// chrome.sockets.udp.getInfo(integer socketId, function callback)
// Retrieves the state of the given socket.
//
// Parameters
// integer  socketId
// The socket identifier.
//
// function callback
// Called when the socket state is available.
//
// The callback parameter should be a function that looks like this:
//
// function( SocketInfo socketInfo) {...};
// SocketInfo   socketInfo
// Object containing the socket information.
func GetInfo(socketId int, callback func(*sockets.SocketInfo)) {
	udp.Call("getInfo", socketId, callback)
}

// getSockets
//
// chrome.sockets.udp.getSockets(function callback)
// Retrieves the list of currently opened sockets owned by the application.
//
// Parameters
// function callback
// Called when the list of sockets is available.
//
// The callback parameter should be a function that looks like this:
//
// function(array of SocketInfo socketInfos) {...};
// array of SocketInfo  socketInfos
// Array of object containing socket information.
func GetSockets(callback func([]*sockets.SocketInfo)) {
	udp.Call("getSockets", callback)
}
