package main

import (
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
	"time"
)

type model struct {
	*js.Object
	Normal    *ko.ObservableArray `js:"normal"`
	Throttled *ko.ObservableArray `js:"throttled"`
	Debounced *ko.ObservableArray `js:"debounced"`
	Start     func()              `js:"go"`
}

func newModel() *model {
	m := new(model)
	m.Object = js.Global.Get("Object").New()
	m.Normal = ko.NewObservableArray()
	m.Throttled = ko.NewObservableArray()
	m.Debounced = ko.NewObservableArray()
	m.Throttled.RateLimit(500)
	m.Debounced.RateLimit(500, true)
	m.Start = func() {
		go func() {
			v := []interface{}{}
			m.Normal.Set(v)
			m.Throttled.Set(v)
			m.Debounced.Set(v)
			println(time.Now().String())
			time.Sleep(500 * time.Millisecond)
			for i := 0; i < 100; i++ {
				m.Normal.Push(i)
				m.Throttled.Push(i)
				m.Debounced.Push(i)
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}
	return m
}

func main() {
	ko.EnableSecureBinding()
	m := newModel()
	println("m:", m)
	ko.ApplyBindings(m)
}
