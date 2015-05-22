package commands

import (
	"github.com/Archs/chrome"
)

var (
	commands = chrome.Get("commands")
)

/*
* Types
 */

type Command map[string]string

/*
* Methods:
 */

// GetAll returns all the registered extension commands for this extension and their shortcut (if active).
func GetAll(callback func(commands []Command)) {
	commands.Call("getAll", callback)
}

/*
* Events
 */

// OnCommand fired when a registered command is activated using a keyboard shortcut.
func OnCommand(callback func(command string)) {
	commands.Get("onCommand").Call("addListener", callback)
}
