//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package display

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMDisplay.h>
#include <stdlib.h>
#include <string.h>

extern int drawCallback(XPLMDrawingPhase inPhase, int inIsBefore, void * inRefcon);
extern int keySnifferCallback(char inChar,XPLMKeyFlags inFlags, char inVirtualKey, void* inRefcon);
*/
import "C"
import (
	"unsafe"
	"github.com/abieberbach/goplane"
	"fmt"
	"github.com/abieberbach/goplane/xplm"
	"github.com/abieberbach/goplane/extra/logging"
	"github.com/go-errors/errors"
)

type DrawingPhase int
type DrawCallback func(phase DrawingPhase, isBefore bool, ref interface{}) int

type KeySnifferCallback func(keyCode xplm.KeyCode, flags xplm.KeyFlags, virtualKeyCode xplm.VirtualKeyCode, ref interface{}) int

const (
	Phase_FirstScene DrawingPhase = 0
	Phase_Terrain DrawingPhase = 5
	Phase_Airports DrawingPhase = 10
	Phase_Vectors DrawingPhase = 15
	Phase_Objects DrawingPhase = 20
	Phase_Airplanes DrawingPhase = 25
	Phase_LastScene DrawingPhase = 30
	Phase_FirstCockpit DrawingPhase = 35
	Phase_Panel DrawingPhase = 40
	Phase_Gauges DrawingPhase = 45
	Phase_Window DrawingPhase = 50
	Phase_LastCockpit DrawingPhase = 55
	Phase_LocalMap3D DrawingPhase = 100
	Phase_LocalMap2D DrawingPhase = 101
	Phase_LocalMapProfile DrawingPhase = 102
)

type regData struct {
	funcPointerAsString string
	callback            DrawCallback
	ref                 interface{}
}

type regSnifferData struct {
	funcPointerAsString string
	callback            KeySnifferCallback
	ref                 interface{}
}


var callbacks = make(map[*C.char]regData)
var keySnifferCallbacks = make(map[*C.char]regSnifferData)


//export drawCallback
func drawCallback(phase, isBefore C.int, ref unsafe.Pointer) C.int {
	defer func() {
		err := recover()
		if err != nil {
			logging.Errorf("Error during drawCallback: %v", err)
			logging.Error(string(errors.Wrap(err, 1).Stack()))
		}
	}()
	id := (*C.char)(ref)
	data, _ := callbacks[id]
	return C.int(data.callback(DrawingPhase(phase), isBefore == 1, data.ref))
}

func RegisterDrawCallback(callback DrawCallback, phase DrawingPhase, wantsBefore bool, ref interface{}) bool {
	id := C.CString(goplane.IdGenerator())
	callbacks[id] = regData{fmt.Sprint(callback), callback, ref}
	return C.XPLMRegisterDrawCallback(C.XPLMDrawCallback_f(unsafe.Pointer(C.drawCallback)), C.XPLMDrawingPhase(phase), C.int(goplane.FromBoolToInt(wantsBefore)), unsafe.Pointer(id)) == 1
}

func UnregisterDrawCallback(callback DrawCallback, phase DrawingPhase, wantsBefore bool, ref interface{}) bool {
	var id *C.char
	for key, info := range callbacks {
		if info.funcPointerAsString == fmt.Sprint(callback)&&info.ref == ref {
			id = key
			defer C.free(unsafe.Pointer(id))
		}
	}
	return C.XPLMUnregisterDrawCallback(C.XPLMDrawCallback_f(unsafe.Pointer(C.drawCallback)), C.XPLMDrawingPhase(phase), C.int(goplane.FromBoolToInt(wantsBefore)), unsafe.Pointer(id)) == 1
}

//export keySnifferCallback
func keySnifferCallback(char C.char, flags C.XPLMKeyFlags, virtualKey C.char, ref unsafe.Pointer) C.int {
	id := (*C.char)(ref)
	data, _ := keySnifferCallbacks[id]
	return C.int(data.callback(xplm.KeyCode(char), xplm.KeyFlags(flags), xplm.VirtualKeyCode(virtualKey), data.ref))

}

func RegisterKeySniffer(callback KeySnifferCallback, beforeWindows bool, ref interface{}) bool {
	id := C.CString(goplane.IdGenerator())
	keySnifferCallbacks[id] = regSnifferData{fmt.Sprint(callback), callback, ref}
	return C.XPLMRegisterKeySniffer(C.XPLMKeySniffer_f(unsafe.Pointer(C.keySnifferCallback)), C.int(goplane.FromBoolToInt(beforeWindows)), unsafe.Pointer(id)) == 1
}

func UnregisterKeySniffer(callback KeySnifferCallback, beforeWindows bool, ref interface{}) bool {
	var id *C.char
	for key, info := range keySnifferCallbacks {
		if info.funcPointerAsString == fmt.Sprint(callback)&&info.ref == ref {
			id = key
			defer C.free(unsafe.Pointer(id))
		}
	}
	return C.XPLMUnregisterKeySniffer(C.XPLMKeySniffer_f(unsafe.Pointer(C.keySnifferCallback)), C.int(goplane.FromBoolToInt(beforeWindows)), unsafe.Pointer(id)) == 1
}