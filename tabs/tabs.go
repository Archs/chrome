package tabs

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	tabs = chrome.Get("tabs")
)

/*
* Types
 */
type Tab struct {
	*js.Object
	Id          int    `js:"id"`
	Index       int    `js:"index"`
	WindowId    int    `js:"windowId"`
	OpenerTabId int    `js:"openerTabId"`
	Selected    bool   `js:"selected"`
	Highlighted bool   `js:"highlighted"`
	Active      bool   `js:"active"`
	Pinned      bool   `js:"pinned"`
	Url         string `js:"url"`
	Title       string `js:"title"`
	FavIconUrl  string `js:"favIconUrl"`
	Status      string `js:"status"`
	Incognito   bool   `js:"incognito"`
	Width       int    `js:"width"`
	Height      int    `js:"height"`
	SessionId   string `js:"sessionId"`
}

type ZoomSettings map[string]string

/*
* Methods:
 */

// Get retrieves details about the specified tab.
func Get(tabId int, callback func(tab Tab)) {
	tabs.Call("get", tabId, callback)
}

// GetCurrent gets the tab that this script call is being made from.
// May be undefined if called from a non-tab context (for example: a background page or popup view).
func GetCurrent(callback func(tab Tab)) {
	tabs.Call("getCurrent", callback)
}

// Connect Connects to the content script(s) in the specified tab.
// The runtime.onConnect event is fired in each content script running in the specified tab
// for the current extension. For more details, see Content Script Messaging (https://developer.chrome.com/extensions/messaging)
func Connect(tabId int, connectInfo interface{}) {
	tabs.Call("connect", tabId, connectInfo)
}

// SendMessage sends a single message to the content script(s) in the specified tab,
//  with an optional callback to run when a response is sent back. The runtime.onMessage
// event is fired in each content script running in the specified tab for the current extension.
func SendMessage(tabId int, message interface{}, responseCallback func(response js.M)) {
	tabs.Call("sendMessage", tabId, message, responseCallback)
}

// GetSelected gets the tab that is selected in the specified window.
func GetSelected(windowId int, callback func(tab Tab)) {
	tabs.Call("getSelected", windowId, callback)
}

// GetAllInWindow gets details about all tabs in the specified window.
func GetAllInWindow(windowId int, callback func(tabs []Tab)) {
	tabs.Call("getAllInWindow", windowId, callback)
}

// Create creates a new tab.
func Create(createProperties interface{}, callback func(tab Tab)) {
	tabs.Call("create", createProperties, callback)
}

// Duplicate duplicates a tab.
func Duplicate(tabId int, callback func(tab Tab)) {
	tabs.Call("duplicate", tabId, callback)
}

// Query gets all tabs that have the specified properties, or all tabs if no properties are specified.
func Query(queryInfo interface{}, callback func(result []Tab)) {
	tabs.Call("query", queryInfo, callback)
}

// Highlight highlights the given tabs
func Highlight(highlightInfo interface{}, callback func(js.Object)) {
	tabs.Call("highlight", highlightInfo, callback)
}

// Update modifies the properties of a tab. Properties that are not specified in updateProperties are not modified.
func Update(tabId int, updateProperties interface{}, callback func(tab Tab)) {
	tabs.Call("highlight", updateProperties, callback)
}

// Move moves one or more tabs to a new position within its window, or to a new window.
// Note that tabs can only be moved to and from normal windows.
func Move(tabIds []interface{}, moveProperties interface{}, callback func(tabs []Tab)) {
	tabs.Call("move", tabIds, moveProperties, callback)
}

// Reload reloads a tab.
func Reload(tabId int, reloadProperties interface{}, callback func(js.Object)) {
	tabs.Call("reload", tabId, reloadProperties, callback)
}

// Remove closes one or more tabs.
func Remove(tabIds []interface{}, callback func(js.Object)) {
	tabs.Call("remove", tabIds, callback)
}

// DetectLanguage detects the primary language of the content in a tab.
func DetectLanguage(tabId int, callback func(language string)) {
	tabs.Call("detectLanguage", tabId, callback)
}

// CaptureVisibleTab captures the visible area of the currently active tab in the specified window.
// You must have <all_urls> permission to use this method.
func CaptureVisibleTab(windowId int, options interface{}, callback func(dataUrl string)) {
	tabs.Call("captureVisibleTab", windowId, options, callback)
}

