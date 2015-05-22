package cookies

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	cookies = chrome.Get("cookies")
)

/*
* Types
 */

type Cookie struct {
	*js.Object
	Name           string `js:"name"`
	Value          string `js:"value"`
	Domain         string `js:"domain"`
	HostOnly       bool   `js:"hostOnly"`
	Path           string `js:"path"`
	Secure         bool   `js:"secure"`
	HttpOnly       bool   `js:"httpOnly"`
	Session        bool   `js:"session"`
	ExpirationDate int64  `js:"expirationDate"`
	StoreId        string `js:"storeId"`
}

type CookieStore struct {
	*js.Object
	Id     string `js:"id"`
	TabIds []int  `js:"tabIds"`
}

/*
* Methods:
 */

// Get retrieves information about a single cookie. If more than one cookie of the same name
// exists for the given URL, the one with the longest path will be returned. For cookies with
// the same path length, the cookie with the earliest creation time will be returned.
func Get(details js.M, callback func(cookie Cookie)) {
	cookies.Call("get", details, callback)
}

// GetAll retrieves all cookies from a single cookie store that match the given information.
// The cookies returned will be sorted, with those with the longest path first. If multiple cookies
// have the same path length, those with the earliest creation time will be first.
func GetAll(details js.M, callback func(cookies []Cookie)) {
	cookies.Call("getAll", details, callback)
}

// Set sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
func Set(details js.M, callback func(cookie Cookie)) {
	cookies.Call("set", details, callback)
}

// Remove deletes a cookie by name.
func Remove(details js.M, callback func(details js.M)) {
	cookies.Call("remove", details, callback)
}

// GetAllCookieStores lists all existing cookie stores.
func GetAllCookieStores(callback func(cookieStores []CookieStore)) {
	cookies.Call("getAllCookieStores", callback)
}

/*
* Events
 */

// OnChanged fired when a cookie is set or removed. As a special case, note that updating a cookie's properties
//  is implemented as a two step process: the cookie to be updated is first removed entirely, generating a
// notification with "cause" of "overwrite" . Afterwards, a new cookie is written with the updated values,
// generating a second notification with "cause" "explicit".
func OnChanged(callback func(changeInfo js.M)) {
	cookies.Get("onChanged").Call("addListener", callback)
}
