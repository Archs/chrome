package contextMenus

import (
	"github.com/Archs/chrome"
	"github.com/Archs/chrome/tabs"
	"github.com/gopherjs/gopherjs/js"
)

var (
	contextMenus = chrome.Get("contextMenus")
	// The maximum number of top level extension items that can be added to an extension action context menu. Any items beyond this limit will be ignored.
	ACTION_MENU_TOP_LEVEL_LIMIT = contextMenus.Get("ACTION_MENU_TOP_LEVEL_LIMIT").Int()
)

/*
* Methods:
 */

// Create creates a new context menu item. Note that if an error occurs during creation,
// you may not find out until the creation callback fires (the details will be in chrome.Runtime.LastError).
func Create(createProperties js.M, callback func()) {
	contextMenus.Call("create", createProperties, callback)
}

// Update updates a previously created context menu item.
func Update(id interface{}, updateProperties js.M, callback func()) {
	contextMenus.Call("update", id, updateProperties, callback)
}

// Remove removes a context menu item.
func Remove(menuItemId interface{}, callback func()) {
	contextMenus.Call("remove", menuItemId, callback)
}

// RemoveAll removes all context menu items added by this extension.
func RemoveAll(callback func()) {
	contextMenus.Call("removeAll", callback)
}

/*
* Events
 */

// OnClicked fired when a context menu item is clicked.
func OnClicked(callback func(info js.M, tab tabs.Tab)) {
	contextMenus.Get("onClicked").Call("addListener", callback)
}
