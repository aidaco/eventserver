package eventmap

type DefaultEventMap struct {
	logger *log.logger
	handlers *map[string]*EventHandler
}

func (em *DefaultEventMap) RegisterHandler(eventname string, handler *EventHandler) {
	if _, exists = em.handlers[eventname]; exists {
		em.logger.Warn("Handler for event '", eventname, "' already exists, replacing.")

	em.handlers[eventname] = handler
	em.logger.Info("Registered Handler:", eventname)
}

func (em *DefaultEventMap) RegisterHandlers(handlers *map[string]*EventHandler) {
	for e, h := range handlers {
		em.RegisterHandler(e, h)
	}
}

func (em *DefaultEventMap) Handle(event Event) (res []byte, err error) {
	if _, exists = em.handlers[event.name]; exists {
		res, err = (*em.handlers[event.name])(*event)
		if err {
			em.logger.Warn("Handler '", event.name, "' failed:", err)
		} else {
			em.logger.Info("Handled '", event.name, "' with response:", res)
		}
	} else {
		em.logger.Warn("Handler '", event.name, "' not registered")
		err = errors.New("Handler '", event.name, "' not registered")
	}
}

func Default(l *log.Logger) (em *EventMap){
	em = &DefaultEventMap{ l, make(*map[string]*EventHandler) }
}