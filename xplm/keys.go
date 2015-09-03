//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package xplm

/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMDisplay.h>
#include <stdlib.h>
#include <string.h>
*/

type KeyFlags int
type KeyCode int
type VirtualKeyCode int

const (
	ShiftFlag KeyFlags = 1
	OptionAltFlag KeyFlags = 2
	ControlFlag KeyFlags = 4
	DownFlag KeyFlags = 8
	UpFlag KeyFlags = 16
)

const (
	KEY_RETURN KeyCode = 13
	KEY_ESCAPE KeyCode = 27
	KEY_TAB KeyCode = 9
	KEY_DELETE KeyCode = 8
	KEY_LEFT  KeyCode = 28
	KEY_RIGHT KeyCode = 29
	KEY_UP  KeyCode = 30
	KEY_DOWN KeyCode = 31
	KEY_0 KeyCode = 48
	KEY_1 KeyCode = 49
	KEY_2 KeyCode = 50
	KEY_3 KeyCode = 51
	KEY_4 KeyCode = 52
	KEY_5 KeyCode = 53
	KEY_6 KeyCode = 54
	KEY_7 KeyCode = 55
	KEY_8 KeyCode = 56
	KEY_9 KeyCode = 57
	KEY_DECIMAL KeyCode = 46
)

