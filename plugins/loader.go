package plugins

import (
	"github.com/aidaco/eventserver/eventmap"
	"os"
)

const esPLUGINDIR = "/opt/eventserver/plugins"

type PluginLoader interface {
	Load(string) (string, eventmap.EventHandler, error)
	LoadFromDir(string) map[string]eventmap.EventHandler
}

func LoadToEventMap(loader PluginLoader, eventMap eventmap.EventMap) eventmap.EventMap {
	dirPath := os.Getenv("esPLUGINDIR")
	if dirPath == "" {
		dirPath = esPLUGINDIR
	}

	handlers := loader.LoadFromDir(dirPath)
	for name, handler := range handlers {
		eventMap.RegisterHandler(name, handler)
	}

	return eventMap
}
