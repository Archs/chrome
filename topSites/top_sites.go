package topSites

import (
	"github.com/Archs/chrome"
)

var (
	topSites = chrome.Get("topSites")
)

/*
* Types
 */

type MostvisitedURL struct {
	Url   string `js:"url"`
	Title string `js:"title"`
}

/*
* Methods
 */

// Get gets a list of top sites.
func Get(callback func(data []MostvisitedURL)) {
	topSites.Call("get", callback)
}
