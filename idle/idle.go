package idle

import (
	"github.com/Archs/chrome"
)

var (
	idle = chrome.Get("idle")
)

/*
* Methods
 */

// QueryState returns "locked" if the system is locked, "idle" if the user has not generated any
// input for a specified number of seconds, or "active" otherwise.
func QueryState(detectionIntervalInSeconds int, callback func(newState string)) {
	idle.Call("queryState", detectionIntervalInSeconds, callback)
}

// SetDetectionInterval sets the interval, in seconds, used to determine when the system is in an idle
//  state for onStateChanged events. The default interval is 60 seconds.
func SetDetectionInterval(intervalInSeconds int) {
	idle.Call("setDetectionInterval", intervalInSeconds)
}

/*
* Events
 */

// OnStateChanged fired when the system changes to an active, idle or locked state. The event fires with
// "locked" if the screen is locked or the screensaver activates, "idle" if the system is unlocked and the
// user has not generated any input for a specified number of seconds, and "active" when the user generates
// input on an idle system.
func OnStateChanged(callback func(newState string)) {
	idle.Get("onStateChanged").Call("addListener", callback)
}
