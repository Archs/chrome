// WARNING: This API only works on Chrome OS

package fileSystemProvider

import (
	"github.com/Archs/chrome"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

var (
	fileSystemProvider = chrome.Get("fileSystemProvider")
)

/*
* Types
 */

type EntryMetadata struct {
	*js.Object
	IsDirectory      bool      `js:"isDirectory"`
	Name             string    `js:"name"`
	Size             int64     `js:"size"`
	ModificationTime time.Time `js:"modificationTime"`
	MimeType         string    `js:"mimeType"`
	Thumbnail        string    `js:"thumbnail"`
}

type Watcher struct {
	*js.Object
	EntryPath string `js:"entryPath"`
	Recursive bool   `js:"recursive"`
	LastTag   string `js:"lastTag"`
}

type FileSystemInfo struct {
	*js.Object
	FileSystemId    string `js:"fileSystemId"`
	DisplayName     string `js:"displayName"`
	Writable        bool   `js:"writable"`
	OpenedFileLimit int64  `js:"openedFileLimit"`
	OpenedFiles     []js.M `js:"openedFiles"`
}

type AddWatcherRequestedOptions struct {
	*js.Object
	FilesystemId string `js:"fileSystemId"`
	RequestId    int    `js:"requestId"`
	EntryPath    string `js:"entryPath"`
	Recursive    bool   `js:"recursive"`
}

type RemoveWatcherRequestedOptions AddWatcherRequestedOptions

type NotifyOptions struct {
	*js.Object
	FilesystemId string `js:"fileSystemId"`
	ObservedPath string `js:"observedPath"`
	Recursive    bool   `js:"recursive"`
	ChangeType   string `js:"changeType"`
	Changes      []js.M `js:"changes"`
	Tag          string `js:"tag"`
}

/*
* Methods:
 */

/* Mount mounts a file system with the given fileSystemId and displayName. displayName will be shown in the left panel of Files.app. displayName can contain any characters including '/', but cannot be an empty string. displayName must be descriptive but doesn't have to be unique. The fileSystemId must not be an empty string. */
func Mount(options js.M, callback func()) {
	fileSystemProvider.Call("mount", options, callback)
}

// Unmount unmounts a file system with the given fileSystemId. It must be called after onUnmountRequested is invoked.
// Also, the providing extension can decide to perform unmounting if not requested (eg. in case of lost connection, or a file error).
func Unmount(options js.M, callback func()) {
	fileSystemProvider.Call("unmount", options, callback)
}

// GetAll returns all file systems mounted by the extension.
func GetAll(callback func(fileSystems []FileSystemInfo)) {
	fileSystemProvider.Call("getAll", callback)
}

// Get returns information about a file system with the passed fileSystemId.
func Get(fileSystemId string, callback func(fileSystemInfo FileSystemInfo)) {
	fileSystemProvider.Call("get", fileSystemId, callback)
}

/*
* Events
 */

// OnUnmountRequested raised when unmounting for the file system with the fileSystemId identifier is requested.
// In the response, the unmount API method must be called together with successCallback .
// If unmounting is not possible (eg. due to a pending operation), then errorCallback must be called.
func OnUnmountRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onUnmountRequested").Call("addListener", callback)
}

// OnGetMetadataRequested raised when metadata of a file or a directory at entryPath is requested.
// The metadata must be returned with the successCallback call. In case of an error, errorCallback must be called.
func OnGetMetadataRequested(callback func(options js.M, successCallback func(metadata EntryMetadata), errorCallback func(err string))) {
	fileSystemProvider.Get("onGetMetadataRequested").Call("addListener", callback)
}

// OnReadDirectoryRequested raised when contents of a directory at directoryPath are requested.
// The results must be returned in chunks by calling the successCallback several times. In case of an error, errorCallback must be called.
func OnReadDirectoryRequested(callback func(options js.M, successCallback func(entries []EntryMetadata, hasMore bool), errorCallback func(err string))) {
	fileSystemProvider.Get("onReadDirectoryRequested").Call("addListener", callback)
}

// OnOpenFileRequested raised when opening a file at filePath is requested. If the file does not exist,
// then the operation must fail. Maximum number of files opened at once can be specified with MountOptions.
func OnOpenFileRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onOpenFileRequested").Call("addListener", callback)
}

// OnCloseFileRequested raised when opening a file previously opened with openRequestId is requested to be closed.
func OnCloseFileRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onCloseFileRequested").Call("addListener", callback)
}

// OnReadFileRequested raised when reading contents of a file opened previously with openRequestId is requested.
// The results must be returned in chunks by calling successCallback several times. In case of an error, errorCallback must be called.
func OnReadFileRequested(callback func(options js.M, successCallback func(data interface{}, hasMore bool), errorCallback func(err string))) {
	fileSystemProvider.Get("onReadFileRequested").Call("addListener", callback)
}

// OnCreateDirectoryRequested raised when creating a directory is requested. The operation must fail with the EXISTS error
// if the target directory already exists. If recursive is true, then all of the missing directories on the directory path must be created.
func OnCreateDirectoryRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onCreateDirectoryRequested").Call("addListener", callback)
}

// OnDeleteEntryRequested raised when deleting an entry is requested. If recursive is true, and the entry is a directory,
// then all of the entries inside must be recursively deleted as well.
func OnDeleteEntryRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onDeleteEntryRequested").Call("addListener", callback)
}

// OnCreateFileReqested raised when creating a file is requested. If the file already exists, then errorCallback must
// be called with the EXISTS error code.
func OnCreateFileReqested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onCreateFileReqested").Call("addListener", callback)
}

// OnCopyEntryRequested raised when copying an entry (recursively if a directory) is requested. If an error occurs,
// then errorCallback must be called.
func OnCopyEntryRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onCopyEntryRequested").Call("addListener", callback)
}

// OnMoveEntryRequested raised when moving an entry (recursively if a directory) is requested. If an error occurs,
// then errorCallback must be called.
func OnMoveEntryRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onMoveEntryRequested").Call("addListener", callback)
}

// OnTruncateRequested raised when truncating a file to a desired length is requested. If an error occurs,
// then errorCallback must be called.
func OnTruncateRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onTruncateRequested").Call("addListener", callback)
}

// OnWriteFileRequested raised when writing contents to a file opened previously with openRequestId is requested.
func OnWriteFileRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onWriteFileRequested").Call("addListener", callback)
}

/* OnAbortRequested raised when aborting an operation with operationRequestId is requested. The operation executed with operationRequestId must be immediately stopped and successCallback of this abort request executed. If aborting fails, then errorCallback must be called. Note, that callbacks of the aborted operation must not be called, as they will be ignored. Despite calling errorCallback, the request may be forcibly aborted. */
func OnAbortRequested(callback func(options js.M, successCallback func(), errorCallback func(err string))) {
	fileSystemProvider.Get("onAbortRequested").Call("addListener", callback)
}
