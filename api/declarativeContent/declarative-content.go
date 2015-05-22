package declarativeContent

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	declarativeContent = chrome.Get("declarativeContent")
	OnPageChanged      = newOnPageChangedEvent(declarativeContent.Get("onPageChanged"))
)

type OnPageChangedEvent struct {
	o *js.Object
}

func newOnPageChangedEvent(o *js.Object) *OnPageChangedEvent {
	d := new(OnPageChangedEvent)
	d.o = o
	return d
}

/*
* Types
 */

type PageStateMatcher struct {
	*js.Object
	PageUrl js.M     `js:"pageUrl"`
	Css     []string `js:"css"`
}

type RequestContentScript struct {
	*js.Object
	Css             []string `js:"css"`
	Js              []string `js:"js"`
	AllFrames       bool     `js:"allFrames"`
	MatchAboutBlank bool     `js:"matchAboutBlank"`
}

/*
* Events
 */

// AddRules takes an array of rule instances as its first parameter and a callback function that is called on completion
func (e *OnPageChangedEvent) AddRules(rules []map[string]interface{}, callback func(details map[string]interface{})) {
	e.o.Call("addRules", rules, callback)
}

// RemoveRules accepts an optional array of rule identifiers as its first parameter and a callback function as its second parameter.
func (e *OnPageChangedEvent) RemoveRules(ruleIdentifiers []string, callback func()) {
	e.o.Call("removeRules", ruleIdentifiers, callback)
}

// GetRules accepts an optional array of rule identifiers with the same semantics as removeRules and a callback function.
func (e *OnPageChangedEvent) GetRules(ruleIdentifiers []string, callback func(details map[string]interface{})) {
	e.o.Call("getRules", ruleIdentifiers, callback)
}
