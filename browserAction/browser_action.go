package browserAction

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	browserAction = chrome.Get("browserAction")
)

/*
* Types
 */

type ColorArray []int

/*
* Methods:
 */

// SetTitle sets the title of the browser action. This shows up in the tooltip.
func SetTitle(details js.M) {
	browserAction.Call("setTitle", details)
}

// GetTitle gets the title of the browser action.
func GetTitle(details js.M, callback func(result string)) {
	browserAction.Call("getTitle", details, callback)
}

// SetIcon sets the icon for the browser action. The icon can be specified either as the path to an image
// file or as the pixel data from a canvas element, or as map of either one of those. Either the path
// or the imageData property must be specified.
func SetIcon(details js.M, callback func()) {
	browserAction.Call("setIcon", details, callback)
}

// SetPopup sets the html document to be opened as a popup when the user clicks on the browser action's icon.
func SetPopup(details js.M) {
	browserAction.Call("setPopup", details)
}

// GetPopup gets the html document set as the popup for this browser action.
func GetPopup(details js.M, callback func(result string)) {
	browserAction.Call("getPopup", details, callback)
}

// SetBadgeText sets the badge text for the browser action. The badge is displayed on top of the icon.
func SetBadgeText(details js.M) {
	browserAction.Call("setBadgeText", details)
}

// getBadgeText gets the badge text of the browser action. If no tab is specified, the non-tab-specific badge text is returned.
func getBadgeText(details js.M, callback func(result string)) {
	browserAction.Call("getBadgeText", details, callback)
}

// SetBadgeBackgroundColor sets the background color for the badge.
func SetBadgeBackgroundColor(details js.M) {
	browserAction.Call("setBadgeBackgroundColor", details)
}

// GetBadgeBackgroundColor gets the background color of the browser action.
func GetBadgeBackgroundColor(details js.M, callback func(result ColorArray)) {
	browserAction.Call("getBadgeBackgroundColor", details, callback)
}

// Enable enables the browser action for a tab. By default, browser actions are enabled.
func Enable(tabId int) {
	browserAction.Call("enable", tabId)
}

// Disable disables the browser action for a tab.
func Disable(tabId int) {
	browserAction.Call("disable", tabId)
}

/*
* Events
 */

// OnClicked fired when a browser action icon is clicked. This event will not fire if the browser action has a popup.
func OnClicked(callback func(tab tabs.Tab)) {
	browserAction.Get("onClicked").Call("addListener", callback)
}
