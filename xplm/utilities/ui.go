//Copyright (c) 2015. The goplane AUTHORS. All rights reserved.
//
// Use of this source code is governed by a license that can be found in the LICENSE file.

package utilities
/*
#cgo CFLAGS: -DLIN -DSIMDATA_EXPORTS -DXPLM200=1 -DXPLM210=1
#cgo LDFLAGS: -Xlinker "--unresolved-symbols=ignore-all"
#include <XPLM/XPLMUtilities.h>
#include <stdlib.h>
#include <string.h>

*/
import "C"
import (
	"unsafe"
	"github.com/abieberbach/goplane/xplm"
)


type CommandKeyID int
type CommandButtonID int
type HostApplicationID int
type LanguageCode int
type DataFileType int

type ErrorCallback func(content string)

const (
	Key_pause CommandKeyID = iota
	Key_revthrust
	Key_jettison
	Key_brakesreg
	Key_brakesmax
	Key_gear
	Key_timedn
	Key_timeup
	Key_fadec
	Key_otto_dis
	Key_otto_atr
	Key_otto_asi
	Key_otto_hdg
	Key_otto_gps
	Key_otto_lev
	Key_otto_hnav
	Key_otto_alt
	Key_otto_vvi
	Key_otto_vnav
	Key_otto_nav1
	Key_otto_nav2
	Key_targ_dn
	Key_targ_up
	Key_hdgdn
	Key_hdgup
	Key_barodn
	Key_baroup
	Key_obs1dn
	Key_obs1up
	Key_obs2dn
	Key_obs2up
	Key_com1_1
	Key_com1_2
	Key_com1_3
	Key_com1_4
	Key_nav1_1
	Key_nav1_2
	Key_nav1_3
	Key_nav1_4
	Key_com2_1
	Key_com2_2
	Key_com2_3
	Key_com2_4
	Key_nav2_1
	Key_nav2_2
	Key_nav2_3
	Key_nav2_4
	Key_adf_1
	Key_adf_2
	Key_adf_3
	Key_adf_4
	Key_adf_5
	Key_adf_6
	Key_transpon_1
	Key_transpon_2
	Key_transpon_3
	Key_transpon_4
	Key_transpon_5
	Key_transpon_6
	Key_transpon_7
	Key_transpon_8
	Key_flapsup
	Key_flapsdn
	Key_cheatoff
	Key_cheaton
	Key_sbrkoff
	Key_sbrkon
	Key_ailtrimL
	Key_ailtrimR
	Key_rudtrimL
	Key_rudtrimR
	Key_elvtrimD
	Key_elvtrimU
	Key_forward
	Key_down
	Key_left
	Key_right
	Key_back
	Key_tower
	Key_runway
	Key_chase
	Key_free1
	Key_free2
	Key_spot
	Key_fullscrn1
	Key_fullscrn2
	Key_tanspan
	Key_smoke
	Key_map
	Key_zoomin
	Key_zoomout
	Key_cycledump
	Key_replay
	Key_tranID
	Key_max
)

const (
	Joy_nothing CommandButtonID = iota
	Joy_start_all
	Joy_start_0
	Joy_start_1
	Joy_start_2
	Joy_start_3
	Joy_start_4
	Joy_start_5
	Joy_start_6
	Joy_start_7
	Joy_throt_up
	Joy_throt_dn
	Joy_prop_up
	Joy_prop_dn
	Joy_mixt_up
	Joy_mixt_dn
	Joy_carb_tog
	Joy_carb_on
	Joy_carb_off
	Joy_trev
	Joy_trm_up
	Joy_trm_dn
	Joy_rot_trm_up
	Joy_rot_trm_dn
	Joy_rud_lft
	Joy_rud_cntr
	Joy_rud_rgt
	Joy_ail_lft
	Joy_ail_cntr
	Joy_ail_rgt
	Joy_B_rud_lft
	Joy_B_rud_rgt
	Joy_look_up
	Joy_look_dn
	Joy_look_lft
	Joy_look_rgt
	Joy_glance_l
	Joy_glance_r
	Joy_v_fnh
	Joy_v_fwh
	Joy_v_tra
	Joy_v_twr
	Joy_v_run
	Joy_v_cha
	Joy_v_fr1
	Joy_v_fr2
	Joy_v_spo
	Joy_flapsup
	Joy_flapsdn
	Joy_vctswpfwd
	Joy_vctswpaft
	Joy_gear_tog
	Joy_gear_up
	Joy_gear_down
	Joy_lft_brake
	Joy_rgt_brake
	Joy_brakesREG
	Joy_brakesMAX
	Joy_speedbrake
	Joy_ott_dis
	Joy_ott_atr
	Joy_ott_asi
	Joy_ott_hdg
	Joy_ott_alt
	Joy_ott_vvi
	Joy_tim_start
	Joy_tim_reset
	Joy_ecam_up
	Joy_ecam_dn
	Joy_fadec
	Joy_yaw_damp
	Joy_art_stab
	Joy_chute
	Joy_JATO
	Joy_arrest
	Joy_jettison
	Joy_fuel_dump
	Joy_puffsmoke
	Joy_prerotate
	Joy_UL_prerot
	Joy_UL_collec
	Joy_TOGA
	Joy_shutdown
	Joy_con_atc
	Joy_fail_now
	Joy_pause
	Joy_rock_up
	Joy_rock_dn
	Joy_rock_lft
	Joy_rock_rgt
	Joy_rock_for
	Joy_rock_aft
	Joy_idle_hilo
	Joy_lanlights
	Joy_max
)

const (
	Host_Unknown                        HostApplicationID = 0
	Host_XPlane                         HostApplicationID = 1
	Host_PlaneMaker                     HostApplicationID = 2
	Host_WorldMaker                     HostApplicationID = 3
	Host_Briefer                        HostApplicationID = 4
	Host_PartMaker                      HostApplicationID = 5
	Host_YoungsMod                      HostApplicationID = 6
	Host_XAuto                          HostApplicationID = 7
)

const (
	Language_Unknown                    LanguageCode = 0
	Language_English                    LanguageCode = 1
	Language_French                     LanguageCode = 2
	Language_German                     LanguageCode = 3
	Language_Italian                    LanguageCode = 4
	Language_Spanish                    LanguageCode = 5
	Language_Korean                     LanguageCode = 6
	Language_Russian                    LanguageCode = 7
	Language_Greek                      LanguageCode = 8
	Language_Japanese                   LanguageCode = 9
	Language_Chinese                    LanguageCode = 10
)

const (
	DataFile_Situation                DataFileType = 1
	DataFile_ReplayMovie              DataFileType = 2
)

func SimulateKeyPress(keyType, key int) {
	C.XPLMSimulateKeyPress(C.int(keyType), C.int(key))
}

func SpeakString(text string) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	C.XPLMSpeakString(cText)
}

func CommandKeyStroke(key CommandKeyID) {
	C.XPLMCommandKeyStroke(C.XPLMCommandKeyID(key))
}

func CommandButtonPress(button CommandButtonID) {
	C.XPLMCommandButtonPress(C.XPLMCommandButtonID(button))
}

func CommandButtonRelease(button CommandButtonID) {
	C.XPLMCommandButtonRelease(C.XPLMCommandButtonID(button))
}

func GetVirtualKeyDescription(virtualKey xplm.VirtualKeyCode) string {
	cDesc := C.XPLMGetVirtualKeyDescription(C.char(virtualKey))
	return C.GoString(cDesc)
}