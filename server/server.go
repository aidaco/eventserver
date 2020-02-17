package server

import (
	"net/http"
	"net/url"
)

const esPORT = "8080"

type EventServer interface {
	Start()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
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
	SetHeader(string, string)
	Text(int, string)
	Json(int, interface{})
	Bytes(int, []byte)
	Writer() *http.ResponseWriter
}
