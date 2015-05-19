package chrome

import (
	"github.com/Archs/chrome"

	"github.com/gopherjs/gopherjs/js"
)

var (
	alarms = chrome.Get("alarms")
)

/*
* Types
 */

type Alarm struct {
	*js.Object
	Name            string `js:"name"`
	ScheduledTime   string `js:"scheduledTime"`
	PeriodInMinutes string `js:"periodInMinutes"`
}

/*
* Methods:
 */

// Create creates an alarm. Near the time(s) specified by alarmInfo, the onAlarm event is fired.
// If there is another alarm with the same name (or no name if none is specified), it will be
// cancelled and replaced by this alarm.
// You must use time.Now().UnixNano() for "when" timestamp in alarmInfo for this to work.
func Create(name string, alarmInfo Object) {
	alarms.Call("create", name, alarmInfo)
}

// Get retrieves details about the specified alarm.
func Get(name string, callback func(alarm Alarm)) {
	alarms.Call("get", name, callback)
}

// GetAll gets an array of all the alarms.
func GetAll(callback func(alarms []Alarm)) {
	alarms.Call("getAll", callback)
}

// Clear clears the alarm with the given name.
func Clear(name string, callback func(wasCleared bool)) {
	alarms.Call("clear", name, callback)
}

// ClearAll clears all alarms.
func ClearAll(callback func(wasCleared bool)) {
	alarms.Call("clearAll", callback)
}

/*
* Events
 */
// OnAlarm is fired when an alarm has elapsed. Useful for event pages.
func OnAlarm(callback func(alarm Alarm)) {
	alarms.Get("onAlarm").Call("addListener", callback)
}
