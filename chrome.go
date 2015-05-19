package chrome

import "github.com/gopherjs/gopherjs/js"

var (
	chrome = js.Global.Get("chrome")
)

type Object js.M

func Get(key string) *js.Object {
	return chrome.Get(key)
}

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
