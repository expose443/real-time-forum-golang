package logger

import (
	"log"
	"os"
)

const (
	green  string = "\033[32m"
	yellow string = "\033[33m"
	red    string = "\033[31m"
	reset  string = "\033[0m"
)

type logLevel struct {
	Info    *log.Logger
	Error   *log.Logger
	Warning *log.Logger
}

func NewLog() *logLevel {
	// logFile, err := os.OpenFile("/home/student/real-time-forum-golang/pkg/logger/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	// if os.IsNotExist(err) {
	// 	logFile, err = os.Create("/home/student/real-time-forum-golang/pkg/logger/log.log")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	infoLog := log.New(os.Stdout, green+"[INFO]   \t"+reset, log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, red+"[ERROR]  \t"+reset, log.Ldate|log.Ltime|log.Lshortfile)
	warningLog := log.New(os.Stdout, yellow+"[WARNING]\t"+reset, log.Ldate|log.Ltime)

	newLog := logLevel{
		Info:    infoLog,
		Error:   errorLog,
		Warning: warningLog,
	}
	return &newLog
}
