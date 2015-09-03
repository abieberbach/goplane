package logging

import (
	"fmt"
	"github.com/abieberbach/goplane/xplm/utilities"
)


type Level struct {
	number byte
	name   string
}

var (
	Debug_Level = Level{1, "DEBUG"}
	Info_Level = Level{2, "INFO"}
	Warning_Level = Level{3, "WARNING"}
	Error_Level = Level{4, "ERROR"}

	MinLevel = Info_Level
	PluginName = "<unknown>"
)

func Debug(msg string) {
	writeMessage(Debug_Level, msg)
}

func Debugf(format string, a... interface{}) {
	if Debug_Level.number >= MinLevel.number {
		Debug(fmt.Sprintf(format, a...))
	}
}

func Info(msg string) {
	writeMessage(Info_Level, msg)
}

func Infof(format string, a... interface{}) {
	if Info_Level.number >= MinLevel.number {
		Info(fmt.Sprintf(format, a...))
	}
}

func Warning(msg string) {
	writeMessage(Warning_Level, msg)
}

func Warningf(format string, a... interface{}) {
	if Warning_Level.number >= MinLevel.number {
		Warning(fmt.Sprintf(format, a...))
	}
}


func Error(msg string) {
	writeMessage(Error_Level, msg)
}

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
