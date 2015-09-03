//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package graphics

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMGraphics.h>
*/
import "C"
import (
	"unsafe"
	"github.com/abieberbach/goplane"
)

type TextureID int

const (
	Tex_GeneralInterface TextureID = 0
	Tex_AircraftPaint TextureID = 1
	Tex_AircraftLiteMap  TextureID = 2

)

func SetGraphicsState(enableFog bool, numberTexUnits int, enableLighting, enableAlphaTesting, enableAlphaBlending, enableDepthTesting, enableDepthWriting bool) {
	C.XPLMSetGraphicsState(
		C.int(goplane.FromBoolToInt(enableFog)),
		C.int(numberTexUnits),
		C.int(goplane.FromBoolToInt(enableLighting)),
		C.int(goplane.FromBoolToInt(enableAlphaTesting)),
		C.int(goplane.FromBoolToInt(enableAlphaBlending)),
		C.int(goplane.FromBoolToInt(enableDepthTesting)),
		C.int(goplane.FromBoolToInt(enableDepthWriting)))
}

func BindTexture2d(textureNum,textureUnit int32) {
	C.XPLMBindTexture2d(C.int(textureNum),C.int(textureUnit))
}

func GenerateTextureNumbers(count int) []int32 {
	result:=make([]int32,count)
	C.XPLMGenerateTextureNumbers((*C.int)(unsafe.Pointer(&result[0])),C.int(count))
	return result
}

func GetTexture(textureId TextureID) int {
	return int(C.XPLMGetTexture(C.XPLMTextureID(textureId)))
}

func LocalToWorld(x, y, z float64) (lat, lon, alt float64) {
C.XPLMLocalToWorld(C.double(x), C.double(y), C.double(z), (*C.double)(unsafe.Pointer(&lat)), (*C.double)(unsafe.Pointer(&lon)), (*C.double)(unsafe.Pointer(&alt)))
return
}

func WorldToLocal(lat, lon, alt float64) (x, y, z float64) {
C.XPLMWorldToLocal(C.double(lat), C.double(lon), C.double(alt), (*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)), (*C.double)(unsafe.Pointer(&z)))
return
}

func DrawTranslucentDarkBox(left,top,right,bottom int) {
	C.XPLMDrawTranslucentDarkBox(C.int(left),C.int(top),C.int(right),C.int(bottom))
}