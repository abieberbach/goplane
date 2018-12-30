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

extern int getIntData(void * inRef);
extern void setIntData(void * inRef, int value);

extern float getFloatData(void * inRef);
extern void setFloatData(void * inRef, float value);

extern double getDoubleData(void * inRef);
extern void setDoubleData(void * inRef, double value);

extern int getIntArrayData(void* inRef, int* values, int offset, int max);
extern void setIntArrayData(void* inRef, int* values, int offset, int count);

extern int getFloatArrayData(void* inRef, float* values, int offset, int max);
extern void setFloatArrayData(void* inRef, float* values, int offset, int count);

extern int getData(void* inRef, void* values, int offset, int max);
extern void setData(void* inRef, void* values, int offset, int count);

*/
import "C"
import (
	"math"
	"unsafe"
)

type GetIntDataFunc func(ref interface{}) int32
type SetIntDataFunc func(ref interface{}, value int32)
type GetFloatDataFunc func(ref interface{}) float32
type SetFloatDataFunc func(ref interface{}, value float32)
type GetDoubleDataFunc func(ref interface{}) float64
type SetDoubleDataFunc func(ref interface{}, value float64)
type GetIntArrayDataFunc func(ref interface{}) []int32
type SetIntArrayDataFunc func(ref interface{}, value []int32, offset, count int)
type GetFloatArrayDataFunc func(ref interface{}) []float32
type SetFloatArrayDataFunc func(ref interface{}, value []float32, offset, count int)
type GetDataFunc func(ref interface{}) []byte
type SetDataFunc func(ref interface{}, value []byte, offset, count int)

var (
	dataRefs = make(map[unsafe.Pointer]regInfo)
)

type regInfo struct {
	accessors DataRefAccessors
	readRef   interface{}
	writeRef  interface{}
}

type DataRefAccessors struct {
	ReadInt         GetIntDataFunc
	WriteInt        SetIntDataFunc
	ReadFloat       GetFloatDataFunc
	WriteFloat      SetFloatDataFunc
	ReadDouble      GetDoubleDataFunc
	WriteDouble     SetDoubleDataFunc
	ReadIntArray    GetIntArrayDataFunc
	WriteIntArray   SetIntArrayDataFunc
	ReadFloatArray  GetFloatArrayDataFunc
	WriteFloatArray SetFloatArrayDataFunc
	ReadData        GetDataFunc
	WriteData       SetDataFunc
}

//export getIntData
func getIntData(ref unsafe.Pointer) C.int {
	value := int32(0)
	regInfo, _ := dataRefs[ref]
	value = regInfo.accessors.ReadInt(regInfo.readRef)
	return C.int(value)
}

//export setIntData
func setIntData(ref unsafe.Pointer, value C.int) {
	regInfo, _ := dataRefs[ref]
	regInfo.accessors.WriteInt(regInfo.writeRef, int32(value))
}

//export getFloatData
func getFloatData(ref unsafe.Pointer) C.float {
	value := float32(0.0)
	regInfo, _ := dataRefs[ref]
	value = regInfo.accessors.ReadFloat(regInfo.readRef)
	return C.float(value)
}

//export setFloatData
func setFloatData(ref unsafe.Pointer, value C.float) {
	regInfo, _ := dataRefs[ref]
	regInfo.accessors.WriteFloat(regInfo.writeRef, float32(value))
}

//export getDoubleData
func getDoubleData(ref unsafe.Pointer) C.double {
	value := 0.0
	regInfo, _ := dataRefs[ref]
	value = regInfo.accessors.ReadDouble(regInfo.readRef)
	return C.double(value)
}

//export setDoubleData
func setDoubleData(ref unsafe.Pointer, value C.double) {
	regInfo, _ := dataRefs[ref]
	regInfo.accessors.WriteDouble(regInfo.writeRef, float64(value))
}

//export getIntArrayData
func getIntArrayData(ref unsafe.Pointer, cValues *C.int, offset, max C.int) C.int {
	length := 0
	regInfo, _ := dataRefs[ref]
	value := regInfo.accessors.ReadIntArray(regInfo.readRef)
	length = len(value)
	if cValues != nil {
		length = int(math.Min(float64(length), float64(max)))
		C.memcpy(unsafe.Pointer(cValues), unsafe.Pointer(&value[offset]), C.size_t(int(unsafe.Sizeof(value[0]))*length))
	}
	return C.int(length)
}

//export setIntArrayData
func setIntArrayData(ref unsafe.Pointer, cValues *C.int, offset, count C.int) {
	regInfo, _ := dataRefs[ref]
	value := make([]int32, int(count))
	C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(cValues), C.size_t(int(unsafe.Sizeof(value[0]))*int(count)))
	regInfo.accessors.WriteIntArray(regInfo.writeRef, value, int(offset), int(count))
}

//export getFloatArrayData
func getFloatArrayData(ref unsafe.Pointer, cValues *C.float, offset, max C.int) C.int {
	length := 0
	regInfo, _ := dataRefs[ref]
	value := regInfo.accessors.ReadFloatArray(regInfo.readRef)
	length = len(value)
	if cValues != nil {
		length = int(math.Min(float64(length), float64(max)))
		C.memcpy(unsafe.Pointer(cValues), unsafe.Pointer(&value[offset]), C.size_t(int(unsafe.Sizeof(value[0]))*length))
	}
	return C.int(length)
}

