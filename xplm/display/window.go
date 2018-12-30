//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package display

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMDisplay.h>
#include <stdlib.h>
#include <string.h>

extern void drawWindowCallback(XPLMWindowID inWindowID, void* inRefcon);
extern int windowHandleKey(XPLMWindowID inWindowID,char inChar,XPLMKeyFlags inFlags, char inVirtualKey, void* inRefcon, int losingFocus);
extern int windowHandleMouseClick(XPLMWindowID inWindowID, int x, int y, XPLMMouseStatus inMouse, void* inRefcon);
extern XPLMCursorStatus windowHandleCursorStatus(XPLMWindowID inWindowID, int x, int y, void* inRefcon);
extern int windowHandleMouseWheel(XPLMWindowID inWindowID,int x, int y, int wheel, int clicks, void* inRefcon);


*/
import "C"
import (
	"github.com/abieberbach/goplane"
	"github.com/abieberbach/goplane/xplm"
	"unsafe"
)

type MouseStatus int
type CursorStatus int
type WindowID unsafe.Pointer

type DrawWindow func(id WindowID, ref interface{})
type HandleKey func(id WindowID, keyCode xplm.KeyCode, flags xplm.KeyFlags, virtualKeyCode xplm.VirtualKeyCode, ref interface{}, losingFocus bool) int
type HandleMouseClick func(id WindowID, x, y int, status MouseStatus, ref interface{}) int
type HandleCursorStatus func(id WindowID, x, y int, ref interface{}) CursorStatus
type HandleMouseWheel func(id WindowID, x, y, wheel, clicks int, ref interface{}) int

type CreateWindowData struct {
	Left                 int
	Top                  int
	Right                int
	Bottom               int
	Visible              bool
	DrawWindowFunc       DrawWindow
	HandleMouseClickFunc HandleMouseClick
	HandleKeyFunc        HandleKey
	HandleCursorFunc     HandleCursorStatus
	HandleMouseWheelFunc HandleMouseWheel
	Ref                  interface{}
}

const (
	MouseDown MouseStatus = 1
	MouseDrag MouseStatus = 2
	MouseUp   MouseStatus = 3
)

const (
	CursorDefault CursorStatus = 0
	CursorHidden  CursorStatus = 1
	CursorArrow   CursorStatus = 2
	CursorCustom  CursorStatus = 3
)

var windows = make(map[*C.char]*CreateWindowData)

func GetScreenSize() (width, height int) {
	C.XPLMGetScreenSize((*C.int)(unsafe.Pointer(&width)), (*C.int)(unsafe.Pointer(&height)))
	return
}

func GetMouseLocation() (x, y int) {
	C.XPLMGetMouseLocation((*C.int)(unsafe.Pointer(&x)), (*C.int)(unsafe.Pointer(&y)))
	return
}

//export drawWindowCallback
func drawWindowCallback(windowId C.XPLMWindowID, ref unsafe.Pointer) {
	id := (*C.char)(ref)
	regInfo := windows[id]
	regInfo.DrawWindowFunc(WindowID(windowId), regInfo.Ref)
}

//export windowHandleKey
func windowHandleKey(windowId C.XPLMWindowID, char C.char, flags C.XPLMKeyFlags, virtualKey C.char, ref unsafe.Pointer, losingFocus C.int) C.int {
	id := (*C.char)(ref)
	regInfo := windows[id]
	return C.int(regInfo.HandleKeyFunc(WindowID(windowId), xplm.KeyCode(char), xplm.KeyFlags(flags), xplm.VirtualKeyCode(virtualKey), regInfo.Ref, losingFocus == 1))
}

//export windowHandleMouseClick
func windowHandleMouseClick(windowId C.XPLMWindowID, x, y C.int, mouse C.XPLMMouseStatus, ref unsafe.Pointer) C.int {
	id := (*C.char)(ref)
	regInfo := windows[id]
	return C.int(regInfo.HandleMouseClickFunc(WindowID(windowId), int(x), int(y), MouseStatus(mouse), regInfo.Ref))
}

//export windowHandleCursorStatus
func windowHandleCursorStatus(windowId C.XPLMWindowID, x, y C.int, ref unsafe.Pointer) C.XPLMCursorStatus {
	id := (*C.char)(ref)
	regInfo := windows[id]
	return C.XPLMCursorStatus(regInfo.HandleCursorFunc(WindowID(windowId), int(x), int(y), regInfo.Ref))
}

//export windowHandleMouseWheel
func windowHandleMouseWheel(windowId C.XPLMWindowID, x, y, wheel, clicks C.int, ref unsafe.Pointer) C.int {
	id := (*C.char)(ref)
	regInfo := windows[id]
	return C.int(regInfo.HandleMouseWheelFunc(WindowID(windowId), int(x), int(y), int(wheel), int(clicks), regInfo.Ref))
}

