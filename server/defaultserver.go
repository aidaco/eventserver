package server

import (
	"github.com/aidaco/eventserver/eventmap"
	"log"
	"net/http"
	"os"
)

type DefaultEventServer struct {
	logger   *log.Logger
	eventmap *eventmap.EventMap
}

func (es *DefaultEventServer) Start() {
	if port := os.Getenv("esPORT"); port == "" {
		port = esPORT
	}
	err := http.ListenAndServe(fmt.SprintF(":%v", port), es)
	es.logger.Error("Server has stopped.")
	log.Fatal(err)
	os.Exit(1)
}

func (es *DefaultEventServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	eventname := r.URL.Path[len("/"):]

	es.logger.Info("Received Request:", r.RemoteAddr, r.Method, r.URL, body)

	event := eventmap.Event{eventname, es.MakeRequest(r), es.MakeResponse(w)}

	if ok := es.eventmap.Handle(event); ok {
		es.logger.Info("Processed Request:", r.RemoteAddr)
	} else {
		es.logger.Warn("Request on Invalid Event:", eventname)
		w.WriteHeader(http.StatusNotFound)
	}
}

func (es *DefaultEventServer) MakeResponse(w http.ResponseWriter) *Response {
	res := &DefaultResponse{w}
	return res
}

func (es *DefaultEventServer) MakeRequest(r *http.Request) *Request {
	req := &DefaultRequest{r, nil}
	return req
}
