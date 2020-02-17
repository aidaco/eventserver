package plugins

import (
	"os"
	"io/ioutil"
	path "path/filepath"
	"plugin"
	"github.com/2sdat/eventserver/log"
	"github.com/2sdat/eventserver/eventmap"
)

const esPLUGINDIR = "/opt/eventserver/plugins"

type PluginLoader interface {
	LoadToMap(*eventmap.eventmap) *map[string]*eventmap.EventHandler
	Load(string) (string, *eventmap.EventHandler)
}
