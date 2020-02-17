package eventmap

import (
	"net/http"
	"net/url"
)

type Request interface {
	Body() []byte
	URL() *url.URL
	Path() string
	Header() http.Header
	Addr() string
	Method() string
	Request() *http.Request
}

type Response interface {
	SetHeader(string, string)
	Text(int, string)
	Json(int, interface{})
	Bytes(int, []byte)
	Writer() *http.ResponseWriter
}

type Event struct {
	Name string
	Req  Request
	Res  Response
}

type EventHandler func(Event) error

type EventMap interface {
	RegisterHandler(string, EventHandler)
	Handle(Event) error
}
