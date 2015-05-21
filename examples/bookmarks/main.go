package main

import (
	"github.com/Archs/chrome/api/bookmarks"
	"github.com/Archs/chrome/api/tabs"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
	"time"
)

var (
	model *MainController
)

func main() {
	ko.EnableSecureBinding()
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
	Toggle func(*ko.ViewModel, *dom.Event) `js:"Toggle"`
	// Show     *js.Object          `js:"Show"`
	Test     func()                         `js:"Test"`
	TextArea *ko.Observable                 `js:"TextArea"`
	Time     *ko.Observable                 `js:"Time"`
	Root     *ko.ViewModel                  `js:"Root"`
	Goto     func(node *ko.ViewModel)       `js:"Goto"`
	Edit     func(node *bookmarks.TreeNode) `js:"Edit"`
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
	b.Toggle = func(vm *ko.ViewModel, e *dom.Event) {
		el := e.Target()
		children := el.QuerySelector("section")
		println("children:", children)
		if children == nil {
			return
		}
		println("children.ClassName:", children.ClassName)
		cls := children.ClassName
		if cls == "" {
			children.ClassName = "children"
		} else {
			children.ClassName = ""
		}
	}
	b.Goto = func(vm *ko.ViewModel) {
		url := vm.Get("url").String()
		tabs.Create(js.M{"url": url}, func(t tabs.Tab) {
			println(t)
		})
	}
	b.Edit = func(node *bookmarks.TreeNode) {
		println("Goto:", node)
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
	m.Title = ko.NewObservable("Bookmark Viewer")
	m.Bookmark = newBkmkCtrl()
	return &m
}

func run() {
	model = newMainCtrl()
	bookmarks.GetTree(func(bs []*bookmarks.TreeNode) {
		bs[0].Title = "Root"
		model.Bookmark.Root = ko.Mapping().FromJS(bs[0])
		println("bs:", bs)
		ko.ApplyBindings(model)
	})
}
