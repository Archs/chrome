package power

import (
	"github.com/Archs/chrome"
)

var (
	power = chrome.Get("power")
)

/*
* Methods
 */

// RequestKeepAwake requests that power management be temporarily disabled. |level| describes the
//degree to which power management should be disabled. If a request previously made by the same
//app is still active, it will be replaced by the new request.
func RequestKeepAwake(level string) {
	power.Call("requestKeepAwake", level)
}

// ReleaseKeepAwake releases a request previously made via requestKeepAwake().
func ReleaseKeepAwake() {
	power.Call("releaseKeepAwake")
}
