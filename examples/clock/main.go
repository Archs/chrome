package main

import (
	"github.com/Archs/js/gopherjs-ko"
	_ "github.com/Archs/js/gopherjs-ko/components/clock"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	ko.EnableSecureBinding()
	ko.ApplyBindings(js.M{
		"Title": "An Analog Clock",
	})
}
