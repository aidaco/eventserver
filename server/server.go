package server

import (
	"github.com/aidaco/eventserver/eventmap"
	"net/http"
	"net/url"
)

const esPORT = "8080"

type EventServer interface {
	Start()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	MakeResponse(w http.ResponseWriter) *Response
	MakeRequest(r *http.Request) *Request
}

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
	Setheader(string, string)
	Text(int, string)
	Json(int, interface{})
	Bytes(int, []byte)
	Writer() *http.ResponseWriter
}

func DefaultServer(em *eventmap.EventMap) EventServer {
	es := DefaultEventServer{em}
}
