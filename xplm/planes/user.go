package planes

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMPlanes.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

func SetUsersAircraft(aircraft string) {
	cPath := C.CString(aircraft)
	defer C.free(unsafe.Pointer(cPath))
	C.XPLMSetUsersAircraft(cPath)
}

func PlaceUserAtAirport(airportCode string) {
	cAirportCode := C.CString(airportCode)
	defer C.free(unsafe.Pointer(cAirportCode))
	C.XPLMPlaceUserAtAirport(cAirportCode)
}
