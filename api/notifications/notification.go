package notifications

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	notifications = chrome.Get("notifications")
)

/*
* Types
 */

type NotificationOptions struct {
	*js.Object
	Type           string `js:"type"`
	IconUrl        string `js:"iconUrl"`
	AppIconMaskUrl string `js:"appIconMaskUrl"`
	Title          string `js:"title"`
	Message        string `js:"message"`
	ContextMessage string `js:"contextMessage"`
	Priority       int    `js:"priority"`
	EventTime      int64  `js:"eventTime"`
	Buttons        []js.M `js:"buttons"`
	ImageUrl       string `js:"imageUrl"`
	Items          []js.M `js:"items"`
	Progress       int    `js:"progress"`
	IsClickable    bool   `js:"isClickable"`
}

/*
* Methods
 */

// Create creates and displays a notification.
func Create(notificationId string, options NotificationOptions, callback func(notificationId string)) {
	notifications.Call("create", notificationId, options, callback)
}

// Update updates an existing notification.
func Update(notificationId string, options NotificationOptions, callback func(wasUpdated bool)) {
	notifications.Call("update", notificationId, options, callback)
}

// Clear clears the specified notification.
func Clear(notificationId string, callback func(notificationId string)) {
	notifications.Call("clear", notificationId, callback)
}

// GetAll retrieves all the notifications.
func GetAll(callback func(notifications js.M)) {
	notifications.Call("getAll", callback)
}

// GetPermissionLevel retrieves whether the user has enabled notifications from this app or extension.
func GetPermissionLevel(callback func(level string)) {
	notifications.Call("getPermissionLevel", callback)
}

/*
* Events
 */

// OnClosed the notification closed, either by the system or by user action.
func OnClosed(callback func(notificationId string, byUser bool)) {
	notifications.Get("onClosed").Call("addListener", callback)
}

// OnClicked the user clicked in a non-button area of the notification.
func OnClicked(callback func(notificationId string)) {
	notifications.Get("onClicked").Call("addListener", callback)
}

// OnButtonClicked the user pressed a button in the notification.
func OnButtonClicked(callback func(notificationId string, buttonIndex int)) {
	notifications.Get("onButtonClicked").Call("addListener", callback)
}

// OnPermissionLevelChanged the user changes the permission level.
func OnPermissionLevelChanged(callback func(level string)) {
	notifications.Get("onPermissionLevelChanged").Call("addListener", callback)
}

// OnShowSettings the user clicked on a link for the app's notification settings.
func OnShowSettings(callback func()) {
	notifications.Get("onShowSettings").Call("addListener", callback)
}
