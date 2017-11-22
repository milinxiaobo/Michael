package logger

import (
	"io"
	"log"
)

var (
	// Trace blabla
	Trace *log.Logger
	// Info blabla
	Info *log.Logger
	// Warning blabla
	Warning *log.Logger
	// Error blabla
	Error *log.Logger
)

// InitLogger blabla
func InitLogger(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Trace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
