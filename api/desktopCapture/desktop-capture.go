package desktopCapture

import (
	"github.com/Archs/chrome"
	"github.com/Archs/chrome/api/tabs"
)

var (
	desktopCapture = chrome.Get("desktopCapture")
)

/*
* Methods:
 */

// ChooseDesktopMedia shows desktop media picker UI with the specified set of sources.
func ChooseDesktopMedia(sources []string, targetTab tabs.Tab, callback func(streamId string)) int {
	return desktopCapture.Call("chooseDesktopMedia", sources, targetTab, callback).Int()
}

// CancelChooseDesktopMedia hides desktop media picker dialog shown by chooseDesktopMedia().
func CancelChooseDesktopMedia(desktopMediaRequestId int) {
	desktopCapture.Call("cancelChooseDesktopMedia", desktopMediaRequestId)
}
