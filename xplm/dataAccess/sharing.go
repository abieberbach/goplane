package dataAccess

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMDataAccess.h>
#include <stdlib.h>

extern void valueChanged(void * inRef);

*/
import "C"
import (
	"unsafe"
)


var callbacks = make(map[unsafe.Pointer]shareRegInfo)

type DataChangedFunc func(ref interface{})

type shareRegInfo struct {
	name string
	callback DataChangedFunc
	ref      interface{}
}

//export valueChanged
func valueChanged(ref unsafe.Pointer) {
	regInfo, _ := callbacks[ref]
	regInfo.callback(regInfo.ref)
}

func ShareData(name string, dataType DataRefType, callback DataChangedFunc, ref interface{}) {
	cName := C.CString(name)
	callbacks[unsafe.Pointer(cName)]=shareRegInfo{name,callback,ref}
	C.XPLMShareData(cName,C.XPLMDataTypeID(dataType),C.XPLMDataChanged_f(unsafe.Pointer(C.valueChanged)),unsafe.Pointer(cName))
}

func UnshareData(name string, dataType DataRefType, callback DataChangedFunc, ref interface{}) {
	for cName,shareRegInfo:=range callbacks {
		if shareRegInfo.name==name {
			C.XPLMUnshareData((*C.char)(cName),C.XPLMDataTypeID(dataType),C.XPLMDataChanged_f(unsafe.Pointer(C.valueChanged)),cName)
			delete(callbacks,cName)
			break
		}
	}
}