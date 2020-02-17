package main

import (
	"github.com/aidaco/eventserver/eventmap"
	"github.com/aidaco/eventserver/log"
	"github.com/aidaco/eventserver/plugins"
	"github.com/aidaco/eventserver/server"
	"net/http"
)

type Modules struct {
	Log          log.Logger
	EventMap     eventmap.EventMap
	PluginLoader plugins.PluginLoader
	Server       server.EventServer
}

func DefaultLoader() *Modules {
	l := log.NewDefaultLogger()
	em := eventmap.NewDefaultEventMap(l)
	pl := plugins.NewDefaultPluginLoader(l)
	plugins.LoadToEventMap(pl, em)
	s := server.NewDefaultEventServer(l, em)
	return &Modules{l, em, pl, s}
}

func testHandler(event eventmap.Event) error {
	event.Res.Text(http.StatusOK, "Hello, World!")
	return nil
}

func main() {
	modules := DefaultLoader()
	modules.EventMap.RegisterHandler("test", testHandler)
	modules.Server.Start()
}
