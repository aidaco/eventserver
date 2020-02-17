package eventmap

import (
	"errors"
	"fmt"
	"github.com/aidaco/eventserver/log"
)

type DefaultEventMap struct {
	logger   *log.Logger
	handlers map[string]*EventHandler
}

func (em *DefaultEventMap) RegisterHandler(eventname string, handler *EventHandler) {
	if _, exists := em.handlers[eventname]; exists {
		(*em.logger).Warn("Handler for event '", eventname, "' already exists, replacing.")
	}
	em.handlers[eventname] = handler
	(*em.logger).Info("Registered Handler:", eventname)
}

func (em *DefaultEventMap) Handle(event Event) error {
	var err error = nil
	if _, exists := em.handlers[event.Name]; exists {
		if err = (*em.handlers[event.Name])(&event); err == nil {
			(*em.logger).Info("Handled '", event.Name)
			return err
		} else {
			(*em.logger).Warn("Handler '", event.Name, "' failed:")
		}
	} else {
		(*em.logger).Warn("Handler '", event.Name, "' not registered")
		err = errors.New(fmt.Sprintf("Handler '%v' not registered", event.Name))
	}

	return err
}

func NewDefaultEventMap(l *log.Logger) *DefaultEventMap {
	return &DefaultEventMap{l, make(map[string]*EventHandler)}
}
