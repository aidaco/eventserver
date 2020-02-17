package log

const esLOGDIR string = "/opt/eventserver/logs"

type Logger interface {
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}
