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

//Basisstruktur für ein X-Plane Plugin.
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

//Status eines Plugins für die Callback-Funktion.
type PluginState int

const (
	PluginStart PluginState = 0 //Plugin wird gestartet
	PluginEnable PluginState = 1 //Plugin wird aktiviert
	PluginDisable PluginState = 2 //Plugin wird deaktiviert
	PluginStop PluginState = 3 //Plugin wird gestoppt
)

//Callback-Funktion für den Statuswechsel eines Plugins.
type PluginStateCallback func(state PluginState, plugin *XPlanePlugin)

//Messagehandler der die empfange Nachricht als Debugmeldung schreibt.
func DebugMessageHandler(msg plugins.Message) {
	name, _, _, _ := plugins.GetPluginInfo(msg.PluginId)
	logging.Debug(fmt.Sprintf("receive message from %v (ID: %v): %v", name, msg.PluginId, msg.MessageId))
}

var (
	plugin *XPlanePlugin //globale Plugin-Instanz
)

//Erzeugt ein neues Plugin.
func NewPlugin(name, signature, description string) *XPlanePlugin {
	logging.PluginName = name
	plugin = &XPlanePlugin{plugins.NO_PLUGIN_ID, name, signature, description, nil, nil, 1.0, nil, nil}
	logging.Info("================================================================")
	logging.Info(fmt.Sprintf("Plugin %v initialized", name))
	logging.Info(fmt.Sprintf("  signature: %v", signature))
	logging.Info(fmt.Sprintf("  description: %v", description))
	logging.Info("================================================================")
	return plugin
}

//Liefert die ID des Plugins.
func (self *XPlanePlugin) GetId() plugins.PluginId {
	if self.id == -plugins.NO_PLUGIN_ID {
		self.id = plugins.GetMyId()
	}
	return self.id
}

//Liefert den Namen des Plugins.
func (self *XPlanePlugin) GetName() string {
	return self.name
}

//Liefert die Beschreibung des Plugins.
func (self *XPlanePlugin) GetDescription() string {
	return self.description
}

//Liefert die Signatur des Plugins.
func (self *XPlanePlugin) GetSignature() string {
	return self.signature
}

//Liefert den MessageHandler des Plugins.
func (self *XPlanePlugin) GetMessageHandler() plugins.MessageHandler {
	return self.messageHandler
}

//Setzt den MessageHander des Plugins.
func (self *XPlanePlugin) SetMessageHandler(handler plugins.MessageHandler) {
	self.messageHandler = handler
}

//Setzt die FlightLoop-Funktion, die beim Start des Plugins automatisch registriert werden soll.
func (self *XPlanePlugin) SetFlightLoopFunc(flightLoopFunc processing.FlightLoopFunc, interval float32) {
	self.flightLoop = flightLoopFunc
	self.flightLoopInterval = interval
}

//Setzt die ErrorCallback-Funktion, die beim Start des Plugins automatisch registriert werden soll.
func (self *XPlanePlugin) SetErrorCallback(callback utilities.ErrorCallback) {
	self.errorCallback = callback
}

//Setzt die Callback-Funktion für den Pluginstatuswechsel
func (self *XPlanePlugin) SetPluginStateCallback(callback PluginStateCallback) {
	self.stateCallback = callback
}

func (self *XPlanePlugin) onStart(name, sig, desc *C.char) {
	copyStringToCPointer(self.name, name)
	copyStringToCPointer(self.signature, sig)
	copyStringToCPointer(self.description, desc)
	if self.errorCallback != nil {
		utilities.SetErrorCallback(self.errorCallback)
	}
	if self.flightLoop != nil {
		processing.RegisterFlightLoopCallback(self.flightLoop, self.flightLoopInterval, self)
	}
}

func copyStringToCPointer(text string, target *C.char) {
	cMsg := C.CString(text)
	defer C.free(unsafe.Pointer(cMsg))
	C.strcpy(target, cMsg);
}


func (self *XPlanePlugin) String() string {
	return fmt.Sprintf("%v (singature: %v, id: %v)", self.GetName(), self.GetSignature(), self.GetId())
}

//Externe Schnittstellen-Methode, die von X-Plane beim Empfang einer Nachricht angesprochen wird
//export XPluginReceiveMessage
func XPluginReceiveMessage(pluginId C.int, messageId C.int, messageData unsafe.Pointer) {
	if plugin.messageHandler != nil {
		plugin.messageHandler(plugins.Message{plugins.PluginId(pluginId), plugins.MessageId(messageId), messageData})
	}
}

//Externe Schnittstellen-Methode, die von X-Plane beim Start des Plugins angesprochen wird
//export XPluginStart
func XPluginStart(outName *C.char, outSig *C.char, outDesc *C.char) int {
	plugin.onStart(outName, outSig, outDesc)
	if plugin.stateCallback != nil {
		plugin.stateCallback(PluginStart, plugin)
	}
	return 1
}


//Externe Schnittstellen-Methode, die von X-Plane beim Aktivieren des Plugins angesprochen wird
//export XPluginEnable
func XPluginEnable() {
	if plugin.stateCallback != nil {
		plugin.stateCallback(PluginEnable, plugin)
	}
}

//Externe Schnittstellen-Methode, die von X-Plane beim Deaktivieren des Plugins angesprochen wird
//export XPluginDisable
func XPluginDisable() {
	if plugin.stateCallback != nil {
		plugin.stateCallback(PluginDisable, plugin)
	}

}

//Externe Schnittstellen-Methode, die von X-Plane beim Stoppen des Plugins angesprochen wird
//export XPluginStop
func XPluginStop() {
	if plugin.stateCallback != nil {
		plugin.stateCallback(PluginStop, plugin)
	}
}
