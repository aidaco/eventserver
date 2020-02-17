package plugins

import (
	"fmt"
	"github.com/aidaco/eventserver/eventmap"
	"github.com/aidaco/eventserver/log"
	"io/ioutil"
	path "path/filepath"
	"plugin"
)

type DefaultPluginLoader struct {
	logger log.Logger
}

func (dl *DefaultPluginLoader) LoadFromDir(dirPath string) map[string]eventmap.EventHandler {
	absDirPath, err := path.Abs(dirPath)
	if err == nil {
		dl.logger.Info("Searching for plugins in dir: ", absDirPath)
		files, err := ioutil.ReadDir(absDirPath)
		if err == nil {
			handlers := make(map[string]eventmap.EventHandler)
			for _, file := range files {
				if path.Ext(file.Name()) == ".so" {
					filePath, err := path.Abs(path.Join(absDirPath, file.Name()))
					eventName, handler, err := dl.Load(filePath)
					if err == nil {
						handlers[eventName] = handler
					} else {
						dl.logger.Warn("Failed to load plugin '", filePath, "' with error:", err)
					}
				}
			}
			return handlers
		}
	}

	dl.logger.Error("Unable to open plugin directory", dirPath, "with error:", err)
	return nil
}

func (dl *DefaultPluginLoader) Load(filePath string) (string, eventmap.EventHandler, error) {
	dl.logger.Info("Attempting to load plugin:", filePath)

	absFilePath, err := path.Abs(filePath)
	if err == nil {
		dl.logger.Info("Loading plugin from file...")
		plug, err := plugin.Open(absFilePath)
		fmt.Println(err)
		if err == nil {
			dl.logger.Info(absFilePath, "loaded as plugin, reading symbols...")
			symEventName, err1 := plug.Lookup("EventName")
			symHandler, err2 := plug.Lookup("Handler")
			if err1 == nil && err2 == nil {
				eventName, ok1 := symEventName.(string)
				handler, ok2 := symHandler.(eventmap.EventHandler)
				if ok1 && ok2 {
					return eventName, handler, nil
				}
				dl.logger.Warn("Error reading symbols from plugin '", filePath, "':", ok1, ok2)
			} else {
				dl.logger.Warn("Error reading symbols from plugin '", filePath, "':", err1, err2)
			}
		}
	}

	dl.logger.Warn("Failed to load plugin '", filePath, "':", err)
	return "", nil, err
}

func NewDefaultPluginLoader(l log.Logger) *DefaultPluginLoader {
	return &DefaultPluginLoader{l}
}
