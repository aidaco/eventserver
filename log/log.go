package log

import (
	"io"
	"log"
	"os"
	path "path/filepath"
)

const esLOGDIR string = "/opt/eventserver/logs"

type Logger interface{
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}



