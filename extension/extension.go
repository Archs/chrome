package extension

import (
	"github.com/Archs/chrome"
	"github.com/Archs/chrome/windows"
	"github.com/gopherjs/gopherjs/js"
)

var (
	extension = chrome.Get("extension")
)

// Set for the lifetime of a callback if an ansychronous extension api has resulted in an error. If no error has occured lastError will be undefined.
func LastError() string {
	return extension.Get("lastError").String()
}

// True for content scripts running inside incognito tabs, and for extension pages running inside an incognito process. The latter only applies to extensions with 'split' incognito_behavior.
func InIncognitoContext() bool {
	return extension.Get("inIncognitoContext").Bool()
}

/*
* Methods:
 */

// GetURL converts a relative path within an extension install directory to a fully-qualified URL.
func GetURL(path string) {
	extension.Call("getURL", path)
}

// GetViews returns an array of the JavaScript 'window' objects for each of the pages running inside the current extension.
// Fix this and the other functions to return windows.Window objects instead of js.Object or whatever else
func GetViews(fetchProperties js.M) []windows.Window {
	wins := []windows.Window{}
	windowObjs := extension.Call("getViews", fetchProperties)
	for i := 0; i < windowObjs.Length(); i++ {
		window := windowObjs.Index(i)
		wins = append(wins, windows.Window{Object: *window})
	}
	return wins
}

// GetBackgroundPage returns the JavaScript 'window' object for the background page running inside
// the current extension. Returns null if the extension has no background page.
func GetBackgroundPage() windows.Window {
	return windows.Window{Object: *extension.Call("getBackgroundPage")}
}

// GetExtensionTabs returns an array of the JavaScript 'window' objects for each of the tabs running inside
// the current extension. If windowId is specified, returns only the 'window' objects of tabs attached to the specified window.
func GetExtensionTabs(windowId int) []windows.Window {
	wins := []windows.Window{}
	windowObjs := extension.Call("getExtensionTabs", windowId)
	for i := 0; i < windowObjs.Length(); i++ {
		window := windowObjs.Index(i)
		wins = append(wins, windows.Window{Object: *window})
	}
	return wins
}

// IsAllowedIncognitoAccess retrieves the state of the extension's access to Incognito-mode
// (as determined by the user-controlled 'Allowed in Incognito' checkbox.
func IsAllowedIncognitoAccess(callback func(isAllowedAccess bool)) {
	extension.Call("isAllowedIncognitoAccess", callback)
}

// IsAllowedFileSchemeAccess retrieves the state of the extension's access to the 'file://'
// scheme (as determined by the user-controlled 'Allow access to File URLs' checkbox.
func IsAllowedFileSchemeAccess(callback func(isAllowedAccess bool)) {
	extension.Call("isAllowedFileSchemeAccess", callback)
}

// SetUpdateUrlData sets the value of the ap CGI parameter used in the extension's update URL.
// This value is ignored for extensions that are hosted in the Chrome Extension Gallery.
func SetUpdateUrlData(data string) {
	extension.Call("setUpdateUrlData", data)
}
