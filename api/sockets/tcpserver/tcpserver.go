package tcpserver

import (
	"github.com/Archs/chrome/api/sockets"
	"github.com/gopherjs/gopherjs/js"
)

var (
	tcp = js.Global.Get("chrome").Get("sockets").Get("tcpServer")
)

// SocketProperties
//
// properties
// boolean	(optional) persistent
// Flag indicating if the socket remains open when the event page of the application is unloaded (see Manage App Lifecycle). The default value is "false." When the application is loaded, any sockets previously opened with persistent=true can be fetched with getSockets.
//
// string	(optional) name
// An application-defined string associated with the socket.
//
// create
//
// chrome.sockets.tcpServer.create( SocketProperties properties, function callback)
// Creates a TCP server socket.
//
// Parameters
// SocketProperties	(optional) properties
// The socket properties (optional).
//
// function	 callback
// Called when the socket has been created.
//
// The callback parameter should be a function that looks like this:
//
// function(object createInfo) {...};
// js.M	 createInfo
// The result of the socket creation.
//
// integer	 socketId
// The ID of the newly created server socket. Note that socket IDs created from this API are not compatible with socket IDs created from other APIs, such as the deprecated socket API.
func Create(callback func(*sockets.CreateInfo)) {
	tcp.Call("create", callback)
}

// chrome.sockets.tcpServer.listen(integer socketId, string address, integer port, integer backlog, function callback)
// Listens for connections on the specified port and address. If the port/address is in use, the callback indicates a failure.
//
// Parameters
// integer	 socketId
// The socket identifier.
//
// string	 address
// The address of the local machine.
//
// integer	 port
// The port of the local machine. When set to 0, a free port is chosen dynamically. The dynamically allocated port can be found by calling getInfo.
//
// integer	(optional) backlog
// Length of the socket's listen queue. The default value depends on the Operating System (SOMAXCONN),
// which ensures a reasonable queue length for most applications.
//
// function	 callback
// Called when listen operation completes.
//
// The callback parameter should be a function that looks like this:
//
// function(integer result) {...};
// integer	 result
// The result code returned from the underlying network call. A negative value indicates an error.
func Listen(socketId int, address string, port int, callback func(result int)) {
	tcp.Call("listen", socketId, address, port, callback)
}

type AcceptInfo struct {
	*js.Object
	// integer	socketId
	// The socket identifier.
	SocketId int `js:"socketId"`
	// integer	 clientSocketId
	// The client socket identifier, i.e. the socket identifier of the newly established connection. This socket identifier should be used only with functions from the chrome.sockets.tcp namespace. Note the client socket is initially paused and must be explictly un-paused by the application to start receiving data.
	ClientSocketId int `js:"clientSocketId"`
}

// onAccept
//
// Event raised when a connection has been made to the server socket.
//
// addListener
// chrome.sockets.tcpServer.onAccept.addListener(function callback)
// Parameters
// function	 callback
// The callback parameter should be a function that looks like this:
//
// function(object info) {...};
// js.M	 info
// The event data.
//
// integer	 socketId
// The server socket identifier.
//
// integer	 clientSocketId
// The client socket identifier, i.e. the socket identifier of the newly established connection. This socket identifier should be used only with functions from the chrome.sockets.tcp namespace. Note the client socket is initially paused and must be explictly un-paused by the application to start receiving data.
func OnAccept(callback func(*AcceptInfo)) {
	tcp.Get("onAccept").Call("addListener", callback)
}

type AcceptError struct {
	*js.Object
	// integer	socketId
	// The socket identifier.
	SocketId int `js:"socketId"`
	// integer	resultCode
	// The result code returned from the underlying network call.
	ResultCode int `js:"resultCode"`
}

// onAcceptError
//
// Event raised when a network error occured while the runtime was waiting for new connections on the socket address and port. Once this event is raised, the socket is set to paused and no more onAccept events are raised for this socket until the socket is resumed.
//
// addListener
// chrome.sockets.tcpServer.onAcceptError.addListener(function callback)
// Parameters
// function	 callback
// The callback parameter should be a function that looks like this:
//
// function(object info) {...};
// js.M	 info
// The event data.
//
// integer	 socketId
// The server socket identifier.
//
// integer	 resultCode
// The result code returned from the underlying network call.
func OnAcceptError(callback func(*AcceptError)) {
	tcp.Get("onAcceptError").Call("addListener", callback)
}

// chrome.sockets.tcpServer.disconnect(integer socketId, function callback)
// Disconnects the listening socket, i.e. stops accepting new connections and releases the address/port the socket is bound to. The socket identifier remains valid, e.g. it can be used with listen to accept connections on a new port and address.
//
// Parameters
// integer	 socketId
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

// chrome.sockets.tcpServer.close(integer socketId, function callback)
// Disconnects and destroys the socket. Each socket created should be closed after use. The socket id is no longer valid as soon at the function is called. However, the socket is guaranteed to be closed only when the callback is invoked.
//
// Parameters
// integer	 socketId
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

// chrome.sockets.tcpServer.getInfo(integer socketId, function callback)
// Retrieves the state of the given socket.
//
// Parameters
// integer	 socketId
// The socket identifier.
//
// function	 callback
// Called when the socket state is available.
//
// The callback parameter should be a function that looks like this:
//
// function( SocketInfo socketInfo) {...};
// SocketInfo	 socketInfo
// js.M containing the socket information.
func GetInfo(socketId int, callback func(*sockets.SocketInfo)) {
	tcp.Call("getInfo", socketId, callback)
}

// chrome.sockets.tcpServer.getSockets(function callback)
// Retrieves the list of currently opened sockets owned by the application.
//
// Parameters
// function	 callback
// Called when the list of sockets is available.
//
// The callback parameter should be a function that looks like this:
//
// function(array of SocketInfo socketInfos) {...};
// array of SocketInfo	 socketInfos
// Array of js.M containing socket information.
func GetSockets(callback func([]*sockets.SocketInfo)) {
	tcp.Call("getSockets", callback)
}

// chrome.sockets.tcpServer.setPaused(integer socketId, boolean paused, function callback)
// Enables or disables a listening socket from accepting new connections. When paused, a listening socket accepts new connections until its backlog (see listen function) is full then refuses additional connection requests. onAccept events are raised only when the socket is un-paused.
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
