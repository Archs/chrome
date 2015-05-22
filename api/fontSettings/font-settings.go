package fontSettings

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	fontSettings = chrome.Get("fontSettings")
)

/*
* Types
 */

type FontName struct {
	*js.Object
	FontId      string `js:"fontId"`
	DisplayName string `js:"displayName"`
}

/*
* Methods:
 */

// ClearFont clears the font set by this extension, if any.
func ClearFont(details js.M, callback func()) {
	fontSettings.Call("clearFont", details, callback)
}

// GetFont gets the font for a given script and generic font family.
func GetFont(details js.M, callback func(details js.M)) {
	fontSettings.Call("getFont", details, callback)
}

// SetFont sets the font for a given script and generic font family.
func SetFont(details js.M, callback func()) {
	fontSettings.Call("setFont", details, callback)
}

// GetFontList gets a list of fonts on the system.
func GetFontList(callback func(results []FontName)) {
	fontSettings.Call("getFontList", callback)
}

// ClearDefaultFontSize clears the default font size set by this extension, if any.
func ClearDefaultFontSize(details js.M, callback func()) {
	fontSettings.Call("clearDefaultFontSize", details, callback)
}

// GetDefaultFontSize gets the default font size.
func GetDefaultFontSize(details js.M, callback func(details js.M)) {
	fontSettings.Call("getDefaultFontSize", details, callback)
}

// SetDefaultFontSize sets the default font size.
func SetDefaultFontSize(details js.M, callback func()) {
	fontSettings.Call("setDefaultFontSize", details, callback)
}

// ClearDefaultFixedFontSize clears the default fixed font size set by this extension, if any.
func ClearDefaultFixedFontSize(details js.M, callback func()) {
	fontSettings.Call("clearDefaultFixedFontSize", details, callback)
}

// GetDefaultFixedFontSize gets the default size for fixed width fonts.
func GetDefaultFixedFontSize(details js.M, callback func(details js.M)) {
	fontSettings.Call("getDefaultFixedFontSize", details, callback)
}

// SetDefaultFixedFontSize sets the default size for fixed width fonts.
func SetDefaultFixedFontSize(details js.M, callback func()) {
	fontSettings.Call("setDefaultFixedFontSize", details, callback)
}

// ClearMinimumFontSize lears the minimum font size set by this extension, if any.
func ClearMinimumFontSize(details js.M, callback func()) {
	fontSettings.Call("clearMinimumFontSize", details, callback)
}

// GetMinimumFontSize gets the minimum font size.
func GetMinimumFontSize(details js.M, callback func(details js.M)) {
	fontSettings.Call("getMinimumFontSize", details, callback)
}

// SetMinimumFontSize sets the minimum font size.
func SetMinimumFontSize(details js.M, callback func()) {
	fontSettings.Call("setMinimumFontSize", details, callback)
}

/*
* Events
 */

// OnFontChanged fired when a font setting changes.
func OnFontChanged(callback func(details js.M)) {
	fontSettings.Get("onFontChanged").Call("addListener", callback)
}

// OnDefaultFontSizeChanged fired when the default font size setting changes.
func OnDefaultFontSizeChanged(callback func(details js.M)) {
	fontSettings.Get("onDefaultFontSizeChanged").Call("addListener", callback)
}

// OnDefaultFixedFontSizeChanged fired when the default fixed font size setting changes.
func OnDefaultFixedFontSizeChanged(callback func(details js.M)) {
	fontSettings.Get("onDefaultFixedFontSizeChanged").Call("addListener", callback)
}

// OnMinimumFontSizeChanged fired when the minimum font size setting changes.
func OnMinimumFontSizeChanged(callback func(details js.M)) {
	fontSettings.Get("onMinimumFontSizeChanged").Call("addListener", callback)
}
