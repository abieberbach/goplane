//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package processing

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMProcessing.h>
#include <stdlib.h>
#include <string.h>

extern float flightLoop(float elapsedSinceLastCall,float elapsedTimeSinceLastFlightLoop, int counter, void* ref);

*/
import "C"
import (
	"unsafe"
	"fmt"
	"github.com/abieberbach/goplane"
)

type FlightLoopPhaseType int
type FlightLoopId unsafe.Pointer

type FlightLoopFunc func(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop float32, counter int, ref interface{}) float32

type CreateFlightLoopData struct {
	Phase    FlightLoopPhaseType
	Callback FlightLoopFunc
	Ref      interface{}
}

type regData struct {
	funcPointerAsString string
	flightLoopData *CreateFlightLoopData
}

const (
	FlightLoop_Phase_BeforeFlightModel  FlightLoopPhaseType = 0
	FlightLoop_Phase_AfterFlightModel   FlightLoopPhaseType = 1
)

var callbacks=make(map[string]regData)

func GetElapsedTime() float32 {
	return float32(C.XPLMGetElapsedTime())
}

func GetCycleNumber() int {
	return int(C.XPLMGetCycleNumber())
}

func RegisterFlightLoopCallback(flightLoop FlightLoopFunc, interval float32, ref interface{}) {
	data:=regData{fmt.Sprint(flightLoop),&CreateFlightLoopData{FlightLoop_Phase_BeforeFlightModel,flightLoop,ref}}
	callbacks[goplane.IdGenerator()]=data
	C.XPLMRegisterFlightLoopCallback(C.XPLMFlightLoop_f(unsafe.Pointer(C.flightLoop)), C.float(interval), unsafe.Pointer(data.flightLoopData))
}

func UnregisterFlightLoopCallback(flightLoop FlightLoopFunc, ref interface{}) {
	for key,data:=range callbacks {
		if data.funcPointerAsString==fmt.Sprint(flightLoop)&&data.flightLoopData.Ref==ref {
			C.XPLMUnregisterFlightLoopCallback(C.XPLMFlightLoop_f(unsafe.Pointer(C.flightLoop)), unsafe.Pointer(data.flightLoopData))
			delete(callbacks,key)
		}
	}

}

func SetFlightLoopCallbackInterval(flightLoop unsafe.Pointer, interval float32, relativeToNow bool, ref unsafe.Pointer) {
	cRelativeToNow := 0
	if relativeToNow {
		cRelativeToNow=1
	}
	C.XPLMSetFlightLoopCallbackInterval(C.XPLMFlightLoop_f(flightLoop), C.float(interval), C.int(cRelativeToNow), ref)
}

//export flightLoop
func flightLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop C.float, counter C.int, ref unsafe.Pointer) C.float {
	regData:=(*CreateFlightLoopData)(ref)
	return C.float(regData.Callback(float32(elapsedSinceLastCall),float32(elapsedTimeSinceLastFlightLoop),int(counter),regData.Ref))
}

func CreateFlightLoop(regData *CreateFlightLoopData) FlightLoopId {
	cRegData := C.XPLMCreateFlightLoop_t{}
	cRegData.structSize=C.int(unsafe.Sizeof(cRegData))
	cRegData.phase=C.XPLMFlightLoopPhaseType(regData.Phase)
	cRegData.callbackFunc=C.XPLMFlightLoop_f(unsafe.Pointer(C.flightLoop))
	cRegData.refcon=unsafe.Pointer(regData)
	flightLoopId:=FlightLoopId(C.XPLMCreateFlightLoop(&cRegData))
	return flightLoopId
}

func DestroyFlightLoop(flightLoopId FlightLoopId) {
	C.XPLMDestroyFlightLoop(C.XPLMFlightLoopID(flightLoopId));
}

func ScheduleFlightLoop(flightLoopId FlightLoopId, interval float32, relativeToNow bool) {
	cRelativeToNow := 0
	if relativeToNow {
		cRelativeToNow=1
	}
	C.XPLMScheduleFlightLoop(C.XPLMFlightLoopID(flightLoopId), C.float(interval), C.int(cRelativeToNow));
}
