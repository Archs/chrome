package main

import (
	"github.com/Archs/chrome/api/bookmarks"
	"github.com/Archs/gopherjs-ko"
	"github.com/Archs/js/JSON"
	"github.com/gopherjs/gopherjs/js"
)

var (
	out  = ko.NewObservable("")
	bmks = ko.NewObservableArray(nil)
)

func main() {
	ko.EnableSecureBinding()
	model := js.M{
		"out":       out,
		"bookmarks": bmks,
	}
	ko.ApplyBindings(model)
	bookmarks.GetTree(func(bs []bookmarks.TreeNode) {
		bmks.Set(bs)
		str := JSON.Stringify(bs)
		println(str)
		out.Set(str)
	})
}
