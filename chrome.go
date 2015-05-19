package chrome

import "github.com/gopherjs/gopherjs/js"

var (
	chrome = js.Global.Get("chrome")
)

type Object js.M

func Get(key string) *js.Object {
	return chrome.Get(key)
}