// ExecuteScript injects JavaScript code into a page. For details, see the programmatic
// injection section of the content scripts doc: (https://developer.chrome.com/extensions/content_scripts#pi)
func ExecuteScript(tabId int, details interface{}, callback func(result []interface{})) {
	tabs.Call("executeScript", tabId, details, callback)
}

// InsertCss injects CSS into a page. For details, see the programmatic injection
// section of the content scripts doc. (https://developer.chrome.com/extensions/content_scripts#pi)
func InsertCss(tabId int, details interface{}, callback func()) {
	tabs.Call("insertCss", tabId, details, callback)
}

// SetZoom zooms a specified tab.
func SetZoom(tabId int, zoomFactor int64, callback func()) {
	tabs.Call("setZoom", tabId, zoomFactor, callback)
}

// GetZoom gets the current zoom factor of a specified tab.
func GetZoom(tabId int, callback func(zoomFactor int64)) {
	tabs.Call("getZoom", tabId, callback)
}

// SetZoomSettings sets the zoom settings for a specified tab, which define how zoom changes are handled.
// These settings are reset to defaults upon navigating the tab.
func SetZoomSettings(tabId int, zoomSettings ZoomSettings, callback func()) {
	tabs.Call("setZoomSettings", tabId, zoomSettings, callback)
}

// GetZoomSettings gets the current zoom settings of a specified tab.
func GetZoomSettings(tabId int, callback func(zoomSettings ZoomSettings)) {
	tabs.Call("getZoomSettings", tabId, callback)
}

/*
* Events
 */

// OnCreated is fired when a tab is created. Note that the tab's URL may not be set at the time
// this event fired, but you can listen to onUpdated events to be notified when a URL is set.
func OnCreated(callback func(tab Tab)) {
	tabs.Get("onCreated").Call("addListener", callback)
}

// OnUpdated fired when a tab is updated.
func OnUpdated(callback func(tabId int, changeInfo js.M, tab Tab)) {
	tabs.Get("onUpdated").Call("addListener", callback)
}

// OnMoved fired when a tab is moved within a window. Only one move event is fired,
// representing the tab the user directly moved. Move events are not fired for the
// other tabs that must move in response. This event is not fired when a tab is moved between windows.
func OnMoved(callback func(tabId int, movedInfo js.M)) {
	tabs.Get("onMoved").Call("addListener", callback)
}

// OnActivated fires when the active tab in a window changes. Note that the tab's URL
// may not be set at the time this event fired, but you can listen to onUpdated events
// to be notified when a URL is set.
func OnActivated(callback func(activeInfo js.M)) {
	tabs.Get("onActivated").Call("addListener", callback)
}

// OnHighlightChanged fired when the highlighted or selected tabs in a window changes.
func OnHighlightChanged(callback func(selectInfo js.M)) {
	tabs.Get("onHighlightChanged").Call("addListener", callback)
}

// OnHighlighted fired when the highlighted or selected tabs in a window changes.
func OnHighlighted(callback func(highlightInfo js.M)) {
	tabs.Get("onHighlighted").Call("addListener", callback)
}

// OnDetached fired when a tab is detached from a window, for example because it is being moved between windows.
func OnDetached(callback func(tabId int, detachInfo js.M)) {
	tabs.Get("onDetached").Call("addListener", callback)
}

// OnAttached fired when a tab is attached to a window, for example because it was moved between windows.
func OnAttached(callback func(tabId int, attachInfo js.M)) {
	tabs.Get("onAttached").Call("addListener", callback)
}

// OnRemoved fired when a tab is closed.
func OnRemoved(callback func(tabId int, removeInfo js.M)) {
	tabs.Get("onRemoved").Call("addListener", callback)
}

// OnReplaced fired when a tab is replaced with another tab due to prerendering or instant.
func OnReplaced(callback func(addedTabId int, removedTabId int)) {
	tabs.Get("OnReplaced").Call("addListener", callback)
}

// OnZoomChange fired when a tab is zoomed.
func OnZoomChange(callback func(zoomChangeInfo js.M)) {
	tabs.Get("onZoomChange").Call("addListener", callback)
}
