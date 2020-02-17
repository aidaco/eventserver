package server

import (
	"fmt"
	"github.com/aidaco/eventserver/eventmap"
	"github.com/aidaco/eventserver/log"
	"net/http"
	"os"
)

type DefaultEventServer struct {
	logger   log.Logger
	eventMap eventmap.EventMap
}

func (es *DefaultEventServer) Start() {
	port := os.Getenv("esPORT")
	if port == "" {
		port = esPORT
	}
	es.logger.Info("Starting server...")
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), es)
	es.logger.Error("Server no longer running with error:", err)
	os.Exit(1)
}

func (es *DefaultEventServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	eventName := url[len("/"):]

	es.logger.Info("Received Request:", r.RemoteAddr, r.Method, r.URL)

	event := eventmap.Event{Name: eventName, Req: MakeDefaultRequest(r), Res: MakeDefaultResponse(w)}

	if err := es.eventMap.Handle(event); err == nil {
		es.logger.Info("Processed Request:", r.RemoteAddr)
	} else {
		es.logger.Warn("Request on Invalid Event:", eventName)
		w.WriteHeader(http.StatusNotFound)
	}
}

func NewDefaultEventServer(logger log.Logger, eventMap eventmap.EventMap) *DefaultEventServer {
	return &DefaultEventServer{logger: logger, eventMap: eventMap}
}

func MakeDefaultResponse(w http.ResponseWriter) *DefaultResponse {
	return &DefaultResponse{w}
}

func MakeDefaultRequest(r *http.Request) *DefaultRequest {
	return &DefaultRequest{r, nil}
}
