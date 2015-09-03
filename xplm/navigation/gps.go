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