package permissions

import (
	"github.com/Archs/chrome"
)

var (
	perms = chrome.Get("permissions")
)

/*
* Methods
 */

// GetAll gets the extension's current set of perms.
func GetAll(callback func(permissions map[string][]string)) {
	perms.Call("getAll", callback)
}

// Contains checks if the extension has the specified perms.
func Contains(permissions map[string][]string, callback func(result bool)) {
	perms.Call("contains", permissions, callback)
}

// Request requests access to the specified perms. These permissions must be defined in the
// optional_permissions field of the manifest. If there are any problems requesting the permissions, runtime.lastError will be set.
func Request(permissions map[string][]string, callback func(granted bool)) {
	perms.Call("request", permissions, callback)
}

// Remove removes access to the specified perms. If there are any problems removing the permissions, runtime.lastError will be set.
func Remove(permissions map[string][]string, callback func(removed bool)) {
	perms.Call("remove", permissions, callback)
}

/*
* Events
 */

// OnAdded fired when the extension acquires new perms.
func OnAdded(callback func(permissions map[string][]string)) {
	perms.Get("onAdded").Call("addListener", callback)
}

// OnRemoved fired when access to permissions has been removed from the extension.
func OnRemoved(callback func(permissions map[string][]string)) {
	perms.Get("onRemoved").Call("addListener", callback)
}
