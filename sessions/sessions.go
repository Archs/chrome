package sessions

import (
	"github.com/Archs/chrome"
	"github.com/Archs/chrome/tabs"
	"github.com/Archs/chrome/windows"
	"github.com/gopherjs/gopherjs/js"
)

var (
	sessions = chrome.Get("sessions")
	// chrome.runtime.id	The ID of the extension/app.
	MAX_SESSION_RESULTS = sessions.Get("MAX_SESSION_RESULTS").Int()
)

/*
* Types
 */

type Filter struct {
	*js.Object
	MaxResults int `js:"maxResults"`
}

type Session struct {
	*js.Object
	LastModified int            `js:"lastModified"`
	Tab          tabs.Tab       `js:"tab"`
	Window       windows.Window `js:"window"`
}

type Device struct {
	*js.Object
	DeviceName string    `js:"deviceName"`
	Sessions   []Session `js:"sessions"`
}

/*
* Methods
 */

// GetRecentlyClosed gets the list of recently closed tabs and/or windows.
func GetRecentlyClosed(filter Filter, callback func(sessions []Session)) {
	sessions.Call("getRecentlyClosed", filter, callback)
}

// GetDevices retrieves all devices with synced sessions.
func GetDevices(filter Filter, callback func(devices []Device)) {
	sessions.Call("getDevices", filter, callback)
}

// Restore reopens a windows.Window or tabs.Tab, with an optional callback to run when the entry has been restored.
func Restore(sessionId string, callback func(restoredSession Session)) {
	sessions.Call("restore", sessionId, callback)
}

/*
* Events
 */

// OnChanged fired when recently closed tabs and/or windows are changed. This event does not monitor synced sessions changes.
func OnChanged(callback func()) {
	sessions.Get("onChanged").Call("addListener", callback)
}
