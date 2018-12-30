//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package plugins

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMPlugin.h>
#include <stdlib.h>

extern void featureCallbackFunc(char * inFeature, void * inRef);

*/
import "C"
import (
	"unsafe"
)

type FeatureEnumerator func(string, unsafe.Pointer)

var featureCallback FeatureEnumerator

func HasFeature(featureName string) bool {
	cFeatureName := C.CString(featureName)
	defer C.free(unsafe.Pointer(cFeatureName))
	return C.XPLMHasFeature(cFeatureName) == 1
}

func IsFeatureEnabled(featureName string) bool {
	cFeatureName := C.CString(featureName)
	defer C.free(unsafe.Pointer(cFeatureName))
	return C.XPLMIsFeatureEnabled(cFeatureName) == 1
}

func EnableFeature(featureName string, enabled bool) {
	cFeatureName := C.CString(featureName)
	defer C.free(unsafe.Pointer(cFeatureName))
	isEnabled := 0
	if enabled {
		isEnabled = 1
	}
	C.XPLMEnableFeature(cFeatureName, C.int(isEnabled))
}

//export featureCallbackFunc
func featureCallbackFunc(cFeatureName *C.char, ref unsafe.Pointer) {
	featureName := C.GoString(cFeatureName)
	if featureCallback != nil {
		featureCallback(featureName, ref)
	}
}

func EnumerateFeatures(enumerator FeatureEnumerator, ref unsafe.Pointer) {
	featureCallback = enumerator
	C.XPLMEnumerateFeatures(C.XPLMFeatureEnumerator_f(unsafe.Pointer(C.featureCallbackFunc)), ref)
	featureCallback = nil
}
