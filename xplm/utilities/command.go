//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package utilities

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMUtilities.h>
#include <stdlib.h>
#include <string.h>

extern int commandCallback(XPLMCommandRef inCommand, XPLMCommandPhase inPhase,void* inRefcon);
*/
import "C"
import (
	"unsafe"
	"github.com/abieberbach/goplane"
	"fmt"
)

type CommandPhase int
type CommandRef unsafe.Pointer

type CommandCallback func(command CommandRef, phase CommandPhase, ref interface{}) int

const (
	Phase_CommandBegin                       CommandPhase = 0
	Phase_CommandContinue                    CommandPhase = 1
	Phase_CommandEnd                         CommandPhase = 2
)

type regInfo struct {
	callbackAsString string
	command CommandRef
	callback CommandCallback
	ref interface{}
}

var callbacks = make(map[*C.char]*regInfo)

func FindCommand(name string) CommandRef {
	cName:=C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return CommandRef(C.XPLMFindCommand(cName))
}

func CommandBegin(command CommandRef) {
	C.XPLMCommandBegin(C.XPLMCommandRef(command))
}

func CommandEnd(command CommandRef) {
	C.XPLMCommandEnd(C.XPLMCommandRef(command))
}

func CommandOnce(command CommandRef) {
	C.XPLMCommandOnce(C.XPLMCommandRef(command))
}

func CreateCommand(name, description string) CommandRef {
	cName:=C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cDescription:=C.CString(description)
	defer C.free(unsafe.Pointer(cDescription))
	return CommandRef(C.XPLMCreateCommand(cName,cDescription))
}

//export commandCallback
func commandCallback(command C.XPLMCommandRef, phase C.XPLMCommandPhase,refcon unsafe.Pointer) C.int {
	id:=(*C.char)(refcon)
	reg:=callbacks[id]
	defer func() {
		err:=recover()
		DebugString("Fatal: "+fmt.Sprint(err))
	}()
	return C.int(reg.callback(CommandRef(command),CommandPhase(phase),reg.ref))
}

func RegisterCommandHandler(command CommandRef, callback CommandCallback, before bool, ref interface{}) {
	cId:=C.CString(goplane.IdGenerator())
	callbacks[cId]=&regInfo{fmt.Sprint(callback),command,callback,ref}
	C.XPLMRegisterCommandHandler(C.XPLMCommandRef(command),C.XPLMCommandCallback_f(unsafe.Pointer(C.commandCallback)),C.int(goplane.FromBoolToInt(before)),unsafe.Pointer(cId))
}

func UnregisterCommandHandler(command CommandRef, callback CommandCallback, before bool, ref interface{}) {
	callbackAsString:=fmt.Sprint(callback)
	for cId,regInfo:=range callbacks {
		if regInfo.command==command && regInfo.callbackAsString==callbackAsString && regInfo.ref==ref {
			C.XPLMUnregisterCommandHandler(C.XPLMCommandRef(command),C.XPLMCommandCallback_f(unsafe.Pointer(C.commandCallback)),C.int(goplane.FromBoolToInt(before)),unsafe.Pointer(cId))
			defer C.free(unsafe.Pointer(cId))
		}
	}
}