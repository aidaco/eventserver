package plugins

import (
	"github.com/aidaco/eventserver/eventmap"
)

const esPLUGINDIR = "/opt/eventserver/plugins"

type PluginLoader interface {
	Load(string) (string, *eventmap.EventHandler)
	LoadFromDir(string) map[string]*eventmap.EventHandler
}

func LoadToEventMap(loader *PluginLoader, eventMap *eventmap.EventMap) *eventmap.EventMap {

	return eventMap
}
