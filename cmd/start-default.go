package main

import (
	"github.com/aidaco/eventserver/eventmap"
	"github.com/aidaco/eventserver/log"
	"github.com/aidaco/eventserver/plugins"
	"github.com/aidaco/eventserver/server"
)

func DefaultLoader() (s *server.EventServer, em *eventmap.EventMap, pl *plugins.Loader, l *log.Logger) {
	log := log.NewDefaultLogger()

	em = eventmap.NewDefaultEventMap(log)
	pl = plugins.NewDefaultLoader(log)
	pl.LoadToMap(em)

	s := server.NewDefaultEventServer(eventmap)
}

func main() {
	server, eventmap, pluginloader, log = DefaultLoader()
	server.Serve()
}
