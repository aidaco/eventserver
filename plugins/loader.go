package plugins

import (
	"github.com/aidaco/eventserver/eventmap"
)

const esPLUGINDIR = "/opt/eventserver/plugins"

type PluginLoader interface {
	LoadToMap(*eventmap.eventmap) *map[string]*eventmap.EventHandler
	Load(string) (string, *eventmap.EventHandler)
}
