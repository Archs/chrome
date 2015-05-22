package debugger

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	debugger = chrome.Get("debugger")
)

/*
* Types
 */

type Debugee struct {
	*js.Object
	TabId       int    `js:"tabId"`
	ExtensionId string `js:"extensionId"`
	TargetId    string `js:"targetId"`
}

type TargetInfo struct {
	*js.Object
	Type        string `js:"type"`
	Id          string `js:"id"`
	TabId       int    `js:"tabId"`
	ExtensionId string `js:"extensionId"`
	Attached    bool   `js:"attached"`
	Title       string `js:"title"`
	Url         string `js:"url"`
	FaviconUrl  string `js:"faviconUrl"`
}

/*
* Methods:
 */

// Attach attaches debugger to the given target.
func Attach(target Debugee, requiredVersion string, callback func()) {
	debugger.Call("attach", target, requiredVersion, callback)
}

// Detach detaches debugger from the given target.
func Detach(target Debugee, callback func()) {
	debugger.Call("detach", target, callback)
}

// SendCommand sends given command to the debugging target.
func SendCommand(target Debugee, method string, commandParams js.M, callback func(result js.M)) {
	debugger.Call("sendCommand", target, method, commandParams, callback)
}

// GetTargets returns the list of available debug targets.
func GetTargets(callback func(result []TargetInfo)) {
	debugger.Call("getTargets", callback)
}

/*
* Events
 */

// OnEvent fired whenever debugging target issues instrumentation event.
func OnEvent(callback func(source Debugee, method string, params js.M)) {
	debugger.Get("onEvent").Call("addListener", callback)
}

// OnDetach fired when browser terminates debugging session for the tab. This happens when
// either the tab is being closed or Chrome DevTools is being invoked for the attached tab.
func OnDetach(callback func(source Debugee, reason string)) {
	debugger.Get("onDetach").Call("addListener", callback)
}
