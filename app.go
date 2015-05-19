package chrome

import (
	"github.com/gopherjs/gopherjs/js"
)

type App struct {
	o       *js.Object
	Runtime *AppRuntime
	Window  *AppWindowManager
}

func newApp(c *Chrome) *App {
	app := c.o.Get("app")
	runtime := &AppRuntime{o: app.Get("runtime")}
	window := &AppWindowManager{o: app.Get("window")}
	return &App{
		o:       app,
		Runtime: runtime,
		Window:  window,
	}
}
