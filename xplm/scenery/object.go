//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package scenery

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMScenery.h>
#include <stdlib.h>
#include <string.h>

extern void objectLoadedCallback(XPLMObjectRef inObject, void* inRefcon);
*/
import "C"
import (
	"github.com/abieberbach/goplane"
	"github.com/abieberbach/goplane/xplm/graphics"
	"unsafe"
)

type ObjectRef unsafe.Pointer

type DrawInfo struct {
	structSize int32
	x          float32
	y          float32
	z          float32
	pitch      float32
	heading    float32
	roll       float32
}

var objLoadedCallbacks = make(map[*C.char]*objectLoadedReg)

type ObjectLoaded func(objRef ObjectRef, ref interface{})

type objectLoadedReg struct {
	callback ObjectLoaded
	ref      interface{}
}

func LoadObject(path string) ObjectRef {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return ObjectRef(C.XPLMLoadObject(cPath))
}

//export objectLoadedCallback
func objectLoadedCallback(objRef C.XPLMObjectRef, ref unsafe.Pointer) {
	id := (*C.char)(ref)
	regInfo, _ := objLoadedCallbacks[id]
	regInfo.callback(ObjectRef(objRef), regInfo.ref)
	delete(objLoadedCallbacks, id)
	C.free(unsafe.Pointer(id))
}

func LoadObjectAsync(path string, callback ObjectLoaded, ref interface{}) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	id := C.CString(goplane.IdGenerator())
	regInfo := &objectLoadedReg{callback, ref}
	objLoadedCallbacks[id] = regInfo
	C.XPLMLoadObjectAsync(cPath, C.XPLMObjectLoaded_f(unsafe.Pointer(C.objectLoadedCallback)), unsafe.Pointer(id))
}

func DrawObjects(objRef ObjectRef, count int, positions []DrawInfo, lighting, earthRelative bool) {
	if len(positions) == 0 {
		C.XPLMDrawObjects(C.XPLMObjectRef(objRef), C.int(count), nil, C.int(goplane.FromBoolToInt(lighting)), C.int(goplane.FromBoolToInt(earthRelative)))
	} else {
		C.XPLMDrawObjects(C.XPLMObjectRef(objRef), C.int(count), (*C.XPLMDrawInfo_t)(unsafe.Pointer(&positions[0])), C.int(goplane.FromBoolToInt(lighting)), C.int(goplane.FromBoolToInt(earthRelative)))
	}
}

func UnloadObject(objRef ObjectRef) {
	C.XPLMUnloadObject(C.XPLMObjectRef(objRef))
}

func NewDrawInfo(lat, lon, alt, pitch, heading, roll float32) DrawInfo {
	x, y, z := graphics.WorldToLocal(float64(lat), float64(lon), float64(alt))
	return NewDrawInfoLocal(float32(x), float32(y), float32(z), pitch, heading, roll)
}

func NewDrawInfoLocal(x, y, z, pitch, heading, roll float32) DrawInfo {
	result := DrawInfo{0, x, y, z, pitch, heading, roll}
	result.structSize = int32(unsafe.Sizeof(result))
	return result
}
