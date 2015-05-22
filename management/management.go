package management

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	management = chrome.Get("management")
)

/*
* Types
 */

type IconInfo struct {
	*js.Object
	Size int    `js:"size"`
	Url  string `js:"url"`
}

type ExtensionInfo struct {
	*js.Object
	Id                   string     `js:"id"`
	Name                 string     `js:"name"`
	ShortName            string     `js:"shortName"`
	Description          string     `js:"description"`
	Version              string     `js:"version"`
	MayDisable           bool       `js:"mayDisable"`
	Enabled              string     `js:"enabled"`
	DisabledReason       string     `js:"disabledReason"`
	IsApp                bool       `js:"isApp"`
	Type                 string     `js:"type"`
	AppLaunchUrl         string     `js:"appLaunchUrl"`
	HomepageUrl          string     `js:"homepageUrl"`
	UpdateUrl            string     `js:"updateUrl"`
	OfflineEnabled       bool       `js:"offlineEnabled"`
	OptionsUrl           string     `js:"optionsUrl"`
	Icons                []IconInfo `js:"icons"`
	Permissions          []string   `js:"permissions"`
	HostPermissions      []string   `js:"hostPermissions"`
	InstallType          string     `js:"installType"`
	LaunchType           string     `js:"launchType"`
	AvailableLaunchTypes []string   `js:"availableLaunchTypes"`
}

/*
* Methods
 */

// GetAll returns a list of information about installed extensions and apps.
func GetAll(callback func(result []ExtensionInfo)) {
	management.Call("getAll", callback)
}

// Get returns information about the installed extension, app, or theme that has the given ID.
func Get(id string, callback func(result ExtensionInfo)) {
	management.Call("get", id, callback)
}

// GetSelf returns information about the calling extension, app, or theme. Note:
// This function can be used without requesting the 'management' permission in the manifest.
func GetSelf(callback func(result ExtensionInfo)) {
	management.Call("getSelf", callback)
}

// GetPermissionWarningsById returns a list of permission warnings for the given extension id.
func GetPermissionWarningsById(id string, callback func(permissionWarnings []string)) {
	management.Call("getPermissionWarningsById", id, callback)
}

// GetPermissionWarningsByManifest returns a list of permission warnings for the given extension manifest
// string. Note: This function can be used without requesting the 'management' permission in the manifest.
func GetPermissionWarningsByManifest(manifestStr string, callback func(permissionWarnings []string)) {
	management.Call("getPermissionWarningsByManifest", manifestStr, callback)
}

// SetEnabled enables or disables an app or extension.
func SetEnabled(id, string, enabled bool, callback func()) {
	management.Call("setEnabled", id, enabled, callback)
}

// Uninstall uninstalls a currently installed app or extension.
func Uninstall(id string, options js.M, callback func()) {
	management.Call("uninstall", id, options, callback)
}

// UninstallSelf uninstalls the calling extension. Note: This function can be used without
// requesting the 'management' permission in the manifest.
func UninstallSelf(options js.M, callback func()) {
	management.Call("uninstallSelf", options, callback)
}

// LaunchApp launches an application.
func LaunchApp(id string, callback func()) {
	management.Call("launchApp", id, callback)
}

// CreateAppShortcut display options to create shortcuts for an app. On Mac, only packaged app shortcuts can be created.
func CreateAppShortcut(id string, callback func()) {
	management.Call("createAppShortcut", id, callback)
}

// SetLaunchType set the launch type of an app.
func SetLaunchType(id string, launchType string, callback func()) {
	management.Call("setLaunchType", id, launchType, callback)
}

// GenerateAppForLink generate an app for a URL. Returns the generated bookmark app.
func GenerateAppForLink(url string, title string, callback func(result ExtensionInfo)) {
	management.Call("generateAppForLink", url, title, callback)
}

/*
* Events
 */

// OnInstalled fired when an app or extension has been installed.
func OnInstalled(callback func(info ExtensionInfo)) {
	management.Get("onInstalled").Call("addListener", callback)
}

// OnUninstalled fired when an app or extension has been uninstalled.
func OnUninstalled(callback func(id string)) {
	management.Get("onUninstalled").Call("addListener", callback)
}

// OnEnabled fired when an app or extension has been enabled.
func OnEnabled(callback func(info ExtensionInfo)) {
	management.Get("onEnabled").Call("addListener", callback)
}

// OnDisabled fired when an app or extension has been disabled.
func OnDisabled(callback func(info ExtensionInfo)) {
	management.Get("onDisabled").Call("addListener", callback)
}
