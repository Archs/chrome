package main

import (
	"github.com/Archs/chrome/api/bookmarks"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	ko.EnableSecureBinding()
	println("Get:", ko.Global().Get("asdfasdfafd") == js.Undefined)
	bookmarks.GetTree(func(bs []*bookmarks.TreeNode) {
		bs[0].Title = "Root"
		vm := ko.Mapping().FromJS(bs[0])
		// vm.Set2("TT", "title change after mapping")
		println("vm.toJSON:", ko.Mapping().ToJSON(vm))
		println("vm.title:", vm.Get("title"))
		vm.Set("tt", 123)
		ko.ApplyBindings(vm)
	})
}
