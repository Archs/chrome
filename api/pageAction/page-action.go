package pageAction

import (
	"github.com/Archs/chrome"
	"github.com/Archs/chrome/api/tabs"
	"github.com/gopherjs/gopherjs/js"
)

var (
	pageAction = chrome.Get("pageAction")
)

/*
* Methods
 */

// Show shows the page action. The page action is shown whenever the tab is selected.
func Show(tabId int) {
	pageAction.Call("show", tabId)
}

// Hide hides the page action.
func Hide(tabId int) {
	pageAction.Call("hide", tabId)
}

// SetTitle sets the title of the page action. This is displayed in a tooltip over the page action.
func SetTitle(details js.M) {
	pageAction.Call("setTitle", details)
}

// GetTitle gets the title of the page action.
func GetTitle(details js.M, callback func(result string)) {
	pageAction.Call("getTitle", details, callback)
}

// SetIcon sets the icon for the page action. The icon can be specified either as the path to an image
// file or as the pixel data from a canvas element, or as dictionary of either one of those. Either the
// path or the imageData property must be specified.
func SetIcon(details js.M, callback func()) {
	pageAction.Call("setIcon", details, callback)
}

// SetPopup sets the html document to be opened as a popup when the user clicks on the page action's icon.
func SetPopup(details js.M) {
	pageAction.Call("setPopup", details)
}

// GetPopup gets the html document set as the popup for this page action.
func GetPopup(details js.M, callback func(result string)) {
	pageAction.Call("getPopup", details, callback)
}

/*
* Events
 */

// OnClicked fired when a page action icon is clicked. This event will not fire if the page action has a popup.
func OnClicked(callback func(tab tabs.Tab)) {
	pageAction.Get("onClicked").Call("addListener", callback)
}
