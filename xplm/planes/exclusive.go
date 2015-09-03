//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package planes

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMPlanes.h>
#include <stdlib.h>
#include <string.h>

static char**makeCharArray(int size) {
        return calloc(sizeof(char*), size);
}

static void setArrayString(char **a, char *s, int n) {
        a[n] = s;
}

static void freeCharArray(char **a, int size) {
        int i;
        for (i = 0; i < size; i++)
                free(a[i]);
        free(a);
}

extern void planesAvailableCallback(void* ref);
*/
import "C"
import (
	"unsafe"
	"github.com/abieberbach/goplane"
)


type  PlaneDrawState struct {
	structSize      int
	GearPosition    float32
	FlapRatio       float32
	SpoilerRatio    float32
	SpeedBrakeRatio float32
	SlatRatio       float32
	WingSweep       float32
	Thrust          float32
	YokePitch       float32
	YokeHeading     float32
	YokeRoll        float32
}

type PlanesAvailable func(ref interface{})

type regInfo struct {
	callback PlanesAvailable
	ref interface{}
}

//export planesAvailableCallback
func planesAvailableCallback(ref unsafe.Pointer) {
	reg:=(*regInfo)(ref)
	reg.callback(reg.ref)
}

func AcquirePlanes(planes []string, callback PlanesAvailable, ref interface{}) int {
	cAircrafts := C.makeCharArray(C.int(len(planes)))
	defer C.freeCharArray(cAircrafts, C.int(len(planes)))
	for i, s := range planes {
		C.setArrayString(cAircrafts, C.CString(s), C.int(i))
	}
	reg:=&regInfo{callback,ref}
	return int(C.XPLMAcquirePlanes(cAircrafts,C.XPLMPlanesAvailable_f(unsafe.Pointer(C.planesAvailableCallback)),unsafe.Pointer(reg)))
}

func ReleasePlanes() {
	C.XPLMReleasePlanes()
}

func SetActiveAircraftCount(count int) {
	C.XPLMSetActiveAircraftCount(C.int(count))
}

func SetAircraftModel(index int, path string) {
	cPath:=C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	C.XPLMSetAircraftModel(C.int(index),cPath)
}

func DisableAIForPlane(index int) {
	C.XPLMDisableAIForPlane(C.int(index))
}

func DrawAircraft(index int, x,y,z,pitch,roll,yaw float32, fullDraw bool,drawState  PlaneDrawState) {
	drawState.structSize=int(unsafe.Sizeof(drawState))
	C.XPLMDrawAircraft(
		C.int(index),
		C.float(x),
		C.float(y),
		C.float(z),
		C.float(pitch),
		C.float(roll),
		C.float(yaw),
		C.int(goplane.FromBoolToInt(fullDraw)),
		(*C.XPLMPlaneDrawState_t)(unsafe.Pointer(&drawState)))
}

func ReinitUsersPlane() {
	C.XPLMReinitUsersPlane()
}