package chrome

import (
	"github.com/gopherjs/gopherjs/js"
)

type AppRuntime struct {
	o *js.Object
}

type EmbedRequest struct {
	*js.Object

	// string	embedderId
	EmbedderId string `js:"embedderId"`

	// 	any	(optional) data
	// Optional developer specified data that the app to be embedded can use when making an embedding decision.
	Any js.Object `js:"data"`

	// function	allow
	// Allows embedderId to embed this app in an <appview> element.
	// Parameters
	// string	url
	// The url specifies the content to embed.
	Allow func(url string) `js:"allow"`

	// function	deny
	// Prevents embedderId from embedding this app in an <appview> element.
	Deny func() `js:"deny"`
}

// onEmbedRequested Fired when an embedding app requests to embed this app.
// This event is only available on dev channel with the flag --enable-app-view.
func (a *AppRuntime) OnEmbedRequested(callback func(*EmbedRequest)) {
	a.o.Get("onEmbedRequested").Call("addListener", callback)
}

type LaunchDataSource string

const (
	APP_LAUNCHER    = "app_launcher"
	NEW_TAB_PAGE    = "new_tab_page"
	RELOAD          = "reload"
	RESTART         = "restart"
	LOAD_AND_LAUNCH = "load_and_launch"
	COMMAND_LINE    = "command_line"
	FILE_HANDLER    = "file_handler"
	URL_HANDLER     = "url_handler"
	SYSTEM_TRAY     = "system_tray"
	ABOUT_PAGE      = "about_page"
	KEYBOARD        = "keyboard"
	EXTENSIONS_PAGE = "extensions_page"
	MANAGEMENT_API  = "management_api"
	EPHEMERAL_APP   = "ephemeral_app"
	BACKGROUND      = "background"
	KIOSK           = "kiosk"
	CHROME_INTERNAL = "chrome_internal"
	TEST            = "test"
)

type LaunchData struct {
	*js.Object
	// 	string	(optional) id
	// Since Chrome 25.
	// The ID of the file or URL handler that the app is being invoked with.
	//Handler IDs are the top-level keys in the file_handlers and/or url_handlers dictionaries in the manifest.
	Id string

	// array of object	(optional) items
	// Since Chrome 25.

	// The file entries for the onLaunched event triggered by a matching file handler in the file_handlers manifest key.

	// Properties of each object
	// FileEntry	entry
	// FileEntry for the file.

	// string	type
	// The MIME type of the file.

	// string	(optional) url
	// Since Chrome 31.
	// The URL for the onLaunched event triggered by a matching URL handler in the url_handlers manifest key.
	Url string

	// string	(optional) referrerUrl
	// Since Chrome 31.
	// The referrer URL for the onLaunched event triggered by a matching URL handler in the url_handlers manifest key.
	ReferrerUrl string

	// boolean	(optional) isKioskSession
	// Since Chrome 31.
	// Whether the app is being launched in a Chrome OS kiosk session.
	IsKioskSession bool

	// Where the app is launched from. Since Chrome 40.
	Source LaunchDataSource
}

// onLaunched Fired when an app is launched from the launcher.
func (a *AppRuntime) OnLaunched(callback func(*LaunchData)) {
	a.o.Get("onLaunched").Call("addListener", callback)
}

// Since Chrome 24.
//
// onRestarted Fired at Chrome startup to apps that were running when Chrome last shut down,
// or when apps have been requested to restart from their previous state for other reasons
// (e.g. when the user revokes access to an app's retained files the runtime will restart the app).
// In these situations if apps do not have an onRestarted handler they will be sent an onLaunched event instead.
func (a *AppRuntime) OnRestarted(callback func()) {
	a.o.Get("onRestarted").Call("addListener", callback)
}
