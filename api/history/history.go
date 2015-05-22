package history

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	history = chrome.Get("history")
)

/*
* Types
 */

type HistoryItem struct {
	*js.Object
	Id            string `js:"id"`
	Url           string `js:"url"`
	Title         string `js:"title"`
	LastVisitTime int64  `js:"lastVisitTime"`
	VisitCount    int    `js:"visitCount"`
	TypedCount    int    `js:"typedCount"`
}

type VisitItem struct {
	*js.Object
	Id               string `js:"id"`
	VisitId          string `js:"visitId"`
	VisitTime        int64  `js:"visitTime"`
	ReferringVisitId string `js:"referringVisitId"`
	Transition       string `js:"transition"`
}

/*
* Methods:
 */

// Search searches the history for the last visit time of each page matching the query.
func Search(query js.M, callback func(results []HistoryItem)) {
	history.Call("search", query, callback)
}

// GetVisits retrieves information about visits to a URL.
func GetVisits(details js.M, callback func(results []VisitItem)) {
	history.Call("getVisits", details, callback)
}

// AddUrl adds a URL to the history at the current time with a transition type of "link".
func AddUrl(details js.M, callback func()) {
	history.Call("addUrl", details, callback)
}

// DeleteUrl removes all occurrences of the given URL from the history.
func DeleteUrl(details js.M, callback func()) {
	history.Call("deleteUrl", details, callback)
}

// DeleteRange removes all items within the specified date range from the history.
// Pages will not be removed from the history unless all visits fall within the range.
func DeleteRange(rang js.M, callback func()) {
	history.Call("deleteRange", rang, callback)
}

// DeleteAll deletes all items from the history.
func DeleteAll(callback func()) {
	history.Call("deleteAll", callback)
}

/*
* Events
 */

// OnVisited fired when a URL is visited, providing the HistoryItem data for that URL.
// This event fires before the page has loaded.
func OnVisited(callback func(result HistoryItem)) {
	history.Get("onVisited").Call("addListener", callback)
}

// OnVisitedRemoved fired when one or more URLs are removed from the history service.
// When all visits have been removed the URL is purged from history.
func OnVisitedRemoved(callback func(removed js.M)) {
	history.Get("onVisitedRemoved").Call("addListener", callback)
}
