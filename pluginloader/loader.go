package pluginloader

import (
	"github.com/aidaco/eventserver/eventmap"
	"os"
)

const esPLUGINDIR = "plugins"

type Plugin struct {
	EventName string
	Handler   func(eventmap.Event) error
}

type PluginLoader interface {
	Load(string) (*Plugin, error)
	LoadFromDir(string) []*Plugin
}

func ToEventMap(loader PluginLoader, eventMap eventmap.EventMap) eventmap.EventMap {
	dirPath := os.Getenv("esPLUGINDIR")
	if dirPath == "" {
		dirPath = esPLUGINDIR
	}

	plugins := loader.LoadFromDir(dirPath)
	for _, p := range plugins {
		eventMap.RegisterHandler(p.EventName, p.Handler)
	}

	return eventMap
}
