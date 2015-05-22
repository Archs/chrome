// WARNING: FileBrowserHandler only works on Chrome OS

package fileBrowserHandler

import (
	"github.com/Archs/chrome"
	"github.com/gopherjs/gopherjs/js"
)

var (
	fileBrowserHandler = chrome.Get("fileBrowserHandler")
)

/*
* Types
 */

type FileHandlerExecuteEventDetails struct {
	*js.Object
	Entries []interface{} `js:"entries"`
	Tab_id  int           `js:"tab_id"`
}

/*
* Methods:
 */

/* SelectFile prompts user to select file path under which file should be saved. When the file is selected, file access permission required to use the file (read, write and create) are granted to the caller. The file will not actually get created during the function call, so function caller must ensure its existence before using it. The function has to be invoked with a user gesture. */
func SelectFile(selectionParams js.M, callback func(result js.M)) {
	fileBrowserHandler.Call("selectFile", selectionParams, callback)
}

/*
* Events
 */

// OnExecute fired when file system action is executed from ChromeOS file browser.
func OnExecute(callback func(id string, details FileHandlerExecuteEventDetails)) {
	fileBrowserHandler.Get("onExecute").Call("addListener", callback)
}
