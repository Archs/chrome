package webstore

import (
	"github.com/Archs/chrome"
)

var (
	webstore = chrome.Get("webstore")
)

/*
* Methods
 */

// Install instals an app from the chrome store
func Install(url string, successCallback func(), failureCallback func(err string, errorCode string)) {
	webstore.Call("install", url, successCallback, failureCallback)
}

/*
* Events
 */

// OnInstallStageChanged fired when an inline installation enters a new InstallStage. In order to receive
// notifications about this event, listeners must be registered before the inline installation begins.
func OnInstallStageChanged(callback func(stage string)) {
	webstore.Get("onInstallStageChanged").Call("addListener", callback)
}

// OnDownloadProgress fired periodically with the download progress of an inline install. In order to
// receive notifications about this event, listeners must be registered before the inline installation begins.
func OnDownloadProgress(callback func(percentDownloaded int64)) {
	webstore.Get("onDownloadProgress").Call("addListener", callback)
}
