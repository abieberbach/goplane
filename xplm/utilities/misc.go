//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package utilities

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMUtilities.h>
#include <stdlib.h>
#include <string.h>

extern void errorCallback(char* inMessage);
*/
import "C"
import "unsafe"

var errCallback ErrorCallback

func ReloadScenery() {
	C.XPLMReloadScenery()
}

func GetSystemPath() string {
	pathBuf := (*C.char)(C.malloc(512))
	defer C.free(unsafe.Pointer(pathBuf))
	C.XPLMGetSystemPath(pathBuf)
	return C.GoString(pathBuf)
}

func GetPrefsPath() string {
	pathBuf := (*C.char)(C.malloc(512))
	defer C.free(unsafe.Pointer(pathBuf))
	C.XPLMGetPrefsPath(pathBuf)
	return C.GoString(pathBuf)
}

func GetDirectorySeparator() string {
	return C.GoString(C.XPLMGetDirectorySeparator())
}

func XPLMExtractFileAndPath(fullPath string) string {
	cFullPath := C.CString(fullPath)
	defer C.free(unsafe.Pointer(cFullPath))
	return C.GoString(C.XPLMExtractFileAndPath(cFullPath))
}

func Initialized() bool {
	return C.XPLMInitialized() == 1
}

func GetVersions() (xplaneVersion, xplmVersion int, appId HostApplicationID) {
	C.XPLMGetVersions(
		(*C.int)(unsafe.Pointer(&xplaneVersion)),
		(*C.int)(unsafe.Pointer(&xplmVersion)),
		(*C.XPLMHostApplicationID)(unsafe.Pointer(&appId)))
	return
}

func GetLanguage() LanguageCode {
	return LanguageCode(C.XPLMGetLanguage())
}

func DebugString(msg string) {
	cMsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cMsg))
	C.XPLMDebugString(cMsg)
}

//export errorCallback
func errorCallback(msg *C.char) {
	errCallback(C.GoString(msg))
}

func SetErrorCallback(callback ErrorCallback) {
	errCallback = callback
	C.XPLMSetErrorCallback(C.XPLMError_f(unsafe.Pointer(C.errorCallback)))
}

func FindSymbol(symbol string) unsafe.Pointer {
	cSymbol := C.CString(symbol)
	defer C.free(unsafe.Pointer(cSymbol))
	return unsafe.Pointer(C.XPLMFindSymbol(cSymbol))
}

func LoadDataFile(fileType DataFileType, filePath string) int {
	cFilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cFilePath))
	return int(C.XPLMLoadDataFile(C.XPLMDataFileType(fileType), cFilePath))
}

func SaveDataFile(fileType DataFileType, filePath string) int {
	cFilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cFilePath))
	return int(C.XPLMSaveDataFile(C.XPLMDataFileType(fileType), cFilePath))
}
