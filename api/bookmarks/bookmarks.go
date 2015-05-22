package bookmarks

import (
	"github.com/Archs/chrome"

	"github.com/gopherjs/gopherjs/js"
)

var (
	bookmarks = chrome.Get("bookmarks")
)

/*
* Types
 */

type TreeNode struct {
	*js.Object
	Id                string     `js:"id"`
	ParentId          string     `js:"parentId"`
	Index             int        `js:"index"`
	Url               string     `js:"url"`
	Title             string     `js:"title"`
	DateAdded         int64      `js:"dateAdded"`
	DateGroupModified int64      `js:"dateGroupModified"`
	Unmodifiable      string     `js:"unmodifiable"`
	Children          []TreeNode `js:"children"`
}

/*
* Methods:
 */

// Get retrieves the specified TreeNode(s).
func Get(idList []string, callback func(results []TreeNode)) {
	bookmarks.Call("get", idList, callback)
}

// GetChildren retrieves the children of the specified TreeNode id.
func GetChildren(id string, callback func(results []TreeNode)) {
	bookmarks.Call("getChildren", id, callback)
}

// GetRecent retrieves the recently added bookmarks.
func GetRecent(numberOfItems int, callback func(results []TreeNode)) {
	bookmarks.Call("getRecent", numberOfItems, callback)
}

// GetTree retrieves the entire Bookmarks hierarchy.
func GetTree(callback func(results []TreeNode)) {
	bookmarks.Call("getTree", callback)
}

// GetSubTree retrieves part of the Bookmarks hierarchy, starting at the specified node.
func GetSubTree(id string, callback func(results []TreeNode)) {
	bookmarks.Call("getSubTree", id, callback)
}

// Search searches for TreeNodes matching the given query. Queries specified
// with an js.M produce TreeNodes matching all specified properties.
func Search(query interface{}, callback func(results []TreeNode)) {
	bookmarks.Call("search", query, callback)
}

// Create creates a bookmark or folder under the specified parentId.
// If url is nil or missing, it will be a folder.
func Create(bookmark js.M, callback func(result TreeNode)) {
	bookmarks.Call("create", bookmark, callback)
}

// Move moves the specified TreeNode to the provided location.
func Move(id string, destination js.M, callback func(result TreeNode)) {
	bookmarks.Call("move", id, destination, callback)
}

// Update updates the properties of a bookmark or folder. Specify only the properties that you want
// to change; unspecified properties will be left unchanged. Note: Currently, only 'title' and 'url' are supported.
func Update(id string, changes js.M, callback func(result TreeNode)) {
	bookmarks.Call("update", id, changes, callback)
}

// Remove removes a bookmark or an empty bookmark folder.
func Remove(id string, callback func()) {
	bookmarks.Call("remove", id, callback)
}

// RemoveTree recursively removes a bookmark folder.
func RemoveTree(id string, callback func()) {
	bookmarks.Call("removeTree", id, callback)
}

/*
* Events
 */

// OnCreated fired when a bookmark or folder is created.
func OnCreated(callback func(id string, bookmark TreeNode)) {
	bookmarks.Get("onCreated").Call("addListener", callback)
}

// OnRemoved fired when a bookmark or folder is removed. When a folder is removed recursively,
// a single notification is fired for the folder, and none for its contents.
func OnRemoved(callback func(id string, removeInfo js.M)) {
	bookmarks.Get("onRemoved").Call("addListener", callback)
}

// onChanged fired when a bookmark or folder changes. Note: Currently, only title and url changes trigger this.
func onChanged(callback func(id string, changeInfo js.M)) {
	bookmarks.Get("onChanged").Call("addListener", callback)
}

// OnMoved fired when a bookmark or folder is moved to a different parent folder.
func OnMoved(callback func(id string, moveInfo js.M)) {
	bookmarks.Get("onMoved").Call("addListener", callback)
}

// OnChildrenReordered fired when the children of a folder have changed their order due to
// the order being sorted in the UI. This is not called as a result of a move().
func OnChildrenReordered(callback func(id string, reorderInfo js.M)) {
	bookmarks.Get("onChildrenReordered").Call("addListener", callback)
}

// OnImportBegan fired when a bookmark import session is begun. Expensive observers should ignore
// onCreated updates until onImportEnded is fired. Observers should still handle other notifications immediately.
func OnImportBegan(callback func()) {
	bookmarks.Get("onImportBegan").Call("addListener", callback)
}

// OnImportEnded fired when a bookmark import session is ended.
func OnImportEnded(callback func()) {
	bookmarks.Get("onImportEnded").Call("addListener", callback)
}
