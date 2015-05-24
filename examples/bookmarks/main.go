package main

import (
	"github.com/Archs/chrome/api/bookmarks"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
	"time"
)

var (
	model *MainController
)

func main() {
	dom.OnDOMContentLoaded(run)
}

type MainController struct {
	*js.Object
	Title    *ko.Observable      `js:"Title"`
	Bookmark *BookmarkController `js:"Bookmark"`
	// Bookmark *ko.Observable `js:"Bookmark"`
}

type BookmarkController struct {
	*js.Object
	// Title ko.Observable       `js:"Title"`
	// Show func(bookmarks.TreeNode) `js:"Show"`
	Show func(*ko.MappedViewModel, *dom.Event) `js:"Show"`
	// Show     *js.Object          `js:"Show"`
	Test     func()              `js:"Test"`
	TextArea *ko.Observable      `js:"TextArea"`
	Time     *ko.Observable      `js:"Time"`
	Root     *ko.MappedViewModel `js:"Root"`
	// Root *ko.Observable `js:"Root"`
	// Str      string              `js:"Str"`
	Str string
}

func newBkmkCtrl() *BookmarkController {
	b := BookmarkController{}
	b.Object = js.Global.Get("Object").New()
	b.TextArea = ko.NewObservable("")
	b.Time = ko.NewObservable(time.Now().String())
	b.Str = "nice to see you"
	b.Test = func() {
		b.Time.Set(time.Now().String())
		// b.TextArea.Set(ko.Mapping().ToJSON(b))
		b.TextArea.Set(b.Time.Get().String())
		model.Title.Set(b.Time.Get().String())
		println(b, b.Str)
	}
	// b.Show = func(n bookmarks.TreeNode) {
	b.Show = func(vm *ko.MappedViewModel, e *dom.Event) {
		e.StopPropagation()
		println(e.Type())
		println("Show", vm.Get("id").String(),
			vm.Get("title").String(),
			vm.Get("url").String())
	}
	// b.Show = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	// 	println("Show:", this.Get("url").Invoke().String())
	// 	return js.Undefined
	// })
	return &b
}

func newMainCtrl() *MainController {
	m := MainController{}
	m.Object = js.Global.Get("Object").New()
	m.Title = ko.NewObservable("Bookmark Manager")
	m.Bookmark = newBkmkCtrl()
	return &m
}

func run() {
	ko.EnableSecureBinding()
	// model := newBkmkCtrl()
	// println("model:", model)
	// ko.ApplyBindings(model)
	println("Get:", ko.Global().Get("asdfasdfafd") == js.Undefined)
	model = newMainCtrl()
	// println(model.Bookmark.Show)
	bookmarks.GetTree(func(bs []*bookmarks.TreeNode) {
		bs[0].Title = "Root"
		model.Bookmark.Root = ko.Mapping().FromJS(bs[0])
		// model.Bookmark.Root = ko.NewObservable(ko.Mapping().FromJS(bs[0]))
		ko.ApplyBindings(model)
	})
}
