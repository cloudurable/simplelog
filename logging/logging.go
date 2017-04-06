package logging

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type LogLevel byte

const (
	EMERGENCY = iota
	ALERT
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

type Logger interface {
	Printf(format string, args ...interface{})
	Emergencyf(format string, args ...interface{})
	Alertf(format string, args ...interface{})
	Criticalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Noticef(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})

	PrintError(message string, err error)
	EmergencyError(message string, err error)
	AlertError(message string, err error)
	CriticalError(message string, err error)
	ErrorError(message string, err error)
	WarnError(message string, err error)
	NoticeError(message string, err error)
	InfoError(message string, err error)
	DebugError(message string, err error)

	Println(args ...interface{})
	Emergency(args ...interface{})
	Alert(args ...interface{})
	Critical(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Notice(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})

    Name() string
}

func NewSimpleLogger(name string) Logger {
	return NewLogger(name, INFO, true, os.Stderr, os.Stderr, os.Stderr, os.Stderr, os.Stderr, os.Stdout, os.Stdout, ioutil.Discard)
}

func NewSimpleDebugLogger(name string) Logger {
	return NewLogger(name, DEBUG, true, os.Stderr, os.Stderr, os.Stderr, os.Stderr, os.Stderr, os.Stdout, os.Stdout, os.Stdout)
}

func NewLogger(name string, logLevel LogLevel, panicOnEmergency bool, emergency io.Writer, alert io.Writer, critical io.Writer,
	error io.Writer, warning io.Writer, notice io.Writer,
	info io.Writer, debug io.Writer) *BasicLogger {

	logger := BasicLogger{}
    logger.name = name
	logger.logLevel = logLevel
	logger.panicOnEmergency = panicOnEmergency
	logger.alert = log.New(alert,
		"ALERT    : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.emergency = log.New(emergency,
		"EMERGENCY: ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.critical = log.New(critical,
		"CRITICAL : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.notice = log.New(notice,
		"NOTICE   : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.debug = log.New(debug,
		"DEBUG    : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.info = log.New(info,
		"INFO     : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.warning = log.New(warning,
		"WARN     : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	logger.error = log.New(error,
		"ERROR    : ["+name+"] - ",
		log.Ldate|log.Ltime|log.Lshortfile)

	return &logger

}

func GetSimpleLogger(flag string, name string) Logger {
	value, wasSet := os.LookupEnv(flag)
	if wasSet && (value == "TRUE" || value == "true") {
		return NewSimpleDebugLogger(name)
	} else {
		return NewSimpleLogger(name)
	}
}
