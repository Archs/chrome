package webRequest

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	webRequest                                        = chrome.Get("webRequest")
	MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES = webRequest.Get("MAX_HANDLER_BEHAVIOR_CHANGED_CALLS_PER_10_MINUTES").Int()
)

/*
* Types
 */

type RequestFilter struct {
	*js.Object
	Urls     []string `js:"urls"`
	Types    string   `js:"types"`
	TabId    int      `js:"tabId"`
	WindowId int      `js:"windowId"`
}

type BlockingResponse struct {
	*js.Object
	Cancel          bool              `js:"cancel"`
	RedirectUrl     string            `js:"redirectUrl"`
	RequestHeaders  js.M              `js:"requestHeaders"`
	ResponseHeaders js.M              `js:"responseHeaders"`
	AuthCredentials map[string]string `js:"authCredentials"`
}

type UploadData struct {
	*js.Object
	Bytes interface{} `js:"bytes"`
	File  string      `js:"file"`
}

/*
* Methods
 */

// HandlerBehaviorChanged needs to be called when the behavior of the webRequest handlers has
// changed to prevent incorrect handling due to caching. This function call is expensive. Don't call it often.
func HandlerBehaviorChanged(callback func()) {
	webRequest.Call("handlerBehaviorChanged", callback)
}

/*
* Events
 */

// OnBeforeRequest fired when a request is about to occur.
func OnBeforeRequest(callback func(details js.M)) {
	webRequest.Get("onBeforeRequest").Call("addListener", callback)
}

// OnBeforeSendHeaders fired before sending an HTTP request, once the request headers are available.
// This may occur after a TCP connection is made to the server, but before any HTTP data is sent.
func OnBeforeSendHeaders(callback func(details js.M)) {
	webRequest.Get("onBeforeSendHeaders").Call("addListener", callback)
}

// OnSendHeaders fired just before a request is going to be sent to the server (modifications of
// previous onBeforeSendHeaders callbacks are visible by the time onSendHeaders is fired).
func OnSendHeaders(callback func(details js.M)) {
	webRequest.Get("onSendHeaders").Call("addListener", callback)
}

// OnHeadersReceived fired when HTTP response headers of a request have been received.
func OnHeadersReceived(callback func(details js.M)) {
	webRequest.Get("onHeadersReceived").Call("addListener", callback)
}

// OnAuthRequired fired when an authentication failure is received. The listener has three options:
// it can provide authentication credentials, it can cancel the request and display the error page,
// or it can take no action on the challenge. If bad user credentials are provided, this may be
// called multiple times for the same request.
func OnAuthRequired(callback func(details js.M)) {
	webRequest.Get("onAuthRequired").Call("addListener", callback)
}

// OnResponseStarted fired when the first byte of the response body is received. For HTTP requests,
// this means that the status line and response headers are available.
func OnResponseStarted(callback func(details js.M)) {
	webRequest.Get("onResponseStarted").Call("addListener", callback)
}

// OnBeforeRedirect fired when a server-initiated redirect is about to occur.
func OnBeforeRedirect(callback func(details js.M)) {
	webRequest.Get("onBeforeRedirect").Call("addListener", callback)
}

// OnCompleted fired when a request is completed.
func OnCompleted(callback func(details js.M)) {
	webRequest.Get("onCompleted").Call("addListener", callback)
}

// OnErrorOccured fired when an error occurs.
func OnErrorOccured(callback func(details js.M)) {
	webRequest.Get("onErrorOccured").Call("addListener", callback)
}
