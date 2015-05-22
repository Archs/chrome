package tts

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	tts = chrome.Get("tts")
)

/*
* Types
 */

type TtsEvent struct {
	*js.Object
	Type         string `js:"type"`
	CharIndex    int64  `js:"charIndex"`
	ErrorMessage string `js:"errorMessage"`
}

type TtsVoice struct {
	*js.Object
	VoiceName   string   `js:"voiceName"`
	Lang        string   `js:"lang"`
	Gender      string   `js:"gender"`
	Remote      bool     `js:"remote"`
	ExtensionId string   `js:"extensionId"`
	EventTypes  []string `js:"eventTypes"`
}

/*
* Methods
 */

// Speak speaks text using a text-to-speech engine.
func Speak(utterance string, options js.M, callback func()) {
	tts.Call("speak", utterance, options, callback)
}

// Stop stops any current speech and flushes the queue of any pending utterances.
// In addition, if speech was paused, it will now be un-paused for the next call to speak.
func Stop() {
	tts.Call("stop")
}

// Pause pauses speech synthesis, potentially in the middle of an utterance.
// A call to resume or stop will un-pause speech.
func Pause() {
	tts.Call("pause")
}

// Resume if speech was paused, resumes speaking where it left off.
func Resume() {
	tts.Call("resume")
}

// IsSpeaking checks whether the engine is currently speaking. On Mac OS X, the result
// is true whenever the system speech engine is speaking, even if the speech wasn't initiated by Chrome.
func IsSpeaking(callback func(speaking bool)) {
	tts.Call("isSpeaking", callback)
}

// GetVoices gets an array of all available voices.
func GetVoices(callback func(voices []TtsVoice)) {
	tts.Call("getVoices", callback)
}
