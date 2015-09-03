# goplane
[![GoDoc](https://godoc.org/github.com/abieberbach/goplane?status.svg)](https://godoc.org/github.com/abieberbach/goplane)

Implementierung des [X-Plane SDKs](http://www.xsquawkbox.net/xpsdk/mediawiki/Main_Page) auf der Basis von [GO](http://www.golang.org)

## Voraussetzungen

1. [GO Version 1.5](http://www.golang.org) muss installiert sein
2. Das [X-Plane SDK](http://www.xsquawkbox.net/xpsdk/mediawiki/Main_Page) (min. Version 2.1.3) muss heruntergeladen und in einem Ordner entpackt sein.

## Beispiel-Plugin
```
package main
import (
	"github.com/abieberbach/goplane/extra"
	"github.com/abieberbach/goplane/extra/logging"
)

func main() {
}

func init() {
	plugin := extra.NewPlugin("TestPlugin", "com.github.abieberbach.goplane.TestPlugin", "TestPlugin")
	plugin.SetPluginStateCallback(onPluginStateChanged)
	logging.MinLevel = logging.Info_Level
}

func onPluginStateChanged(state extra.PluginState, plugin *extra.XPlanePlugin) {
	switch state {
	case extra.PluginStart:
		onPluginStart()
	case extra.PluginEnable:
		onPluginEnable()
	}
}

func onPluginStart() {
	logging.Info("Plugin start")
}


func onPluginEnable() {
	logging.Info("Plugin enable")
}
```

## Plugin kompilieren

Damit das Plugin kompiliert werden kann, muss dem System mitgeteilt werden, wo sich die Include-Dateien des SDKs befinden.
Dazu muss die Umgebungsvariable "CGO_CFLAGS" auf den Pfad des SDKs gesetzt werden.

```
Linux: export CGO_CFLAGS='-I <Pfad zum SDK>/CHeaders'
```

Im Anschluss kann das Plugin mit dem go Tool gebaut werden:

```
Linux: go build -o lin.xpl -buildmode=c-shared  TestPlugin.go
```

## Unterstütze Plattformen

- Linux: x64
- Mac: nicht getestet
- Windows: nicht unterstützt (siehe offenen Issue bei GO https://github.com/golang/go/issues/11058)

