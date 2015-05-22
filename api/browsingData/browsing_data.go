package browsingData

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	browsingData = chrome.Get("browsingData")
)

/*
* Types
 */

type RemovalOptions struct {
	*js.Object
	Since       float64         `js:"since"`
	OriginTypes map[string]bool `js:"originTypes"`
}

type DataTypeSet map[string]bool

/*
* Methods:
 */

// Settings reports which types of data are currently selected in the 'Clear browsing data'
// settings UI. Note: some of the data types included in this API are not available in the
// settings UI, and some UI settings control more than one data type listed here.
func Settings(callback func(result js.M)) {
	browsingData.Call("settings", callback)
}

// Remove clears various types of browsing data stored in a user's profile.
func Remove(options RemovalOptions, dataToRemove DataTypeSet, callback func()) {
	browsingData.Call("remove", options, dataToRemove, callback)
}

// RemoveAppCache clears websites' appcache data.
func RemoveAppCache(options RemovalOptions, callback func()) {
	browsingData.Call("removeAppCache", options, callback)
}

// RemoveCache clears the browser's cache.
func RemoveCache(options RemovalOptions, callback func()) {
	browsingData.Call("removeCache", options, callback)
}

// RemoveCookies clears the browser's cookies and server-bound certificates modified within a particular timeframe.
func RemoveCookies(options RemovalOptions, callback func()) {
	browsingData.Call("removeCookies", options, callback)
}

// RemoveDownloads clears the browser's list of downloaded files (not the downloaded files themselves).
func RemoveDownloads(options RemovalOptions, callback func()) {
	browsingData.Call("removeDownloads", options, callback)
}

// RemoveFileSystems clears websites' file system data.
func RemoveFileSystems(options RemovalOptions, callback func()) {
	browsingData.Call("removeFileSystems", options, callback)
}

// RemoveFormData clears the browser's stored form data (autofill).
func RemoveFormData(options RemovalOptions, callback func()) {
	browsingData.Call("removeFormData", options, callback)
}

// RemoveHistory clears the browser's history.
func RemoveHistory(options RemovalOptions, callback func()) {
	browsingData.Call("removeHistory", options, callback)
}

// RemoveIndexedDB clears websites' IndexedDB data.
func RemoveIndexedDB(options RemovalOptions, callback func()) {
	browsingData.Call("removeIndexedDB", options, callback)
}

// RemoveLocalStorage clears websites' local storage data.
func RemoveLocalStorage(options RemovalOptions, callback func()) {
	browsingData.Call("removeLocalStorage", options, callback)
}

// RemovePluginData clears plugins' data.
func RemovePluginData(options RemovalOptions, callback func()) {
	browsingData.Call("removePluginData", options, callback)
}

// RemovePasswords clears the browser's stored passwords.
func RemovePasswords(options RemovalOptions, callback func()) {
	browsingData.Call("removePasswords", options, callback)
}

// RemoveWebSQL clears websites' WebSQL data.
func RemoveWebSQL(options RemovalOptions, callback func()) {
	browsingData.Call("removeWebSQL", options, callback)
}
