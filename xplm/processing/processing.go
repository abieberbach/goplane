//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package processing

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMProcessing.h>
#include <stdlib.h>
#include <string.h>

extern float flightLoop(float elapsedSinceLastCall,float elapsedTimeSinceLastFlightLoop, int counter, void* ref);

*/
import "C"
import (
	"fmt"
	"github.com/abieberbach/goplane"
	"github.com/abieberbach/goplane/extra/logging"
	"github.com/go-errors/errors"
	"unsafe"
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
	flightLoopData      *CreateFlightLoopData
}

const (
	FlightLoop_Phase_BeforeFlightModel FlightLoopPhaseType = 0
	FlightLoop_Phase_AfterFlightModel  FlightLoopPhaseType = 1
)

var callbacks = make(map[*C.char]regData)

func GetElapsedTime() float32 {
	return float32(C.XPLMGetElapsedTime())
}

func GetCycleNumber() int {
	return int(C.XPLMGetCycleNumber())
}

func RegisterFlightLoopCallback(flightLoop FlightLoopFunc, interval float32, ref interface{}) {
	data := regData{fmt.Sprint(flightLoop), &CreateFlightLoopData{FlightLoop_Phase_BeforeFlightModel, flightLoop, ref}}
	id := C.CString(goplane.IdGenerator())
	callbacks[id] = data
	C.XPLMRegisterFlightLoopCallback(C.XPLMFlightLoop_f(unsafe.Pointer(C.flightLoop)), C.float(interval), unsafe.Pointer(id))
}

func UnregisterFlightLoopCallback(flightLoop FlightLoopFunc, ref interface{}) {
	for key, data := range callbacks {
		if data.funcPointerAsString == fmt.Sprint(flightLoop) && data.flightLoopData.Ref == ref {
			C.XPLMUnregisterFlightLoopCallback(C.XPLMFlightLoop_f(unsafe.Pointer(C.flightLoop)), unsafe.Pointer(data.flightLoopData))
			delete(callbacks, key)
		}
	}

}

func SetFlightLoopCallbackInterval(flightLoop FlightLoopFunc, interval float32, relativeToNow bool, ref interface{}) {
	cRelativeToNow := 0
	if relativeToNow {
		cRelativeToNow = 1
	}
	for key, current := range callbacks {
		if current.funcPointerAsString == fmt.Sprint(flightLoop) {
			current.flightLoopData.Ref = ref
			C.XPLMSetFlightLoopCallbackInterval(C.XPLMFlightLoop_f(C.flightLoop), C.float(interval), C.int(cRelativeToNow), unsafe.Pointer(key))
			return
		}
	}
}

//export flightLoop
func flightLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop C.float, counter C.int, ref unsafe.Pointer) C.float {
	defer func() {
		err := recover()
		if err != nil {
			logging.Errorf("Error during flightLoop: %v", err)
			logging.Error(errors.Wrap(err, 1).ErrorStack())
		}
	}()

	id := (*C.char)(ref)
	regData, _ := callbacks[id]
	return C.float(regData.flightLoopData.Callback(float32(elapsedSinceLastCall), float32(elapsedTimeSinceLastFlightLoop), int(counter), regData.flightLoopData.Ref))
}

func CreateFlightLoop(data *CreateFlightLoopData) FlightLoopId {
	id := C.CString(goplane.IdGenerator())
	cRegData := C.XPLMCreateFlightLoop_t{}
	cRegData.structSize = C.int(unsafe.Sizeof(cRegData))
	cRegData.phase = C.XPLMFlightLoopPhaseType(data.Phase)
	cRegData.callbackFunc = C.XPLMFlightLoop_f(unsafe.Pointer(C.flightLoop))
	cRegData.refcon = unsafe.Pointer(id)
	callbacks[id] = regData{fmt.Sprint(data.Callback), data}
	flightLoopId := FlightLoopId(C.XPLMCreateFlightLoop(&cRegData))
	return flightLoopId
}

func DestroyFlightLoop(flightLoopId FlightLoopId) {
	C.XPLMDestroyFlightLoop(C.XPLMFlightLoopID(flightLoopId))
}

func ScheduleFlightLoop(flightLoopId FlightLoopId, interval float32, relativeToNow bool) {
	cRelativeToNow := 0
	if relativeToNow {
		cRelativeToNow = 1
	}
	C.XPLMScheduleFlightLoop(C.XPLMFlightLoopID(flightLoopId), C.float(interval), C.int(cRelativeToNow))
}
