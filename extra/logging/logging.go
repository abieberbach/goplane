package logging

import (
	"fmt"
	"github.com/abieberbach/goplane/xplm/utilities"
	"strings"
)

//Informationen über ein Loglevel
type Level struct {
	number byte   //Nummer des Loglevels
	name   string //Name des Loglevels
}

var (
//Loglevel für Tracemeldungen
	Trace_Level = Level{1, "TRACE"}
//Loglevel für Debugmeldungen
	Debug_Level = Level{2, "DEBUG"}
//Loglevel für Infomeldungen
	Info_Level = Level{3, "INFO"}
//Loglevel für Warnungen
	Warning_Level = Level{4, "WARNING"}
//Loglevel für Fehler
	Error_Level = Level{5, "ERROR"}

//Level ab dem die Meldungen ausgegeben werden
	MinLevel = Info_Level
//aktueller Pluginname
	PluginName = "<unknown>"
)

//Ermittelt aus einem String das entsprechende Loglevel. Mögliche Werte sind: TRACE, DEBUG, INFO, WARNING, ERROR.
//Wird ein anderer String verwendet, dann liefert die Methode das Info-Level
func GetLevelFromString(level string) Level {
	switch strings.ToUpper(level) {
	case "TRACE":
		return Trace_Level
	case "DEBUG":
		return Debug_Level
	case "INFO":
		return Info_Level
	case "WARNING":
		return Warning_Level
	case "ERROR":
		return Error_Level
	default:
		return Info_Level
	}
}

//Schreibt eine Tracemeldung in die Logdatei
func Trace(msg string) {
	writeMessage(Trace_Level, msg)
}

//Schreibt eine formatierte Tracemeldung in die Logdatei
func Tracef(format string, a... interface{}) {
	if Trace_Level.number >= MinLevel.number {
		Trace(fmt.Sprintf(format, a...))
	}
}

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
