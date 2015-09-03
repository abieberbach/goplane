//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package navigation

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMNavigation.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

func GetGPSDestinationType() NavType {
	return NavType(C.XPLMGetGPSDestinationType())
}

func GetGPSDestination() NavRef {
	return NavRef(C.XPLMGetGPSDestination())
}