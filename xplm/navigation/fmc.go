//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package navigation

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMNavigation.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

func CountFMSEntries() int {
	return int(C.XPLMCountFMSEntries())
}

func GetDisplayedFMSEntry() int {
	return int(C.XPLMGetDisplayedFMSEntry())
}

func GetDestinationFMSEntry() int {
	return int(C.XPLMGetDestinationFMSEntry())
}

func SetDisplayedFMSEntry(index int) {
	C.XPLMSetDisplayedFMSEntry(C.int(index))
}

func SetDestinationFMSEntry(index int) {
	C.XPLMSetDestinationFMSEntry(C.int(index))
}

func GetFMSEntryInfo(index int) (navType NavType, id string, navRef NavRef, alt int, lat, lon float32) {
	cId := (*C.char)(C.malloc(32))
	defer C.free(unsafe.Pointer(cId))
	C.XPLMGetFMSEntryInfo(
		C.int(index),
		(*C.XPLMNavType)(unsafe.Pointer(&navType)),
		cId,
		(*C.XPLMNavRef)(unsafe.Pointer(&navRef)),
		(*C.int)(unsafe.Pointer(&alt)),
		(*C.float)(unsafe.Pointer(&lat)),
		(*C.float)(unsafe.Pointer(&lon)))

	id = C.GoString(cId)
	return
}

func SetFMSEntryInfo(index int, navRef NavRef, altitude int) {
	C.XPLMSetFMSEntryInfo(C.int(index), C.XPLMNavRef(navRef), C.int(altitude))
}

func SetFMSEntryLatLon(index int, lat, lon float32, altitude int) {
	C.XPLMSetFMSEntryLatLon(C.int(index), C.float(lat), C.float(lon), C.int(altitude))
}

func ClearFMSEntry(index int) {
	C.XPLMClearFMSEntry(C.int(index))
}
