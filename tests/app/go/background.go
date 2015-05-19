package main

import (
	"github.com/Archs/chrome"
)

func main() {
	c := chrome.NewChrome()
	c.App.Runtime.OnLaunched(func(data *chrome.LaunchData) {
		c.App.Window.CreateM("app.html", chrome.Object{
			"bounds": chrome.Object{
				"width":  400,
				"height": 500,
			}}, func(*chrome.AppWindow) {})
	})
}
