//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package plugins

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMPlugin.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

const (
	NO_PLUGIN_ID  PluginId = -1
	PLUGIN_XPLANE PluginId = 0
)

type PluginId C.int

func GetMyId() PluginId {
	return PluginId(C.XPLMGetMyID())
}

func CountPlugins() int {
	return int(C.XPLMCountPlugins())
}

func GetNthPlugin(index int) PluginId {
	return PluginId(C.XPLMGetNthPlugin(C.int(index)))
}

func FindPluginByPath(path string) PluginId {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return PluginId(C.XPLMFindPluginByPath(cPath))
}

func FindPluginBySignature(signature string) PluginId {
	cSignature := C.CString(signature)
	defer C.free(unsafe.Pointer(cSignature))
	return PluginId(C.XPLMFindPluginBySignature(cSignature))
}

func GetPluginInfo(id PluginId) (name, filePath, signature, description string) {
	nameBuf := (*C.char)(C.malloc(255))
	defer C.free(unsafe.Pointer(nameBuf))
	pathBuf := (*C.char)(C.malloc(255))
	defer C.free(unsafe.Pointer(pathBuf))
	sigBuf := (*C.char)(C.malloc(255))
	defer C.free(unsafe.Pointer(sigBuf))
	descBuf := (*C.char)(C.malloc(255))
	defer C.free(unsafe.Pointer(descBuf))
	C.XPLMGetPluginInfo(C.XPLMPluginID(id), nameBuf, pathBuf, sigBuf, descBuf)
	name = C.GoString(nameBuf)
	filePath = C.GoString(pathBuf)
	signature = C.GoString(sigBuf)
	description = C.GoString(descBuf)
	return
}

func IsPluginEnabled(id PluginId) bool {
	return C.XPLMIsPluginEnabled(C.XPLMPluginID(id)) == 1
}

func EnablePlugin(id PluginId) bool {
	return C.XPLMEnablePlugin(C.XPLMPluginID(id)) == 1
}

func DisablePlugin(id PluginId) {
	C.XPLMDisablePlugin(C.XPLMPluginID(id))
}

func ReloadPlugins(id PluginId) {
	C.XPLMReloadPlugins()
}
