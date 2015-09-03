package logging

import (
	"fmt"
	"github.com/abieberbach/goplane/xplm/utilities"
)

//Informationen über ein Loglevel
type Level struct {
	number byte   //Nummer des Loglevels
	name   string //Name des Loglevels
}

var (
//Loglevel für Debugmeldungen
	Debug_Level = Level{1, "DEBUG"}
//Loglevel für Infomeldungen
	Info_Level = Level{2, "INFO"}
//Loglevel für Warnungen
	Warning_Level = Level{3, "WARNING"}
//Loglevel für Fehler
	Error_Level = Level{4, "ERROR"}

//Level ab dem die Meldungen ausgegeben werden
	MinLevel = Info_Level
//aktueller Pluginname
	PluginName = "<unknown>"
)

//Schreibt eine Debugmeldung in die Logdatei
func Debug(msg string) {
	writeMessage(Debug_Level, msg)
}

//Schreibt eine formatierte Debugmeldung in die Logdatei
func Debugf(format string, a... interface{}) {
	if Debug_Level.number >= MinLevel.number {
		Debug(fmt.Sprintf(format, a...))
	}
}

//Schreibt eine Infomeldung in die Logdatei
func Info(msg string) {
	writeMessage(Info_Level, msg)
}

//Schreibt eine formatierte Infomeldung in die Logdatei
func Infof(format string, a... interface{}) {
	if Info_Level.number >= MinLevel.number {
		Info(fmt.Sprintf(format, a...))
	}
}

//Schreibt eine Warnung in die Logdatei
func Warning(msg string) {
	writeMessage(Warning_Level, msg)
}

//Schreibt eine formatierte Warnung in die Logdatei
func Warningf(format string, a... interface{}) {
	if Warning_Level.number >= MinLevel.number {
		Warning(fmt.Sprintf(format, a...))
	}
}


//Schreibt eine Fehlermeldung in die Logdatei
func Error(msg string) {
	writeMessage(Error_Level, msg)
}

//Schreibt eine formatierte Fehlermeldung in die Logdatei
func Errorf(format string, a... interface{}) {
	if Error_Level.number >= MinLevel.number {
		Error(fmt.Sprintf(format, a...))
	}
}

func writeMessage(level Level, msg string) {
	if level.number >= MinLevel.number {
		utilities.DebugString(fmt.Sprintf("[%v] %v: %v\n", PluginName, level.name, msg))
	}
}
