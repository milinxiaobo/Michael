package logger

import (
	"log"
	"os"
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
func InitLogger() {
	fTrace, err := os.OpenFile("pcapagent_trace.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		panic(err)
	}
	fInfo, err := os.OpenFile("pcapagent_info.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		panic(err)
	}
	Trace = log.New(fTrace, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(fInfo, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(fTrace, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(fTrace, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