//export setFloatArrayData
func setFloatArrayData(ref unsafe.Pointer, cValues *C.float, offset, count C.int) {
	regInfo, _ := dataRefs[ref]
	value := make([]float32, int(count))
	C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(cValues), C.size_t(int(unsafe.Sizeof(value[0]))*int(count)))
	regInfo.accessors.WriteFloatArray(regInfo.writeRef, value, int(offset), int(count))
}

//export getData
func getData(ref unsafe.Pointer, cValues unsafe.Pointer, offset, max C.int) C.int {
	length := 0
	regInfo, _ := dataRefs[ref]
	value := regInfo.accessors.ReadData(regInfo.readRef)
	length = len(value)
	if cValues != nil {
		length = int(math.Min(float64(length), float64(max)))
		C.memcpy(cValues, unsafe.Pointer(&value[offset]), C.size_t(length))
	}
	return C.int(length)
}

//export setData
func setData(ref unsafe.Pointer, cValues unsafe.Pointer, offset, count C.int) {
	regInfo, _ := dataRefs[ref]
	value := make([]byte, int(count))
	C.memcpy(unsafe.Pointer(&value[0]), unsafe.Pointer(cValues), C.size_t(count))
	regInfo.accessors.WriteData(regInfo.writeRef, value, int(offset), int(count))
}

func RegisterDataAccessor(name string, dataType DataRefType, isWriteable bool, accessors DataRefAccessors, readRef, writeRef interface{}) DataRef {
	cName := C.CString(name)
	dataRefs[unsafe.Pointer(cName)] = regInfo{accessors, readRef, writeRef}
	writable := 0
	if isWriteable {
		writable = 1
	}

	var (
		intGetter        C.XPLMGetDatai_f
		intSetter        C.XPLMSetDatai_f
		floatGetter      C.XPLMGetDataf_f
		floatSetter      C.XPLMSetDataf_f
		doubleGetter     C.XPLMGetDatad_f
		doubleSetter     C.XPLMSetDatad_f
		intArrayGetter   C.XPLMGetDatavi_f
		intArraySetter   C.XPLMSetDatavi_f
		floatArrayGetter C.XPLMGetDatavf_f
		floatArraySetter C.XPLMSetDatavf_f
		dataGetter       C.XPLMGetDatab_f
		dataSetter       C.XPLMSetDatab_f
	)

	if accessors.ReadInt != nil {
		intGetter = C.XPLMGetDatai_f(unsafe.Pointer(C.getIntData))
	}
	if accessors.WriteInt != nil {
		intSetter = C.XPLMSetDatai_f(unsafe.Pointer(C.setIntData))
	}
	if accessors.ReadFloat != nil {
		floatGetter = C.XPLMGetDataf_f(unsafe.Pointer(C.getFloatData))
	}
	if accessors.WriteFloat != nil {
		floatSetter = C.XPLMSetDataf_f(unsafe.Pointer(C.setFloatData))
	}
	if accessors.ReadDouble != nil {
		doubleGetter = C.XPLMGetDatad_f(unsafe.Pointer(C.getDoubleData))
	}
	if accessors.WriteDouble != nil {
		doubleSetter = C.XPLMSetDatad_f(unsafe.Pointer(C.setDoubleData))
	}
	if accessors.ReadIntArray != nil {
		intArrayGetter = C.XPLMGetDatavi_f(unsafe.Pointer(C.getIntArrayData))
	}
	if accessors.WriteIntArray != nil {
		intArraySetter = C.XPLMSetDatavi_f(unsafe.Pointer(C.setIntArrayData))
	}
	if accessors.ReadFloatArray != nil {
		floatArrayGetter = C.XPLMGetDatavf_f(unsafe.Pointer(C.getFloatArrayData))
	}
	if accessors.WriteFloatArray != nil {
		floatArraySetter = C.XPLMSetDatavf_f(unsafe.Pointer(C.setFloatArrayData))
	}
	if accessors.ReadData != nil {
		dataGetter = C.XPLMGetDatab_f(unsafe.Pointer(C.getData))
	}
	if accessors.WriteData != nil {
		dataSetter = C.XPLMSetDatab_f(unsafe.Pointer(C.setData))
	}

	return DataRef(C.XPLMRegisterDataAccessor(
		cName,
		C.XPLMDataTypeID(dataType),
		C.int(writable),
		intGetter,
		intSetter,
		floatGetter,
		floatSetter,
		doubleGetter,
		doubleSetter,
		intArrayGetter,
		intArraySetter,
		floatArrayGetter,
		floatArraySetter,
		dataGetter,
		dataSetter,
		unsafe.Pointer(cName),
		unsafe.Pointer(cName)))
}

func UnregisterDataAccessor(dataRef DataRef) {
	C.XPLMUnregisterDataAccessor(C.XPLMDataRef(dataRef))
}
