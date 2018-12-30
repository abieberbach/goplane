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
import (
	"math"
	"unsafe"
)

type NavType int
type NavRef int

const (
	Nav_Unknown      NavType = 0
	Nav_Airport      NavType = 1
	Nav_NDB          NavType = 2
	Nav_VOR          NavType = 4
	Nav_ILS          NavType = 8
	Nav_Localizer    NavType = 16
	Nav_GlideSlope   NavType = 32
	Nav_OuterMarker  NavType = 64
	Nav_MiddleMarker NavType = 128
	Nav_InnerMarker  NavType = 256
	Nav_Fix          NavType = 512
	Nav_DME          NavType = 1024
	Nav_LatLon       NavType = 2048
)

const NAV_NOT_FOUND NavRef = -1

func GetFirstNavAid() NavRef {
	return NavRef(C.XPLMGetFirstNavAid())
}

func GetNextNavAid(navAidRef NavRef) NavRef {
	return NavRef(C.XPLMGetNextNavAid(C.XPLMNavRef(navAidRef)))
}

func FindFirstNavAidOfType(navType NavType) NavRef {
	return NavRef(C.XPLMFindFirstNavAidOfType(C.XPLMNavType(navType)))
}

func FindLastNavAidOfType(navType NavType) NavRef {
	return NavRef(C.XPLMFindLastNavAidOfType(C.XPLMNavType(navType)))
}

func FindNavAid(nameFragment, idFrament string, lat, lon float32, freq int, navType NavType) NavRef {
	var cNameFragment (*C.char) = nil
	var cIdFrament (*C.char) = nil
	if len(nameFragment) > 0 {
		cNameFragment = C.CString(nameFragment)
		defer C.free(unsafe.Pointer(cNameFragment))
	}
	if len(idFrament) > 0 {
		cIdFrament = C.CString(idFrament)
		defer C.free(unsafe.Pointer(cIdFrament))
	}
	var (
		cLat, cLon *C.float
		cFreq      *C.int
	)
	if lat != math.MaxFloat32 {
		cLat = (*C.float)(unsafe.Pointer(&lat))
	}
	if lon != math.MaxFloat32 {
		cLat = (*C.float)(unsafe.Pointer(&lon))
	}
	if freq != math.MaxInt32 {
		cFreq = (*C.int)(unsafe.Pointer(&freq))
	}
	return NavRef(C.XPLMFindNavAid(cNameFragment, cIdFrament, cLat, cLon, cFreq, C.XPLMNavType(navType)))
}

func GetNavAidInfo(navAidRef NavRef) (navType NavType, lat, lon, height float32, frequency int, heading float32, id, name, reg string) {
	cId := (*C.char)(C.malloc(32))
	defer C.free(unsafe.Pointer(cId))
	cName := (*C.char)(C.malloc(256))
	defer C.free(unsafe.Pointer(cName))
	cReg := (*C.char)(C.malloc(256))
	defer C.free(unsafe.Pointer(cReg))
	C.XPLMGetNavAidInfo(
		C.XPLMNavRef(navAidRef),
		(*C.XPLMNavType)(unsafe.Pointer(&navType)),
		(*C.float)(unsafe.Pointer(&lat)),
		(*C.float)(unsafe.Pointer(&lon)),
		(*C.float)(unsafe.Pointer(&height)),
		(*C.int)(unsafe.Pointer(&frequency)),
		(*C.float)(unsafe.Pointer(&heading)),
		cId, cName, cReg)
	id = C.GoString(cId)
	name = C.GoString(cName)
	reg = C.GoString(cReg)
	return
}
