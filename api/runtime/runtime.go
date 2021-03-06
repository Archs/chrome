package runtime

import (
	"github.com/Archs/chrome"
	"github.com/Archs/chrome/api/tabs"
	"github.com/gopherjs/gopherjs/js"
)

var (
	runtime = chrome.Get("runtime")
	// chrome.runtime.id
	// The ID of the extension/app.
	Id = runtime.Get("id").String()
)

func GetLastError() string {
	msg := runtime.Get("lastError").Get("message")
	if msg == js.Undefined {
		return ""
	}
	return msg.String()
}

/*
* Types
 */

// An object which allows two way communication with other pages.
type Port struct {
	*js.Object
	Name         string        `js:"name"`
	onDisconnect *js.Object    `js:"onDisconnect"`
	onMessage    *js.Object    `js:"onMessage"`
	disconnect   *js.Object    `js:"disconnect"`
	postMessage  *js.Object    `js:"postMessage"`
	Sender       MessageSender `js:"sender"`
}

func (p *Port) PostMessage(message interface{}) {
	p.postMessage.Invoke(message)
}

func (p *Port) OnMessage(callback func(message interface{}, sender MessageSender, sendResponse func(interface{}))) {
	p.onMessage.Call("addListener", callback)
}

func (p *Port) OnDisconnect(callback func()) {
	p.onDisconnect.Call("addListener", callback)
}

func (p *Port) Disconnect() {
	p.disconnect.Invoke()
}

type MessageSender struct {
	*js.Object
	Tab          tabs.Tab `js:"tab"`
	FrameId      int      `js:"frameId"`
	Id           string   `js:"id"`
	Url          string   `js:"url"`
	TlsChannelId string   `js:"tlsChannelId"`
}

// An object containing information about the current platform.
type PlatformInfo struct {
	*js.Object
	// enum of "mac", "win", "android", "cros", "linux", or "openbsd"	os
	// The operating system chrome is running on.
	Os string `js:"os"`

	// enum of "arm", "x86-32", or "x86-64"	arch
	// The machine's processor architecture.
	Arch string `js:"arch"`

	// enum of "arm", "x86-32", or "x86-64"	nacl_arch
	// The native client architecture. This may be different from arch on some platforms.
	NaclArch string `js:"nacl_arch"`
}

/*
* Methods
 */

func GetBackgroundPage(callback func(backgroundPage interface{})) {
	runtime.Call("getBackgroundPage", callback)
}

func GetManifest() *js.Object {
	return runtime.Call("getManifest")
}

func GetURL(path string) string {
	return runtime.Call("getURL", path).String()
}

func Reload() {
	runtime.Call("reload")
}

// Maybe make status an Enum string with specific values.
func RequestUpdateCheck(callback func(status string, details interface{})) {
	runtime.Call("requestUpdateCheck", callback)
}

func Restart() {
	runtime.Call("restart")
}

func Connect(extensionId string, connectInfo interface{}) *Port {
	portObj := runtime.Call("connect", extensionId, connectInfo)
	return &Port{Object: portObj}
}

func ConnectNative(application string) *Port {
	portObj := runtime.Call("connectNative", application)
	return &Port{Object: portObj}
}

func SendMessage(extensionId string, message interface{}, options interface{}, responseCallback func(response interface{})) {
	runtime.Call("sendMessage", extensionId, message, options, responseCallback)
}

func SendNativeMessage(application string, message interface{}, responseCallback func(response interface{})) {
	runtime.Call("sendNativeMessage", application, message, responseCallback)
}

func GetPlatformInfo(callback func(platformInfo PlatformInfo)) {
	runtime.Call("getPlatformInfo", callback)
}

func GetPackageDirectoryEntry(callback func(directoryEntry interface{})) {
	runtime.Call("getPackageDirectoryEntry", callback)
}

/*
* Events
 */

func OnStartup(callback func()) {
	runtime.Get("onStartup").Call("addListener", callback)
}

func OnInstalled(callback func(details map[string]string)) {
	runtime.Get("onInstalled").Call("addListener", callback)
}

func OnSuspend(callback func()) {
	runtime.Get("onSuspend").Call("addListener", callback)
}

func OnSuspendCanceled(callback func()) {
	runtime.Get("onSuspendCanceled").Call("addListener", callback)
}

func OnUpdateAvailable(callback func(details map[string]string)) {
	runtime.Get("onUpdateAvailable").Call("addListener", callback)
}

func OnConnect(callback func(port Port)) {
	runtime.Get("onConnect").Call("addListener", callback)
}

func OnConnectExternal(callback func(port Port)) {
	runtime.Get("onConnectExternal").Call("addListener", callback)
}

func OnMessage(callback func(message interface{}, sender MessageSender, sendResponse func(interface{}))) {
	runtime.Get("onMessage").Call("addListener", callback)
}

func OnMessageExternal(callback func(message interface{}, sender MessageSender, sendResponse func(interface{}))) {
	runtime.Get("onMessageExternal").Call("addListener", callback)
}

func OnRestartRequired(callback func(reason string)) {
	runtime.Get("onRestartRequired").Call("addListener", callback)
}
