//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package dataAccess

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMDataAccess.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"unsafe"
)

type DataRef C.XPLMDataRef
type DataRefType int

const (
	TypeUnknown    DataRefType = 0
	TypeInt        DataRefType = 1
	TypeFloat      DataRefType = 2
	TypeDouble     DataRefType = 4
	TypeFloatArray DataRefType = 8
	TypeIntArray   DataRefType = 16
	TypeData       DataRefType = 32
)

func FindDataRef(dataRefName string) (DataRef, bool) {
	cDataRefName := C.CString(dataRefName)
	defer C.free(unsafe.Pointer(cDataRefName))
	dataRef := C.XPLMFindDataRef(cDataRefName)
	return DataRef(dataRef), dataRef != nil
}

func CanWriteDataRef(dataRef DataRef) bool {
	return C.XPLMCanWriteDataRef(C.XPLMDataRef(dataRef)) == 1
}

func IsDataRefGood(dataRef DataRef) bool {
	return C.XPLMIsDataRefGood(C.XPLMDataRef(dataRef)) == 1
}

func GetDataRefTypes(dataRef DataRef) DataRefType {
	return DataRefType(C.XPLMGetDataRefTypes(C.XPLMDataRef(dataRef)))
}

func GetIntData(dataRef DataRef) int {
	return int(C.XPLMGetDatai(C.XPLMDataRef(dataRef)))
}

func SetIntData(dataRef DataRef, value int) {
	C.XPLMSetDatai(C.XPLMDataRef(dataRef), C.int(value))
}

func GetFloatData(dataRef DataRef) float32 {
	return float32(C.XPLMGetDataf(C.XPLMDataRef(dataRef)))
}

func SetFloatData(dataRef DataRef, value float32) {
	C.XPLMSetDataf(C.XPLMDataRef(dataRef), C.float(value))
}

func GetDoubleData(dataRef DataRef) float64 {
	return float64(C.XPLMGetDatad(C.XPLMDataRef(dataRef)))
}

func SetDoubleData(dataRef DataRef, value float64) {
	C.XPLMSetDatad(C.XPLMDataRef(dataRef), C.double(value))
}

func GetIntArrayData(dataRef DataRef) []int {
	length := int(C.XPLMGetDatavi(C.XPLMDataRef(dataRef), nil, 0, 0))
	result := make([]int, length)
	C.XPLMGetDatavi(C.XPLMDataRef(dataRef), (*C.int)(unsafe.Pointer(&result[0])), 0, C.int(length))
	return result
}

func SetIntArrayData(dataRef DataRef, value []int) {
	C.XPLMSetDatavi(C.XPLMDataRef(dataRef), (*C.int)(unsafe.Pointer(&value[0])), 0, C.int(len(value)))
}

func GetFloatArrayData(dataRef DataRef) []float32 {
	length := int(C.XPLMGetDatavf(C.XPLMDataRef(dataRef), nil, 0, 0))
	result := make([]float32, length)
	C.XPLMGetDatavf(C.XPLMDataRef(dataRef), (*C.float)(unsafe.Pointer(&result[0])), 0, C.int(length))
	return result
}

func SetFloatArrayData(dataRef DataRef, value []float32) {
	C.XPLMSetDatavf(C.XPLMDataRef(dataRef), (*C.float)(unsafe.Pointer(&value[0])), 0, C.int(len(value)))
}

func GetData(dataRef DataRef) []byte {
	length := int(C.XPLMGetDatab(C.XPLMDataRef(dataRef), nil, 0, 0))
	result := make([]byte, length)
	C.XPLMGetDatab(C.XPLMDataRef(dataRef), unsafe.Pointer(&result[0]), 0, C.int(length))
	return result
}

func SetData(dataRef DataRef, value []byte) {
	C.XPLMSetDatab(C.XPLMDataRef(dataRef), unsafe.Pointer(&value[0]), 0, C.int(len(value)))
}

func GetString(dataRef DataRef) string {
	length := int(C.XPLMGetDatab(C.XPLMDataRef(dataRef), nil, 0, 0))
	valueBuffer := (*C.char)(C.malloc(C.size_t(length)))
	defer C.free(unsafe.Pointer(valueBuffer))
	C.XPLMGetDatab(C.XPLMDataRef(dataRef), unsafe.Pointer(valueBuffer), 0, C.int(length))
	return C.GoString(valueBuffer)
}

func SetString(dataRef DataRef, value string) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	C.XPLMSetDatab(C.XPLMDataRef(dataRef), unsafe.Pointer(cValue), 0, C.int(C.strlen(cValue)))
}
