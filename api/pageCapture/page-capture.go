package pageCapture

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	pageCapture = chrome.Get("pageCapture")
)

/*
* Methods
 */

// SaveAsMHTML saves the content of the tab with given id as MHTML.
func SaveAsMHTML(details js.M, callback func(mhtmlData interface{})) {
	pageCapture.Call("saveAsMHTML", details, callback)
}
