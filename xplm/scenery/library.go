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

extern void lookupObjectsEnum(char* inFilePath,void* inRef);

*/
import "C"
import (
	"unsafe"
)

type LibraryEnumerator func(path string, ref interface{})

type lookupRegInfo struct {
	enumerator LibraryEnumerator
	ref        interface{}
}

//export lookupObjectsEnum
func lookupObjectsEnum(cPath *C.char, ref unsafe.Pointer) {
	regInfo := (*lookupRegInfo)(ref)
	regInfo.enumerator(C.GoString(cPath), regInfo.ref)
}

func LookupObjects(path string, lat, lon float32, enumerator LibraryEnumerator, ref interface{}) int {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	regInfo := &lookupRegInfo{enumerator, ref}
	return int(C.XPLMLookupObjects(cPath, C.float(lat), C.float(lon), C.XPLMLibraryEnumerator_f(unsafe.Pointer(C.lookupObjectsEnum)), unsafe.Pointer(regInfo)))
}
