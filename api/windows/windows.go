package windows

import (
	"github.com/Archs/chrome"
	"github.com/Archs/chrome/api/tabs"
	"github.com/gopherjs/gopherjs/js"
)

var (
	windows = chrome.Get("windows")
	// -1	chrome.windows.WINDOW_ID_NONE
	// Since Chrome 6.
	WINDOW_ID_NONE = windows.Get("WINDOW_ID_NONE").Int()
	// The windowId value that represents the absence of a chrome browser window.
	// -2	chrome.windows.WINDOW_ID_CURRENT
	// Since Chrome 18.
	WINDOW_ID_CURRENT = windows.Get("WINDOW_ID_CURRENT").Int()
)

/*
* Types
 */

type Window struct {
	js.Object
	Id          int        `js:"id"`
	Focused     bool       `js:"focused"`
	Top         int        `js:"top"`
	Left        int        `js:"left"`
	Width       int        `js:"width"`
	Height      int        `js:"height"`
	Tabs        []tabs.Tab `js:"tabs"`
	Incognito   bool       `js:"incognito"`
	Type        string     `js:"type"`
	State       string     `js:"state"`
	AlwaysOnTop bool       `js:"alwaysOnTop"`
	SessionId   string     `js:"sessionId"`
}

/*
* Methods
 */

// Get gets details about a window.
func Get(windowId int, getInfo js.M, callback func(window Window)) {
	windows.Call("get", windowId, getInfo, callback)
}

// GetCurrent gets the current window.
func GetCurrent(getInfo js.M, callback func(window Window)) {
	windows.Call("getCurrent", getInfo, callback)
}

// GetLastFocused gets the window that was most recently focused — typically the window 'on top'.
func GetLastFocused(getInfo js.M, callback func(window Window)) {
	windows.Call("getLastFocused", getInfo, callback)
}

// GetAll gets all windows.
func GetAll(getInfo js.M, callback func(windows []Window)) {
	windows.Call("getAll", getInfo, callback)
}

// Create creates (opens) a new browser with any optional sizing, position or default URL provided.
func Create(createData js.M, callback func(window Window)) {
	windows.Call("create", createData, callback)
}

// Update updates the properties of a window. Specify only the properties that you
// want to change; unspecified properties will be left unchanged.
func Update(windowId int, updateInfo js.M, callback func(window Window)) {
	windows.Call("update", windowId, updateInfo, callback)
}

// Remove removes (closes) a window, and all the tabs inside it.
func Remove(windowId int, callback func(js.Object)) {
	windows.Call("remove", windowId, callback)
}

/*
* Events
 */

// OnCreated fired when a window is created.
func OnCreated(callback func(window Window)) {
	windows.Get("onCreated").Call("addListener", callback)
}

// OnRemoved fired when a window is removed (closed).
func OnRemoved(callback func(windowId int)) {
	windows.Get("onRemoved").Call("addListener", callback)
}

// onFocusChanged fired when the currently focused window changes.
// Will be chrome.windows.WINDOW_ID_NONE if all chrome windows have lost focus.
// Note: On some Linux window managers, WINDOW_ID_NONE will always be sent immediately
// preceding a switch from one chrome window to another.
func onFocusChanged(callback func(windowId int)) {
	windows.Get("onFocusChanged").Call("addListener", callback)
}
