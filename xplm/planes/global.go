package planes
import (
	"unsafe"
	"github.com/abieberbach/goplane/xplm/plugins"
)

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include "../../../../../../SDK/CHeaders/XPLM/XPLMPlanes.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
const USER_AIRCRAFT = 0

func CountAircraft() (totalAircraft int, activeAircraft int, pluginId plugins.PluginId) {
	C.XPLMCountAircraft((*C.int)(unsafe.Pointer(&totalAircraft)),(*C.int)(unsafe.Pointer(&activeAircraft)),(*C.XPLMPluginID)(unsafe.Pointer(&pluginId)))
	return
}

func GetNthAircraftModel(index int) (fileName, path string) {
	nameBuf := (*C.char)(C.malloc(256))
	defer C.free(unsafe.Pointer(nameBuf))
	pathBuf := (*C.char)(C.malloc(512))
	defer C.free(unsafe.Pointer(pathBuf))
	C.XPLMGetNthAircraftModel((C.int)(index),nameBuf,pathBuf)
	fileName=C.GoString(nameBuf)
	path=C.GoString(pathBuf)
	return
}