func CreateWindow(left, top, right, bottom int, isVisible bool, drawCallback DrawWindow, keyCallback HandleKey, mouseClickCallback HandleMouseClick, ref interface{}) WindowID {
	regInfo := &CreateWindowData{}
	regInfo.Left = left
	regInfo.Top = top
	regInfo.Right = right
	regInfo.Bottom = bottom
	regInfo.Visible = isVisible
	regInfo.DrawWindowFunc = drawCallback
	regInfo.HandleKeyFunc = keyCallback
	regInfo.HandleMouseClickFunc = mouseClickCallback
	regInfo.Ref = ref

	id := C.CString(goplane.IdGenerator())
	windows[id] = regInfo

	return WindowID(C.XPLMCreateWindow(
		C.int(left),
		C.int(top),
		C.int(right),
		C.int(bottom),
		C.int(goplane.FromBoolToInt(isVisible)),
		C.XPLMDrawWindow_f(unsafe.Pointer(C.drawWindowCallback)),
		C.XPLMHandleKey_f(unsafe.Pointer(C.windowHandleKey)),
		C.XPLMHandleMouseClick_f(unsafe.Pointer(C.windowHandleMouseClick)),
		unsafe.Pointer(id)))
}

func CreateWindowEx(params *CreateWindowData) WindowID {
	id := C.CString(goplane.IdGenerator())
	windows[id] = params
	cCreateWindowData := C.XPLMCreateWindow_t{}
	cCreateWindowData.structSize = C.int(unsafe.Sizeof(cCreateWindowData))
	cCreateWindowData.left = C.int(params.Left)
	cCreateWindowData.top = C.int(params.Top)
	cCreateWindowData.right = C.int(params.Right)
	cCreateWindowData.bottom = C.int(params.Bottom)
	cCreateWindowData.visible = C.int(goplane.FromBoolToInt(params.Visible))
	cCreateWindowData.drawWindowFunc = C.XPLMDrawWindow_f(unsafe.Pointer(C.drawWindowCallback))
	cCreateWindowData.handleMouseClickFunc = C.XPLMHandleMouseClick_f(unsafe.Pointer(C.windowHandleMouseClick))
	cCreateWindowData.handleKeyFunc = C.XPLMHandleKey_f(unsafe.Pointer(C.windowHandleKey))
	cCreateWindowData.handleCursorFunc = C.XPLMHandleCursor_f(unsafe.Pointer(C.windowHandleCursorStatus))
	cCreateWindowData.handleMouseWheelFunc = C.XPLMHandleMouseWheel_f(unsafe.Pointer(C.windowHandleMouseWheel))
	cCreateWindowData.refcon = unsafe.Pointer(id)
	return WindowID(C.XPLMCreateWindowEx((*C.XPLMCreateWindow_t)(unsafe.Pointer(&cCreateWindowData))))
}

func DestroyWindow(windowId WindowID) {
	C.XPLMDestroyWindow(C.XPLMWindowID(windowId))
}

func GetWindowGeometry(windowId WindowID) (left, top, right, bottom int) {
	C.XPLMGetWindowGeometry(
		C.XPLMWindowID(windowId),
		(*C.int)(unsafe.Pointer(&left)),
		(*C.int)(unsafe.Pointer(&top)),
		(*C.int)(unsafe.Pointer(&right)),
		(*C.int)(unsafe.Pointer(&bottom)))
	return
}

func SetWindowGeometry(windowId WindowID, left, top, right, bottom int) {
	C.XPLMSetWindowGeometry(C.XPLMWindowID(windowId), C.int(left), C.int(top), C.int(right), C.int(bottom))
}

func GetWindowIsVisible(windowId WindowID) bool {
	return C.XPLMGetWindowIsVisible(C.XPLMWindowID(windowId)) == 1
}

func SetWindowIsVisible(windowId WindowID, isVisible bool) {
	C.XPLMSetWindowIsVisible(C.XPLMWindowID(windowId), C.int(goplane.FromBoolToInt(isVisible)))
}

func GetWindowRefCon(windowId WindowID) interface{} {
	id := (*C.char)(C.XPLMGetWindowRefCon(C.XPLMWindowID(windowId)))
	return windows[id].Ref
}

func SetWindowRefCon(windowId WindowID, ref interface{}) {
	id := (*C.char)(C.XPLMGetWindowRefCon(C.XPLMWindowID(windowId)))
	windows[id].Ref = ref
}

func TakeKeyboardFocus(windowId WindowID) {
	C.XPLMTakeKeyboardFocus(C.XPLMWindowID(windowId))
}

func BringWindowToFront(windowId WindowID) {
	C.XPLMBringWindowToFront(C.XPLMWindowID(windowId))
}

func IsWindowInFront(windowId WindowID) bool {
	return C.XPLMIsWindowInFront(C.XPLMWindowID(windowId)) == 1
}
