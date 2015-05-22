package i18n

import (
	"github.com/Archs/chrome"
)

var (
	i18n = chrome.Get("i18n")
)

/*
* Methods:
 */

// GetAcceptLanguages gets the accept-languages of the browser. This is different from
// the locale used by the browser; to get the locale, use i18n.getUILanguage.
func GetAcceptLanguages(callback func(languages []string)) {
	i18n.Call("getAcceptLanguages", callback)
}

// GetMessage gets the localized string for the specified message. If the message is missing,
// this method returns an empty string (''). If the format of the getMessage() call is wrong —
// for example, messageName is not a string or the substitutions array has more than 9 elements —
// this method returns undefined.
func GetMessage(messageName string, substitutions interface{}) string {
	return i18n.Call("getMessage", messageName, substitutions).String()
}

// GetUILanguage gets the browser UI language of the browser. This is different from
// i18n.getAcceptLanguages which returns the preferred user languages.
func GetUILanguage() string {
	return i18n.Call("getUILanguage").String()
}
