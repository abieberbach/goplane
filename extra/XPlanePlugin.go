package extra

/*
#include <stdlib.h>
#include <string.h>
 */
import "C"
import (
	"fmt"
	"github.com/abieberbach/goplane/xplm/plugins"
	"github.com/abieberbach/goplane/xplm/processing"
	"github.com/abieberbach/goplane/xplm/utilities"
	"unsafe"
	"github.com/abieberbach/goplane/extra/logging"
)

type XPlanePlugin struct {
	id                 plugins.PluginId
	name               string
	signature          string
	description        string
	messageHandler     plugins.MessageHandler
	flightLoop         processing.FlightLoopFunc
	flightLoopInterval float32
	errorCallback      utilities.ErrorCallback
	stateCallback      PluginStateCallback
}

type PluginState int

const (
	PluginStart PluginState = 0
	PluginEnable PluginState = 1
	PluginDisable PluginState = 2
	PluginStop PluginState = 3
)

type PluginStateCallback func(state PluginState, plugin *XPlanePlugin)

func DebugMessageHandler(msg plugins.Message) {
	name, _, _, _ := plugins.GetPluginInfo(msg.PluginId)
	logging.Debug(fmt.Sprintf("receive message from %v (ID: %v): %v", name, msg.PluginId, msg.MessageId))
}

var (
	plugin *XPlanePlugin //globale Plugin-Instanz
)


func NewPlugin(name, signature, description string) *XPlanePlugin {
	logging.PluginName=name
	plugin=&XPlanePlugin{plugins.NO_PLUGIN_ID, name, signature, description, nil, nil, 1.0, nil, nil}
	logging.Info("================================================================")
	logging.Info(fmt.Sprintf("Plugin %v initialized", name))
	logging.Info(fmt.Sprintf("  signature: %v", signature))
	logging.Info(fmt.Sprintf("  description: %v", description))
	logging.Info("================================================================")
	return plugin
}

func (self *XPlanePlugin) GetId() plugins.PluginId {
	if self.id==-plugins.NO_PLUGIN_ID {
		self.id=plugins.GetMyId()
	}
	return self.id
}

func (self *XPlanePlugin) GetName() string {
	return self.name
}

func (self *XPlanePlugin) GetDescription() string {
	return self.description
}

func (self *XPlanePlugin) GetSignature() string {
	return self.signature
}

func (self *XPlanePlugin) GetMessageHandler() plugins.MessageHandler {
	return self.messageHandler
}

func (self *XPlanePlugin) SetMessageHandler(handler plugins.MessageHandler) {
	self.messageHandler=handler
}

func (self *XPlanePlugin) SetFlightLoopFunc(flightLoopFunc processing.FlightLoopFunc, interval float32) {
	self.flightLoop=flightLoopFunc
	self.flightLoopInterval=interval
}

func (self *XPlanePlugin) SetErrorCallback(callback utilities.ErrorCallback) {
	self.errorCallback=callback
}

func (self *XPlanePlugin) SetPluginStateCallback(callback PluginStateCallback) {
	self.stateCallback=callback
}

func (self *XPlanePlugin) onStart(name, sig, desc *C.char) {
	copyStringToCPointer(self.name, name)
	copyStringToCPointer(self.signature, sig)
	copyStringToCPointer(self.description, desc)
	if self.errorCallback!=nil {
		utilities.SetErrorCallback(self.errorCallback)
	}
	if self.flightLoop!=nil {
		processing.RegisterFlightLoopCallback(self.flightLoop, self.flightLoopInterval, self)
	}
}

func copyStringToCPointer(text string, target *C.char) {
	cMsg := C.CString(text)
	defer C.free(unsafe.Pointer(cMsg))
	C.strcpy(target, cMsg);
}


func (self *XPlanePlugin) String() string {
	return fmt.Sprintf("%v (singature: %v, id: %v)",self.GetName(),self.GetSignature(),self.GetId())
}

//export XPluginReceiveMessage
func XPluginReceiveMessage(pluginId C.int, messageId C.int, messageData unsafe.Pointer) {
	if plugin.messageHandler!=nil {
		plugin.messageHandler(plugins.Message{plugins.PluginId(pluginId), plugins.MessageId(messageId), messageData})
	}
}

//export XPluginStart
func XPluginStart(outName *C.char, outSig *C.char, outDesc *C.char) int {
	plugin.onStart(outName, outSig, outDesc)
	if plugin.stateCallback!=nil {
		plugin.stateCallback(PluginStart,plugin)
	}
	return 1
}


//export XPluginEnable
func XPluginEnable() {
	if plugin.stateCallback!=nil {
		plugin.stateCallback(PluginEnable,plugin)
	}
}

//export XPluginDisable
func XPluginDisable() {
	if plugin.stateCallback!=nil {
		plugin.stateCallback(PluginDisable,plugin)
	}

}

//export XPluginStop
func XPluginStop() {
	if plugin.stateCallback!=nil {
		plugin.stateCallback(PluginStop,plugin)
	}
}
