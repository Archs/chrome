package omnibox

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	omnibox = chrome.Get("omnibox")
)

/*
* Types
 */

type SuggestResult struct {
	*js.Object
	Content     string `js:"content"`
	Description string `js:"description"`
}

/*
* Methods
 */

// SetDefaultSuggestion sets the description and styling for the default suggestion.
// The default suggestion is the text that is displayed in the first suggestion row underneath the URL bar.
func SetDefaultSuggestion(suggestion js.M) {
	omnibox.Call("setDefaultSuggestion", suggestion)
}

/*
* Events
 */

// OnInputStarted user has started a keyword input session by typing the extension's keyword.
// This is guaranteed to be sent exactly once per input session, and before any onInputChanged events.
func OnInputStarted(callback func()) {
	omnibox.Get("onInputStarted").Call("addListener", callback)
}

// OnInputChanged user has changed what is typed into the omnibox.
func OnInputChanged(callback func(text string, suggest func(suggestResults []SuggestResult))) {
	omnibox.Get("onInputChanged").Call("addListener", callback)
}

// OnInputEntered user has accepted what is typed into the omnibox.
func OnInputEntered(callback func(text string, disposition string)) {
	omnibox.Get("onInputEntered").Call("addListener", callback)
}

// OnInputCancelled user has ended the keyword input session without accepting the input.
func OnInputCancelled(callback func()) {
	omnibox.Get("onInputCancelled").Call("addListener", callback)
}
