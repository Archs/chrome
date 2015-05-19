package main

import (
	"github.com/Archs/chrome/api/bookmarks"
	// "github.com/Archs/js/JSON"
	"github.com/Archs/js/gopherjs-ko"
	// "github.com/gopherjs/gopherjs/js"
)

var (
	out  = ko.NewObservable("")
	bmks = ko.NewObservableArray(nil)
)

func main() {
	ko.EnableSecureBinding()
	// model := js.M{
	// 	"out":       out,
	// 	"bookmarks": bmks,
	// }
	bookmarks.GetTree(func(bs []*bookmarks.TreeNode) {
		bs[0].Title = "Root"
		vm := ko.Mapping().FromJS(bs[0])
		vm.Call("title", "title change after mapping")
		println("vm:", vm)
		println("vm.toJSON:", ko.Mapping().ToJSON(vm))
		ko.ApplyBindings(vm)
	})
}
