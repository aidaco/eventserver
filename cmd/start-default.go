package main

import (
	"github.com/2sdat/eventserver/server"
	"github.com/2sdat/eventserver/eventmap"
	"github.com/2sdat/eventserver/plugins"
	"github.com/2sdat/eventserver/log"
)

func DefaultLoader() (server *server.EventServer, eventmap *eventmap.EventMap, pluginloader *plugins.Loader(), logger *log.Logger) {
	log := log.NewDefaultLogger()

	eventmap := eventmap.NewDefaultEventMap(log)
	pluginloader := plugins.NewDefaultLoader(log)
	pluginloader.LoadToMap(eventmap)

	server := server.NewDefaultEventServer(eventmap)
}

func main() {
	server, eventmap, pluginloader, log = DefaultLoader()
	server.Serve()
}