const (
	VK_BACK VirtualKeyCode = 0x08
	VK_TAB VirtualKeyCode = 0x09
	VK_CLEAR VirtualKeyCode = 0x0C
	VK_RETURN  VirtualKeyCode = 0x0D
	VK_ESCAPE   VirtualKeyCode = 0x1B
	VK_SPACE VirtualKeyCode = 0x20
	VK_PRIOR VirtualKeyCode = 0x21
	VK_NEXT VirtualKeyCode = 0x22
	VK_END VirtualKeyCode = 0x23
	VK_HOME VirtualKeyCode = 0x24
	VK_LEFT VirtualKeyCode = 0x25
	VK_UP VirtualKeyCode = 0x26
	VK_RIGHT VirtualKeyCode = 0x27
	VK_DOWN VirtualKeyCode = 0x28
	VK_SELECT VirtualKeyCode = 0x29
	VK_PRINT VirtualKeyCode = 0x2A
	VK_EXECUTE VirtualKeyCode = 0x2B
	VK_SNAPSHOT VirtualKeyCode = 0x2C
	VK_INSERT  VirtualKeyCode = 0x2D
	VK_DELETE VirtualKeyCode = 0x2E
	VK_HELP VirtualKeyCode = 0x2F
	VK_0 VirtualKeyCode = 0x30
	VK_1 VirtualKeyCode = 0x31
	VK_2 VirtualKeyCode = 0x32
	VK_3 VirtualKeyCode = 0x33
	VK_4 VirtualKeyCode = 0x34
	VK_5 VirtualKeyCode = 0x35
	VK_6 VirtualKeyCode = 0x36
	VK_7 VirtualKeyCode = 0x37
	VK_8 VirtualKeyCode = 0x38
	VK_9 VirtualKeyCode = 0x39
	VK_A VirtualKeyCode = 0x41
	VK_B VirtualKeyCode = 0x42
	VK_C VirtualKeyCode = 0x43
	VK_D VirtualKeyCode = 0x44
	VK_E VirtualKeyCode = 0x45
	VK_F VirtualKeyCode = 0x46
	VK_G VirtualKeyCode = 0x47
	VK_H VirtualKeyCode = 0x48
	VK_I VirtualKeyCode = 0x49
	VK_J VirtualKeyCode = 0x4A
	VK_K VirtualKeyCode = 0x4B
	VK_L VirtualKeyCode = 0x4C
	VK_M VirtualKeyCode = 0x4D
	VK_N VirtualKeyCode = 0x4E
	VK_O VirtualKeyCode = 0x4F
	VK_P VirtualKeyCode = 0x50
	VK_Q VirtualKeyCode = 0x51
	VK_R VirtualKeyCode = 0x52
	VK_S VirtualKeyCode = 0x53
	VK_T VirtualKeyCode = 0x54
	VK_U VirtualKeyCode = 0x55
	VK_V VirtualKeyCode = 0x56
	VK_W VirtualKeyCode = 0x57
	VK_X VirtualKeyCode = 0x58
	VK_Y VirtualKeyCode = 0x59
	VK_Z VirtualKeyCode = 0x5A
	VK_NUMPAD0 VirtualKeyCode = 0x60
	VK_NUMPAD1 VirtualKeyCode = 0x61
	VK_NUMPAD2 VirtualKeyCode = 0x62
	VK_NUMPAD3 VirtualKeyCode = 0x63
	VK_NUMPAD4 VirtualKeyCode = 0x64
	VK_NUMPAD5 VirtualKeyCode = 0x65
	VK_NUMPAD6 VirtualKeyCode = 0x66
	VK_NUMPAD7 VirtualKeyCode = 0x67
	VK_NUMPAD8 VirtualKeyCode = 0x68
	VK_NUMPAD9 VirtualKeyCode = 0x69
	VK_MULTIPLY VirtualKeyCode = 0x6A
	VK_ADD VirtualKeyCode = 0x6B
	VK_SEPARATOR VirtualKeyCode = 0x6C
	VK_SUBTRACT VirtualKeyCode = 0x6D
	VK_DECIMAL VirtualKeyCode = 0x6E
	VK_DIVIDE VirtualKeyCode = 0x6F
	VK_F1 VirtualKeyCode = 0x70
	VK_F2 VirtualKeyCode = 0x71
	VK_F3 VirtualKeyCode = 0x72
	VK_F4 VirtualKeyCode = 0x73
	VK_F5 VirtualKeyCode = 0x74
	VK_F6 VirtualKeyCode = 0x75
	VK_F7 VirtualKeyCode = 0x76
	VK_F8 VirtualKeyCode = 0x77
	VK_F9 VirtualKeyCode = 0x78
	VK_F10 VirtualKeyCode = 0x79
	VK_F11 VirtualKeyCode = 0x7A
	VK_F12 VirtualKeyCode = 0x7B
	VK_F13 VirtualKeyCode = 0x7C
	VK_F14 VirtualKeyCode = 0x7D
	VK_F15 VirtualKeyCode = 0x7E
	VK_F16 VirtualKeyCode = 0x7F
	VK_F17 VirtualKeyCode = 0x80
	VK_F18 VirtualKeyCode = 0x81
	VK_F19 VirtualKeyCode = 0x82
	VK_F20 VirtualKeyCode = 0x83
	VK_F21 VirtualKeyCode = 0x84
	VK_F22 VirtualKeyCode = 0x85
	VK_F23 VirtualKeyCode = 0x86
	VK_F24 VirtualKeyCode = 0x87
	VK_EQUAL VirtualKeyCode = 0xB0
	VK_MINUS VirtualKeyCode = 0xB1
	VK_RBRACE VirtualKeyCode = 0xB2
	VK_LBRACE VirtualKeyCode = 0xB3
	VK_QUOTE VirtualKeyCode = 0xB4
	VK_SEMICOLON VirtualKeyCode = 0xB5
	VK_BACKSLASH VirtualKeyCode = 0xB6
	VK_COMMA VirtualKeyCode = 0xB7
	VK_SLASH VirtualKeyCode = 0xB8
	VK_PERIOD VirtualKeyCode = 0xB9
	VK_BACKQUOTE VirtualKeyCode = 0xBA
	VK_ENTER VirtualKeyCode = 0xBB
	VK_NUMPAD_ENT VirtualKeyCode = 0xBC
	VK_NUMPAD_EQ VirtualKeyCode = 0xBD
)