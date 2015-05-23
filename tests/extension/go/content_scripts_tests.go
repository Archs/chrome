package main

import (
	"github.com/Archs/chrome/api/runtime"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	/*
	* Runtime Events
	 */

	// Listen for OnMessage Event
	runtime.OnMessage(func(message interface{}, sender runtime.MessageSender, sendResponse func(interface{})) {
		// fmt.Println("Runtime.OnMessage received: ", message.(map[string]interface{}))
		println("Runtime.OnMessage received: ", message.(map[string]interface{}), sender)
		resp := js.M{
			"farewell": "goodbye",
		}
		sendResponse(resp)
	})
}
