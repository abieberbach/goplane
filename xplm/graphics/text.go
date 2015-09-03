//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package graphics

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMGraphics.h>
#include <stdlib.h>
*/
import "C"
import (
	"image/color"
	"unsafe"
	"github.com/abieberbach/goplane"
)

type FontID int

const (
	Font_Basic                           FontID = 0
	Font_Menus                           FontID = 1
	Font_Metal                           FontID = 2
	Font_Led                             FontID = 3
	Font_LedWide                         FontID = 4
	Font_PanelHUD                        FontID = 5
	Font_PanelEFIS                       FontID = 6
	Font_PanelGPS                        FontID = 7
	Font_RadiosGA                        FontID = 8
	Font_RadiosBC                        FontID = 9
	Font_RadiosHM                        FontID = 10
	Font_RadiosGANarrow                  FontID = 11
	Font_RadiosBCNarrow                  FontID = 12
	Font_RadiosHMNarrow                  FontID = 13
	Font_Timer                           FontID = 14
	Font_FullRound                       FontID = 15
	Font_SmallRound                      FontID = 16
	Font_Menus_Localized                 FontID = 17
	Font_Proportional                    FontID = 18
)

func DrawString(color color.Color,xOffset,yOffset int, text string,wordWrapWidth[]int32, fontId FontID) {
	r,g,b,_:=color.RGBA()
	cColor:=[3]float32{float32(r>>8)/255.0,float32(g>>8)/255.0,float32(b>>8)/255.0}
	cText:=C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	var cWordWrapWidth unsafe.Pointer
	if len(wordWrapWidth)>0 {
		cWordWrapWidth=unsafe.Pointer(&wordWrapWidth[0])
	}
	C.XPLMDrawString((*C.float)(unsafe.Pointer(&cColor[0])),C.int(xOffset),C.int(yOffset),cText,(*C.int)(cWordWrapWidth),C.XPLMFontID(fontId))
}

func DrawNumber(color color.Color,xOffset,yOffset int, value float64,digits,decimal int, showSign bool, fontId FontID){
	r,g,b,_:=color.RGBA()
	cColor:=[3]float32{float32(r>>8)/255.0,float32(g>>8)/255.0,float32(b>>8)/255.0}
	C.XPLMDrawNumber((*C.float)(unsafe.Pointer(&cColor[0])),C.int(xOffset),C.int(yOffset),C.double(value),C.int(digits),C.int(decimal),C.int(goplane.FromBoolToInt(showSign)),C.XPLMFontID(fontId))
}

func GetFontDimensions(fontId FontID) (charWidth,charHeight int,digitsOnly bool) {
	var cDigitsOnly int
	C.XPLMGetFontDimensions(C.XPLMFontID(fontId),(*C.int)(unsafe.Pointer(&charWidth)),(*C.int)(unsafe.Pointer(&charHeight)),(*C.int)(unsafe.Pointer(&cDigitsOnly)))
	digitsOnly=cDigitsOnly==1
	return
}

func MeasureString(fontId FontID, text string) float32 {
	cText:=C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	return float32(C.XPLMMeasureString(C.XPLMFontID(fontId),cText,C.int(len(text))))
}
