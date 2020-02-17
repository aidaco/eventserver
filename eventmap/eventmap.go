package eventmap

import (
	"github.com/2sdat/eventserver/log"
	"github.com/2sdat/eventserver/server"
)

type Event struct {
	Name string
	Req *server.Request
	Res *server.Response
}

type EventHandler func(*Event) error

type EventMap interface {
	RegisterHandler(string, *EventHandler)
	Handle(Event) string
}