package ttsEngine

import (
	"github.com/Archs/chrome"
	"github.com/Archs/chrome/tts"
	"github.com/gopherjs/gopherjs/js"
)

var (
	ttsEngine = chrome.Get("ttsEngine")
)

/*
* Events
 */

// OnSpeak called when the user makes a call to tts.speak() and one of the voices from this
// extension's manifest is the first to match the options js.M.
func OnSpeak(callback func(utterance string, options js.M, sendItsEvent func(event tts.TtsEvent))) {
	ttsEngine.Get("onSpeak").Call("addListener", callback)
}

// OnStop fired when a call is made to tts.stop and this extension may be in the middle of speaking.
// If an extension receives a call to onStop and speech is already stopped, it should do nothing
// (not raise an error). If speech is in the paused state, this should cancel the paused state.
func OnStop(callback func()) {
	ttsEngine.Get("onStop").Call("addListener", callback)
}

// OnPause is optional: if an engine supports the pause event, it should pause the current utterance
// being spoken, if any, until it receives a resume event or stop event. Note that a stop event should
// also clear the paused state.
func OnPause(callback func()) {
	ttsEngine.Get("onPause").Call("addListener", callback)
}

// OnResume is optional: if an engine supports the pause event, it should also support the resume event,
// to continue speaking the current utterance, if any. Note that a stop event should also clear the paused state.
func OnResume(callback func()) {
	ttsEngine.Get("onResume").Call("addListener", callback)
}
