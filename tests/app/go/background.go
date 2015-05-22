package main

import (
	"github.com/Archs/chrome/api/app/runtime"
	"github.com/Archs/chrome/api/app/window"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	runtime.OnLaunched(func(data *runtime.LaunchData) {
		window.CreateM("app.html", js.M{
			"bounds": js.M{
				"width":  400,
				"height": 500,
			}}, func(*window.AppWindow) {})
	})
}
