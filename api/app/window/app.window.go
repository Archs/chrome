package window

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	window = chrome.Get("app").Get("window")
)

type AppWindow struct {
	*js.Object
	// Window	contentWindow
	// The JavaScript 'window' object for the created child.
	ContentWindow *js.Object `js:"contentWindow"`

	// string	id
	// Since Chrome 33.
	// The id the window was created with.
	Id string `js:"id"`

	// Bounds	innerBounds
	// Since Chrome 35.
	// The position, size and constraints of the window's content, which does not include window decorations. This property is new in Chrome 36.
	InnerBounds *Bounds `js:"innerBounds"`

	// Bounds	outerBounds
	// Since Chrome 35.
	// The position, size and constraints of the window, which includes window decorations, such as the title bar and frame. This property is new in Chrome 36.
	OuterBounds *Bounds `js:"outerBounds"`
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

type Bounds struct {
	BoundsSpecification
}

// function	 setPosition
// Since Chrome 35.
//
// Set the left and top position of the content or window.
//
// Parameters
// integer	 left
// integer	 top
func (b *Bounds) SetPosition(left, top int) {
	b.Call("setPosition", left, top)
}

// function	 setSize
// Since Chrome 35.
//
// Set the width and height of the content or window.
//
// Parameters
// integer	 width
// integer	 height
func (b *Bounds) SetSize(width, height int) {
	b.Call("setSize", width, height)
}

// function	 setMinimumSize
// Since Chrome 35.
//
// Set the minimum size constraints of the content or window. The minimum width or height can be set to null to remove the constraint. A value of undefined will leave a constraint unchanged.
//
// Parameters
// integer	 minWidth
// integer	 minHeight
func (b *Bounds) SetMinimumSize(minWidth, minHeight int) {
	b.Call("setMinimumSize", minWidth, minHeight)
}

// function	 setMaximumSize
// Since Chrome 35.
//
// Set the maximum size constraints of the content or window. The maximum width or height can be set to null to remove the constraint. A value of undefined will leave a constraint unchanged.
//
// Parameters
// integer	 maxWidth
// integer	 maxHeight
func (b *Bounds) SetMaximumSize(minWidth, minHeight int) {
	b.Call("setMaximumSize", minWidth, minHeight)
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
func Create(url string, options CreateWindowOptions, callback func(*AppWindow)) {
	window.Call("create", url, options, callback)
}

func CreateM(url string, options js.M, callback func(*AppWindow)) {
	window.Call("create", url, options, callback)
}

func CreateEx(url string) {
	window.Call("create", url)
}

// current

// AppWindow chrome.app.window.current()
// Returns an AppWindow object for the current script context (ie JavaScript 'window' object).
// This can also be called on a handle to a script context for another page,
// for example: otherWindow.chrome.app.window.current().
func Current() *AppWindow {
	return &AppWindow{
		Object: window.Call("current"),
	}
}

// getAll
//
// array of AppWindow chrome.app.window.getAll()
// Since Chrome 33.
//
// Gets an array of all currently created app windows. This method is new in Chrome 33.
func AetAll() []*AppWindow {
	ws := window.Call("getAll")
	return ws.Interface().([]*AppWindow)
}

// get
//
// AppWindow chrome.app.window.get(string id)
// Since Chrome 33.
//
// Gets an AppWindow with the given id. If no window with the given id exists null is returned. This method is new in Chrome 33.
//
// Parameters
// string	 id
func Get(id string) *AppWindow {
	return &AppWindow{
		Object: window.Call("get", id),
	}
}

// canSetVisibleOnAllWorkspaces
//
// boolean chrome.app.window.canSetVisibleOnAllWorkspaces()
// Since Chrome 42.
//
// Does the current platform support windows being visible on all workspaces?
func CanSetVisibleOnAllWorkspaces() bool {
	return window.Call("canSetVisibleOnAllWorkspaces").Bool()
}

// function	focus
// Focus the window.
func (a *AppWindow) Focus() {
	a.Call("focus")
}

// function	fullscreen
// Since Chrome 28.
//
// Fullscreens the window.
// The user will be able to restore the window by pressing ESC. An application can prevent the fullscreen state to be left when ESC is pressed by requesting the app.window.fullscreen.overrideEsc permission and canceling the event by calling .preventDefault(), in the keydown and keyup handlers, like this:
// window.onkeydown = window.onkeyup = function(e) { if (e.keyCode == 27 /* ESC */) { e.preventDefault(); } };
//
// Note window.fullscreen() will cause the entire window to become fullscreen and does not require a user gesture. The HTML5 fullscreen API can also be used to enter fullscreen mode (see Web APIs for more details).
func (a *AppWindow) Fullscreen() {
	a.Call("fullscreen")
}

// function	isFullscreen
// Since Chrome 27.
//
// Is the window fullscreen?
// This will be true if the window has been created fullscreen or was made fullscreen via the AppWindow or HTML5 fullscreen APIs.
// Returns	boolean.
func (a *AppWindow) IsFullscreen() bool {
	return a.Call("fullscreen").Bool()
}

// function	minimize
// Minimize the window.
func (a *AppWindow) Minimize() {
	a.Call("minimize")
}

// function	isMinimized
// Since Chrome 25.
// Is the window minimized?
func (a *AppWindow) IsMinimized() bool {
	return a.Call("isMinimized").Bool()
}

// function	maximize
// Maximize the window.
func (a *AppWindow) Maximize() {
	a.Call("maximize")
}

// function	isMaximized
// Since Chrome 25.
// Is the window maximized?
// Returns	boolean.
func (a *AppWindow) IsMaximized() bool {
	return a.Call("isMaximized").Bool()
}

// function	restore
// Restore the window, exiting a maximized, minimized, or fullscreen state.
func (a *AppWindow) Restore() {
	a.Call("restore")
}

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
//
// Draw attention to the window.
func (a *AppWindow) DrawAttention() {
	a.Call("drawAttention")
}

// function	clearAttention
// Since Chrome 24.
//
// Clear attention to the window.
func (a *AppWindow) ClearAttention() {
	a.Call("clearAttention")
}

// function	close
// Since Chrome 24.

// Close the window.
func (a *AppWindow) Close() {
	a.Call("close")
}

// function	show
// Since Chrome 24.
// Show the window. Does nothing if the window is already visible. Focus the window if |focused| is set to true or omitted.
// Parameters
// boolean	(optional) focused
// Since Chrome 34.
func (a *AppWindow) Show(focused bool) {
	a.Call("show", focused)
}

// function	hide
// Since Chrome 24.
// Hide the window. Does nothing if the window is already hidden.
func (a *AppWindow) Hide() {
	a.Call("hide")
}

// function	getBounds
// Deprecated since Chrome 36. Use innerBounds or outerBounds.

// Get the window's inner bounds as a ContentBounds js.M.

// Returns	ContentBounds.
// function	setBounds
// Deprecated since Chrome 36. Use innerBounds or outerBounds.

// Set the window's inner bounds.

// Parameters
// ContentBounds	bounds

// function	isAlwaysOnTop
// Since Chrome 32.
//
// Is the window always on top?
// Returns	boolean.
func (a *AppWindow) IsAlwaysOnTop() bool {
	return a.Call("isAlwaysOnTop").Bool()
}

// function	setAlwaysOnTop
// Since Chrome 32.
//
// Set whether the window should stay above most other windows. Requires the alwaysOnTopWindows permission.
//
// Parameters
// boolean	alwaysOnTop
func (a *AppWindow) SetAlwaysOnTop(alwaysOnTop bool) {
	a.Call("alwaysOnTop", alwaysOnTop)
}

// function	setVisibleOnAllWorkspaces
// Since Chrome 39.
//
// For platforms that support multiple workspaces, is this window visible on all of them?
//
// Parameters
// boolean	alwaysVisible
func (a *AppWindow) SetVisibleOnAllWorkspaces(alwaysVisible bool) {
	a.Call("setVisibleOnAllWorkspaces", alwaysVisible)
}

// function	setInterceptAllKeys
// Since Chrome 41.
//
// Set whether the window should get all keyboard events including system keys that are usually not sent. This is best-effort subject to platform specific constraints. Requires the "app.window.allKeys" permission. This is currently available only in dev channel on Windows.
//
// Parameters
// boolean	wantAllKeys
func (a *AppWindow) SetInterceptAllKeys(wantAllKeys bool) {
	a.Call("setInterceptAllKeys", wantAllKeys)
}
