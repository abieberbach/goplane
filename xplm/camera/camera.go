//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package camera

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMCamera.h>
#include <stdlib.h>
#include <string.h>

extern int cameraControlCallback(XPLMCameraPosition_t * outCameraPosition,int inIsLosingControl, void* inRefcon);
*/
import "C"
import "unsafe"

type CameraControlDuration int

type CameraPosition struct {
	X       float32
	Y       float32
	Z       float32
	Pitch   float32
	Heading float32
	Roll    float32
	Zoom    float32
}

type CameraControl func(position *CameraPosition, losingControl bool, ref interface{}) int

type regInfo struct {
	callback CameraControl
	ref      interface{}
}

const (
	ControlCameraUntilViewChanges CameraControlDuration = 1
	ControlCameraForever          CameraControlDuration = 2
)

//export cameraControlCallback
func cameraControlCallback(cameraPosition *C.XPLMCameraPosition_t, isLosingControl C.int, ref unsafe.Pointer) C.int {
	reg := (*regInfo)(ref)
	position := (*CameraPosition)(unsafe.Pointer(cameraPosition))
	return C.int(reg.callback(position, isLosingControl == 1, reg.ref))
}

func ControlCamera(duration CameraControlDuration, cameraControl CameraControl, ref interface{}) {
	reg := &regInfo{cameraControl, ref}
	C.XPLMControlCamera(C.XPLMCameraControlDuration(duration), C.XPLMCameraControl_f(unsafe.Pointer(C.cameraControlCallback)), unsafe.Pointer(reg))
}

func DontControlCamera() {
	C.XPLMDontControlCamera()
}

func IsCameraBeingControlled() (bool, CameraControlDuration) {
	var duration CameraControlDuration
	result := C.XPLMIsCameraBeingControlled((*C.XPLMCameraControlDuration)(unsafe.Pointer(&duration)))
	return result == 1, duration

}

func XPLMReadCameraPosition() CameraPosition {
	var result CameraPosition
	C.XPLMReadCameraPosition((*C.XPLMCameraPosition_t)(unsafe.Pointer(&result)))
	return result
}
