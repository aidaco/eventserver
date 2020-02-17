package plugins

import (
	"github.com/aidaco/eventserver/eventmap"
	"github.com/aidaco/eventserver/log"
	"io/ioutil"
	"os"
	path "path/filepath"
	"plugin"
)

type DefaultPluginLoader struct {
	logger *log.Logger
}

func (dl *DefaultPluginLoader) LoadToMap(em *eventmap.EventMap) {
	dirpath := os.Getenv("esPLUGINDIR")
	if dirpath == "" {
		dirpath = esPLUGINDIR
	}

	dirpath, err := path.Abs(dirpath)
	(*dl.logger).Info("Searching for plugins in dir: ", dirpath)

	direntries, err := ioutil.ReadDir(dirpath)
	if err != nil {
		(*dl.logger).Error("Unable to open plugin directory:", err)
		os.Exit(1)
	}

	for _, f := range direntries {
		if path.Ext(f.Name()) == ".so" {
			pathtofile, err := path.Abs(path.Join(dirpath, f.Name()))
			eventname, handler, err := dl.Load(pathtofile)
			if err == nil {
				(*em).RegisterHandler(eventname, handler)
			}
		}
	}
}

func (dl *DefaultPluginLoader) Load(file string) (eventname string, handler *eventmap.EventHandler, err error) {
	dl.logger.Info("Attempting to load plugin:", f)

	if plug, err := plugin.Open(dirpath + f); err == nil {
		symEventName, err1 := plug.Lookup("EventName")
		symHandler, err2 := plug.Lookup("Handler")
		if err1 == nil && err2 == nil {
			eventname, err1 = symEventName.(string)
			handler, err2 = symHandler.(*eventmap.EventHandler)
			if err1 == nil && err2 == nil {
				err == nil
			} else {
				dl.logger.Warn("Error reading symbols from plugin '", f, "':", err1, err2)
			}
		} else {
			dl.logger.Warn("Error reading symbols from plugin '", f, "':", err1, err2)
		}
		if err1 == nil {
			return err2
		}
	} else {
		dl.logger.Warn("Failed to load plugin '", f, "':", err)
	}
}

func Default(l *Logger) *PluginLoader {
	return &DefaultPluginLoader{l, nil}
}
