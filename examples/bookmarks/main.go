package main

import (
	"github.com/Archs/chrome/api/bookmarks"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/Archs/js/utils/property"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	dom.OnDOMContentLoaded(run)
}

func run() {
	ko.EnableSecureBinding()
	println("Get:", ko.Global().Get("asdfasdfafd") == js.Undefined)
	bookmarks.GetTree(func(bs []*bookmarks.TreeNode) {
		bs[0].Title = "Root"
		vm := ko.Mapping().FromJS(bs[0])
		println("vm.toJSON:", ko.Mapping().ToJSON(vm))
		println("vm.title:", vm.Get("title"))
		println("vm.children:", vm.Get("title"))
		ko.ApplyBindings(vm)
	})
}
