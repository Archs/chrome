package downloads

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	downloads = chrome.Get("downloads")
)

/*
* Types
 */

type DownloadItem struct {
	*js.Object
	Id               int    `js:"id"`
	Url              string `js:"url"`
	Referrer         string `js:"referrer"`
	Filename         string `js:"filename"`
	Incognito        bool   `js:"incognito"`
	Danger           string `js:"danger"`
	Mime             string `js:"mime"`
	StartTime        string `js:"startTime"`
	EndTime          string `js:"endTime"`
	EstimatedEndTime string `js:"estimatedEndTime"`
	State            string `js:"state"`
	Paused           bool   `js:"paused"`
	CanResume        bool   `js:"canResume"`
	Error            string `js:"error"`
	BytesReceived    int64  `js:"bytesReceived"`
	TotalBytes       int64  `js:"totalBytes"`
	FileSize         int64  `js:"fileSize"`
	Exists           bool   `js:"exists"`
	ByExtensionId    string `js:"byExtensionId"`
	ByExtensionName  string `js:"byExtensionName"`
}

type StringDelta map[string]string

type DoubleDelta map[string]int64

type BooleanDelta map[string]bool

/*
* Methods:
 */

/* Download download a URL. If the URL uses the HTTP[S] protocol, then the request will include all cookies currently set for its hostname. If both filename and saveAs are specified, then the Save As dialog will be displayed, pre-populated with the specified filename. If the download started successfully, callback will be called with the new DownloadItem's downloadId. If there was an error starting the download, then callback will be called with downloadId=undefined and runtime.lastError will contain a descriptive string. The error strings are not guaranteed to remain backwards compatible between releases. Extensions must not parse it. */
func Download(options js.M, callback func(downloadId int)) {
	downloads.Call("download", options, callback)
}

/* Search find DownloadItem. Set query to the empty js.M to get all DownloadItem. To get a specific DownloadItem, set only the id field. To page through a large number of items, set orderBy: ['-startTime'], set limit to the number of items per page, and set startedAfter to the startTime of the last item from the last page. */
func Search(query js.M, callback func(results []DownloadItem)) {
	downloads.Call("search", query, callback)
}

/* Pause pause the download. If the request was successful the download is in a paused state. Otherwise runtime.lastError contains an error message. The request will fail if the download is not active. */
func Pause(downloadId int, callback func()) {
	downloads.Call("pause", downloadId, callback)
}

/* Resume resume a paused download. If the request was successful the download is in progress and unpaused. Otherwise runtime.lastError contains an error message. The request will fail if the download is not active. */
func Resume(downloadId int, callback func()) {
	downloads.Call("resume", downloadId, callback)
}

/* Cancel cancel a download. When callback is run, the download is cancelled, completed, interrupted or doesn't exist anymore. */
func Cancel(downloadId int, callback func()) {
	downloads.Call("cancel", downloadId, callback)
}

/* GetFileIcon retrieve an icon for the specified download. For new downloads, file icons are available after the onCreated event has been received. The image returned by this function while a download is in progress may be different from the image returned after the download is complete. Icon retrieval is done by querying the underlying operating system or toolkit depending on the platform. The icon that is returned will therefore depend on a number of factors including state of the download, platform, registered file types and visual theme. If a file icon cannot be determined, runtime.lastError will contain an error message. */
func GetFileIcon(downloadId int, options js.M, callback func(iconURL string)) {
	downloads.Call("getFileIcon", downloadId, options, callback)
}

/* Open open the downloaded file now if the DownloadItem is complete; otherwise returns an error through runtime.lastError. Requires the "downloads.open" permission in addition to the "downloads" permission. An onChanged event will fire when the item is opened for the first time. */
func Open(downloadId int) {
	downloads.Call("open", downloadId)
}

/* Show show the downloaded file in its folder in a file manager. */
func Show(downloadId int) {
	downloads.Call("show", downloadId)
}

/* ShowDefaultFolder show the default Downloads folder in a file manager. */
func ShowDefaultFolder() {
	downloads.Call("showDefaultFolder")
}

/* Erase erases matching DownloadItem from history without deleting the downloaded file. An onErased event will fire for each DownloadItem that matches query, then callback will be called. */
func Erase(query js.M, callback func(erasedIds []int)) {
	downloads.Call("erase", query, callback)
}

/* RemoveFile remove the downloaded file if it exists and the DownloadItem is complete; otherwise return an error through runtime.lastError. */
func RemoveFile(downloadId int, callback func()) {
	downloads.Call("removeFile", downloadId, callback)
}

/* AcceptDanger prompt the user to accept a dangerous download. Does not automatically accept dangerous downloads. If the download is accepted, then an onChanged event will fire, otherwise nothing will happen. When all the data is fetched into a temporary file and either the download is not dangerous or the danger has been accepted, then the temporary file is renamed to the target filename, the |state| changes to 'complete', and onChanged fires. */
func AcceptDanger(downloadId int, callback func()) {
	downloads.Call("acceptDanger", downloadId, callback)
}

/* Drag initiate dragging the downloaded file to another application. Call in a javascript ondragstart handler. */
func Drag(downloadId int) {
	downloads.Call("drag", downloadId)
}

/* SetShelfEnabled enable or disable the gray shelf at the bottom of every window associated with the current browser profile. The shelf will be disabled as long as at least one extension has disabled it. Enabling the shelf while at least one other extension has disabled it will return an error through runtime.lastError. Requires the "downloads.shelf" permission in addition to the "downloads" permission. */
func SetShelfEnabled(enabled bool) {
	downloads.Call("setShelfEnabled", enabled)
}

/*
* Events
 */

/* OnCreated this event fires with the DownloadItem js.M when a download begins. */
func OnCreated(callback func(downloadItem DownloadItem)) {
	downloads.Get("onCreated").Call("addListener", callback)
}

/* OnErased fires with the downloadId when a download is erased from history. */
func OnErased(callback func(downloadId int)) {
	downloads.Get("onErased").Call("addListener", callback)
}

/* OnChanged when any of a DownloadItem's properties except bytesReceived and estimatedEndTime changes, this event fires with the downloadId and an js.M containing the properties that changed. */
func OnChanged(callback func(downloadDelta js.M)) {
	downloads.Get("onChanged").Call("addListener", callback)
}

/* OnDeterminingFilename during the filename determination process, extensions will be given the opportunity to override the target DownloadItem.filename. Each extension may not register more than one listener for this event. Each listener must call suggest exactly once, either synchronously or asynchronously. If the listener calls suggest asynchronously, then it must return true. If the listener neither calls suggest synchronously nor returns true, then suggest will be called automatically. The DownloadItem will not complete until all listeners have called suggest. Listeners may call suggest without any arguments in order to allow the download to use downloadItem.filename for its filename, or pass a suggestion js.M to suggest in order to override the target filename. If more than one extension overrides the filename, then the last extension installed whose listener passes a suggestion js.M to suggest wins. In order to avoid confusion regarding which extension will win, users should not install extensions that may conflict. If the download is initiated by download and the target filename is known before the MIME type and tentative filename have been determined, pass filename to download instead. */
func OnDeterminingFilename(callback func(downloadItem DownloadItem, suggest func(suggestion js.M))) {
	downloads.Get("onDeterminingFilename").Call("addListener", callback)
}
