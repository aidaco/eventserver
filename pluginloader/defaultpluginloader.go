package pluginloader

import (
	"github.com/aidaco/eventserver/log"
	"io/ioutil"
	path "path/filepath"
	"plugin"
)

type DefaultPluginLoader struct {
	logger log.Logger
}

func (dl *DefaultPluginLoader) LoadFromDir(dirPath string) []*Plugin {
	absDirPath, err := path.Abs(dirPath)
	if err == nil {
		dl.logger.Info("Searching for plugin in dir: ", absDirPath)
		files, err := ioutil.ReadDir(absDirPath)
		if err == nil {
			var plugins []*Plugin
			for _, file := range files {
				if path.Ext(file.Name()) == ".so" {
					filePath, err := path.Abs(path.Join(absDirPath, file.Name()))
					loaded, err := dl.Load(filePath)
					if err == nil {
						plugins = append(plugins, loaded)
					} else {
						dl.logger.Warn("Failed to load plugin '", filePath, "' with error:", err)
					}
				}
			}
			return plugins
		}
	}

	dl.logger.Error("Unable to open plugin directory", dirPath, "with error:", err)
	return nil
}

func (dl *DefaultPluginLoader) Load(filePath string) (*Plugin, error) {
	dl.logger.Info("Attempting to load plugin:", filePath)

	absFilePath, err := path.Abs(filePath)
	if err != nil {
		dl.logger.Warn("Failed to generate absolute path to file '", filePath, "':", err)
		return nil, err
	}

	dl.logger.Info("Loading from file...")
	plug, err := plugin.Open(absFilePath)
	if err != nil {
		dl.logger.Warn("Failed to open file as plugin:", filePath)
		return nil, err
	}

	dl.logger.Info(absFilePath, "Reading symbols...")
	symbol, err := plug.Lookup("EventHandler")
	if err != nil {
		dl.logger.Warn("Failed to find symbol 'Eventhandler' in loaded plugin:", err)
		return nil, err
	}

	dl.logger.Info("Verifying plugin...")
	loaded, ok := symbol.(*Plugin)
	if !ok {
		dl.logger.Warn("Failed to assert type 'Plugin' on loaded symbols: '", filePath, "':", ok)
		return nil, err
	}

	dl.logger.Info("Successfully loaded handler for event: '", loaded.EventName, "'")
	return loaded, nil
}

func NewDefaultPluginLoader(l log.Logger) *DefaultPluginLoader {
	return &DefaultPluginLoader{l}
}
