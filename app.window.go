package chrome

import (
	"github.com/gopherjs/gopherjs/js"
)

type AppWindowManager struct {
	o *js.Object
}

type AppWindow struct {
	o *js.Object
}

// Since Chrome 35.
type BoundsSpecification struct {
	*js.Object
	// integer	(optional) left
	// The X coordinate of the content or window.
	Left int `js:"left"`

	// integer	(optional) top
	// The Y coordinate of the content or window.
	Top int `js:"top"`

	// integer	(optional) width
	// The width of the content or window.
	Width int `js:"width"`

	// integer	(optional) height
	// The height of the content or window.
	Height int `js:"height"`

	// integer	(optional) minWidth
	// The minimum width of the content or window.
	MinWidth int `js:"minWidth"`

	// integer	(optional) minHeight
	// The minimum height of the content or window.
	MinHeigth int `js:"minHeight"`

	// integer	(optional) maxWidth
	// The maximum width of the content or window.
	MaxWidth int `js:"maxWidth"`

	// integer	(optional) maxHeight
	// The maximum height of the content or window.
	MaxHeigth int `js:"maxHeight"`
}

// Since Chrome 35.
type ContentBounds struct {
	*js.Object
	// integer	(optional) left
	Left int `js:"left"`
	// integer	(optional) top
	Top int `js:"top"`
	// integer	(optional) width
	Width int `js:"width"`
	// integer	(optional) height
	Height int `js:"height"`
}

type CreateWindowOptions struct {
	*js.Object
	// string	(optional) id
	// Id to identify the window. This will be used to remember the size and position of the window and restore that geometry when a window with the same id is later opened. If a window with a given id is created while another window with the same id already exists, the currently opened window will be focused instead of creating a new window.
	Id string `js:"id"`

	// BoundsSpecification	(optional) innerBounds
	// Since Chrome 35.
	// Used to specify the initial position, initial size and constraints of the window's content (excluding window decorations). If an id is also specified and a window with a matching id has been shown before, the remembered bounds will be used instead.
	// Note that the padding between the inner and outer bounds is determined by the OS. Therefore setting the same bounds property for both the innerBounds and outerBounds will result in an error.
	// This property is new in Chrome 36.
	InnerBounds BoundsSpecification `js:"innerBounds"`

	// BoundsSpecification	(optional) outerBounds
	// Since Chrome 35.
	// Used to specify the initial position, initial size and constraints of the window (including window decorations such as the title bar and frame). If an id is also specified and a window with a matching id has been shown before, the remembered bounds will be used instead.
	// Note that the padding between the inner and outer bounds is determined by the OS. Therefore setting the same bounds property for both the innerBounds and outerBounds will result in an error.
	// This property is new in Chrome 36.
	OuterBounds BoundsSpecification `js:"outerBounds"`

	// integer	(optional) minWidth
	// Deprecated since Chrome 36. Use innerBounds or outerBounds.

	// Minimum width of the window.

	// integer	(optional) minHeight
	// Deprecated since Chrome 36. Use innerBounds or outerBounds.

	// Minimum height of the window.

	// integer	(optional) maxWidth
	// Deprecated since Chrome 36. Use innerBounds or outerBounds.

	// Maximum width of the window.

	// integer	(optional) maxHeight
	// Deprecated since Chrome 36. Use innerBounds or outerBounds.

	// Maximum height of the window.

	// string or FrameOptions	(optional) frame
	// Frame type: none or chrome (defaults to chrome). For none, the -webkit-app-region CSS property can be used to apply draggability to the app's window. -webkit-app-region: drag can be used to mark regions draggable. no-drag can be used to disable this style on nested elements.

	// Use of FrameOptions is new in M36.

	// ContentBounds	(optional) bounds
	// Deprecated since Chrome 36. Use innerBounds or outerBounds.
	// Size and position of the content in the window (excluding the titlebar). If an id is also specified and a window with a matching id has been shown before, the remembered bounds of the window will be used instead.
	Bounds ContentBounds `js:"bounds"`

	// enum of "normal", "fullscreen", "maximized", or "minimized"	(optional) state
	// Since Chrome 28.
	// The initial state of the window, allowing it to be created already fullscreen, maximized, or minimized. Defaults to 'normal'.
	State string `js:"state"`

	// boolean	(optional) hidden
	// Since Chrome 24.
	// If true, the window will be created in a hidden state. Call show() on the window to show it once it has been created. Defaults to false.
	Hidden bool `js:"hidden"`

	// boolean	(optional) resizable
	// Since Chrome 27.
	// If true, the window will be resizable by the user. Defaults to true.
	Resizable bool `js:"resizable"`

	// // boolean	(optional) singleton
	// // Deprecated since Chrome 34. Multiple windows with the same id is no longer supported.

	// // By default if you specify an id for the window, the window will only be created if another window with the same id doesn't already exist. If a window with the same id already exists that window is activated instead. If you do want to create multiple windows with the same id, you can set this property to false.
	// Singleton bool

	// boolean	(optional) alwaysOnTop
	// Since Chrome 32.
	// If true, the window will stay above most other windows. If there are multiple windows of this kind, the currently focused window will be in the foreground. Requires the alwaysOnTopWindows permission. Defaults to false.
	// Call setAlwaysOnTop() on the window to change this property after creation.
	AlwaysOnTop bool `js:"alwaysOnTop"`

	// boolean	(optional) focused
	// Since Chrome 33.
	// If true, the window will be focused when created. Defaults to true.
	Focused bool `js:"focused"`

	// boolean	(optional) visibleOnAllWorkspaces
	// Since Chrome 39.
	// If true, the window will be visible on all workspaces.
	VisibleOnAllWorkspaces bool `js:"visibleOnAllWorkspaces"`
}

