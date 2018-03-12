package utils

import (
	"io"
	"log"
	"os"
)

const (
	TRACE_FILE_PATH   = "trace.log"
	INFO_FILE_PATH    = "info.log"
	WARNING_FILE_PATH = "warning.log"
	ERROR_FILE_PATH   = "error.log"
)

var (
	TRACE   *log.Logger
	INFO    *log.Logger
	WARNING *log.Logger
	ERROR   *log.Logger
)

func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	TRACE = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	INFO = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	WARNING = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ERROR = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}

// InitLoggers To initialize loggers
func InitLoggers() {
	traceLogFile, err1 := os.Create(TRACE_FILE_PATH)
	infoLogFile, err2 := os.Create(INFO_FILE_PATH)
	warningLogFile, err3 := os.Create(WARNING_FILE_PATH)
	errorLogFile, err4 := os.Create(ERROR_FILE_PATH)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.Fatalln("Fail to open log file")
	}
	initLog(traceLogFile, infoLogFile, warningLogFile, errorLogFile)
	INFO.Println("Loggers have been set up successfully.")
}

// func RedisConnect() redis.Conn {

// }
