package gcm

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	gcm = chrome.Get("gcm")
	// 4,096	chrome.gcm.MAX_MESSAGE_SIZE	The maximum size (in bytes) of all key/value pairs in a message.
	MAX_MESSAGE_SIZE = gcm.Get("MAX_MESSAGE_SIZE").Int()
)

/*
* Methods:
 */

// Register registers the application with GCM. The registration ID will be returned by the callback.
// If register is called again with the same list of senderIds, the same registration ID will be returned.
func Register(senderIds []string, callback func(registrationId string)) {
	gcm.Call("register", senderIds, callback)
}

// Unregister unregisters the application from GCM.
func Unregister(callback func()) {
	gcm.Call("unregister", callback)
}

// Send sends a message according to its contents.
func Send(message js.M, callback func(messageId string)) {
	gcm.Call("send", message, callback)
}

/*
* Events
 */

// OnMessage fired when a message is received through GCM.
func OnMessage(callback func(message js.M)) {
	gcm.Get("onMessage").Call("addListener", callback)
}

// OnMessageDeleted fired when a GCM server had to delete messages sent by an app server to the application.
// See Messages deleted event section of Cloud Messaging documentation for details on handling this event.
func OnMessageDeleted(callback func()) {
	gcm.Get("onMessageDeleted").Call("addListener", callback)
}

// OnSendError fired when it was not possible to send a message to the GCM server.
func OnSendError(callback func(error js.M)) {
	gcm.Get("onSendError").Call("addListener", callback)
}
