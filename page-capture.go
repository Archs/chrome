package chrome

import "github.com/gopherjs/gopherjs/js"

type PageCapture struct {
	o *js.Object
}

/*
* Methods
 */

// SaveAsMHTML saves the content of the tab with given id as MHTML.
func (p *PageCapture) SaveAsMHTML(details js.M, callback func(mhtmlData interface{})) {
	p.o.Call("saveAsMHTML", details, callback)
}