// chrome.app.window.create(string url, CreateWindowOptions options, function callback)
// The size and position of a window can be specified in a number of different ways. The most simple option is not specifying anything at all, in which case a default size and platform dependent position will be used.
//
// To set the position, size and constraints of the window, use the innerBounds or outerBounds properties. Inner bounds do not include window decorations. Outer bounds include the window's title bar and frame. Note that the padding between the inner and outer bounds is determined by the OS. Therefore setting the same property for both inner and outer bounds is considered an error (for example, setting both innerBounds.left and outerBounds.left).
//
// To automatically remember the positions of windows you can give them ids. If a window has an id, This id is used to remember the size and position of the window whenever it is moved or resized. This size and position is then used instead of the specified bounds on subsequent opening of a window with the same id. If you need to open a window with an id at a location other than the remembered default, you can create it hidden, move it to the desired location, then show it.
func (a *AppWindowManager) Create(url string, options CreateWindowOptions, callback func(*AppWindow)) {
	a.o.Call("create", url, options, callback)
}

func (a *AppWindowManager) CreateM(url string, options Object, callback func(*AppWindow)) {
	a.o.Call("create", url, options, callback)
}

// function	focus
// Focus the window.
func (a *AppWindow) Focus() {
	a.o.Call("focus")
}

// function	fullscreen
// Since Chrome 28.
//
// Fullscreens the window.
func (a *AppWindow) Fullscreen() {
	a.o.Call("fullscreen")
}

// The user will be able to restore the window by pressing ESC. An application can prevent the fullscreen state to be left when ESC is pressed by requesting the app.window.fullscreen.overrideEsc permission and canceling the event by calling .preventDefault(), in the keydown and keyup handlers, like this:
// window.onkeydown = window.onkeyup = function(e) { if (e.keyCode == 27 /* ESC */) { e.preventDefault(); } };

// Note window.fullscreen() will cause the entire window to become fullscreen and does not require a user gesture. The HTML5 fullscreen API can also be used to enter fullscreen mode (see Web APIs for more details).

// function	isFullscreen
// Since Chrome 27.

// Is the window fullscreen? This will be true if the window has been created fullscreen or was made fullscreen via the AppWindow or HTML5 fullscreen APIs.

// Returns	boolean.
// function	minimize
// Minimize the window.

// function	isMinimized
// Since Chrome 25.

// Is the window minimized?

// Returns	boolean.
// function	maximize
// Maximize the window.

// function	isMaximized
// Since Chrome 25.

// Is the window maximized?

// Returns	boolean.
// function	restore
// Restore the window, exiting a maximized, minimized, or fullscreen state.

// function	moveTo
// Deprecated since Chrome 43. Use outerBounds.

// Move the window to the position (|left|, |top|).

// Parameters
// integer	left
// Since Chrome 25.

// integer	top
// Since Chrome 25.

// function	resizeTo
// Deprecated since Chrome 43. Use outerBounds.

// Resize the window to |width|x|height| pixels in size.

// Parameters
// integer	width
// integer	height
// function	drawAttention
// Since Chrome 24.

// Draw attention to the window.

// function	clearAttention
// Since Chrome 24.

// Clear attention to the window.

// function	close
// Since Chrome 24.

// Close the window.
func (a *AppWindow) Close() {
	a.o.Call("close")
}

// function	show
// Since Chrome 24.
// Show the window. Does nothing if the window is already visible. Focus the window if |focused| is set to true or omitted.
// Parameters
// boolean	(optional) focused
// Since Chrome 34.
func (a *AppWindow) Show(focused bool) {
	a.o.Call("show", focused)
}

// function	hide
// Since Chrome 24.
// Hide the window. Does nothing if the window is already hidden.
func (a *AppWindow) Hide() {
	a.o.Call("hide")
}

// function	getBounds
// Deprecated since Chrome 36. Use innerBounds or outerBounds.

// Get the window's inner bounds as a ContentBounds object.

// Returns	ContentBounds.
// function	setBounds
// Deprecated since Chrome 36. Use innerBounds or outerBounds.

// Set the window's inner bounds.

// Parameters
// ContentBounds	bounds
// function	isAlwaysOnTop
// Since Chrome 32.

// Is the window always on top?

// Returns	boolean.
// function	setAlwaysOnTop
// Since Chrome 32.

// Set whether the window should stay above most other windows. Requires the alwaysOnTopWindows permission.

// Parameters
// boolean	alwaysOnTop
// function	setVisibleOnAllWorkspaces
// Since Chrome 39.

// For platforms that support multiple workspaces, is this window visible on all of them?

// Parameters
// boolean	alwaysVisible
// function	setInterceptAllKeys
// Since Chrome 41.

// Set whether the window should get all keyboard events including system keys that are usually not sent. This is best-effort subject to platform specific constraints. Requires the "app.window.allKeys" permission. This is currently available only in dev channel on Windows.

// Parameters
// boolean	wantAllKeys
// Window	contentWindow
// The JavaScript 'window' object for the created child.

// string	id
// Since Chrome 33.

// The id the window was created with.

// Bounds	innerBounds
// Since Chrome 35.

// The position, size and constraints of the window's content, which does not include window decorations. This property is new in Chrome 36.

// Bounds	outerBounds
// Since Chrome 35.

// The position, size and constraints of the window, which includes window decorations, such as the title bar and frame. This property is new in Chrome 36.
