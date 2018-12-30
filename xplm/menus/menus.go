//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package menus

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1 -DXPLM300=1 -DXPLM301=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMMenus.h>
#include <stdlib.h>
#include <string.h>

extern void menuHandler(void* menuRef, void* itemRef);

*/
import "C"
import (
	"github.com/abieberbach/goplane"
	"unsafe"
)

type MenuCheck int
type MenuID unsafe.Pointer
type MenuHandler func(menuRef, itemRef interface{})

type regInfo struct {
	handler  MenuHandler
	menuId   MenuID
	menuRef  interface{}
	itemRefs map[*C.char]interface{}
}

const (
	Menu_NoCheck   MenuCheck = 0
	Menu_Unchecked MenuCheck = 1
	Menu_Checked   MenuCheck = 2
)

var handlers = make(map[*C.char]*regInfo)

//export menuHandler
func menuHandler(cMenuRef, cItemRef unsafe.Pointer) {
	info, _ := handlers[(*C.char)(cMenuRef)]
	itemRef := info.itemRefs[(*C.char)(cItemRef)]
	info.handler(info.menuRef, itemRef)
}

func FindPluginsMenu() MenuID {
	return MenuID(C.XPLMFindPluginsMenu())
}

func CreateMenu(name string, parentMenu MenuID, parentItem int, handler MenuHandler, menuRef interface{}) MenuID {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var cHandler C.XPLMMenuHandler_f
	var internalMenuRef *C.char
	if handler != nil {
		internalMenuRef = C.CString(goplane.IdGenerator())
		cHandler = C.XPLMMenuHandler_f(unsafe.Pointer(C.menuHandler))
	}
	id := MenuID(C.XPLMCreateMenu(cName, C.XPLMMenuID(parentMenu), C.int(parentItem), cHandler, unsafe.Pointer(internalMenuRef)))
	if handler != nil {
		handlers[internalMenuRef] = &regInfo{handler, id, menuRef, make(map[*C.char]interface{})}
	}
	return id
}

func DestroyMenu(menuId MenuID) {
	C.XPLMDestroyMenu(C.XPLMMenuID(menuId))
	id, _ := findMenu(menuId)
	delete(handlers, id)
	C.free(unsafe.Pointer(id))
}

func ClearAllMenuItems(menuId MenuID) {
	C.XPLMClearAllMenuItems(C.XPLMMenuID(menuId))
}

func AppendMenuItem(menuId MenuID, itemName string, itemRef interface{}, forceEnglish bool) int {
	cItemName := C.CString(itemName)
	cForceEnglish := 0
	if forceEnglish {
		cForceEnglish = 1
	}
	_, info := findMenu(menuId)
	if info == nil {
		return -1
	}
	info.itemRefs[cItemName] = itemRef
	return int(C.XPLMAppendMenuItem(C.XPLMMenuID(menuId), cItemName, unsafe.Pointer(cItemName), C.int(cForceEnglish)))
}

func AppendMenuSeparator(menuId MenuID) {
	C.XPLMAppendMenuSeparator(C.XPLMMenuID(menuId))
}

func SetMenuItemName(menuId MenuID, index int, itemName string, forceEnglish bool) {
	cItemName := C.CString(itemName)
	cForceEnglish := 0
	if forceEnglish {
		cForceEnglish = 1
	}
	C.XPLMSetMenuItemName(C.XPLMMenuID(menuId), C.int(index), cItemName, C.int(cForceEnglish))
}

func CheckMenuItem(menuId MenuID, index int, check MenuCheck) {
	C.XPLMCheckMenuItem(C.XPLMMenuID(menuId), C.int(index), C.XPLMMenuCheck(check))
}

func CheckMenuItemState(menuId MenuID, index int) MenuCheck {
	var result MenuCheck
	C.XPLMCheckMenuItemState(C.XPLMMenuID(menuId), C.int(index), (*C.XPLMMenuCheck)(unsafe.Pointer(&result)))
	return result
}

func EnableMenuItem(menuId MenuID, index int, enabled bool) {
	cEnabled := 0
	if enabled {
		cEnabled = 1
	}
	C.XPLMEnableMenuItem(C.XPLMMenuID(menuId), C.int(index), C.int(cEnabled))
}

func RemoveMenuItem(menuId MenuID, index int) {
	C.XPLMRemoveMenuItem(C.XPLMMenuID(menuId), C.int(index))
}

func findMenu(menuId MenuID) (*C.char, *regInfo) {
	for key, infos := range handlers {
		if infos.menuId == menuId {
			return key, infos
		}
	}
	return nil, nil
}
