package util

import (
	"log"
	"os"

	u "github.com/Truth1984/awadau-go"
)

type logger struct {
	Trace logMethod
	Debug logMethod
	Info  logMethod
	Warn  logMethod
	Error logMethod
	Fatal logMethod
}

type logMethod struct {
	Print   func(...interface{})
	Printf  func(string, ...interface{})
	Println func(...interface{})
}

func LogMap(param map[string]interface{}, value map[string]interface{}) string {
	res, err := u.JsonToString(u.Map("param", param, "value", value), "")
	if err != nil {
		Logger().Error.Println("LogMap", err)
	}
	return res
}

func EHWarn(err error, line ...interface{}) {
	if err != nil {
		Logger().Warn.Println(append([]interface{}{err}, line...)...)

	}
}

func EHError(err error, line ...interface{}) {
	if err != nil {
		Logger().Error.Println(append([]interface{}{err}, line...)...)
	}
}

func EHFatal(err error, line ...interface{}) {
	if err != nil {
		Logger().Fatal.Println(append([]interface{}{err}, line...)...)
		os.Exit(1)
	}
}

func emptyLog() logMethod {
	return logMethod{
		Print:   func(...interface{}) {},
		Printf:  func(string, ...interface{}) {},
		Println: func(...interface{}) {},
	}
}

func extractLog(logger *log.Logger) logMethod {
	return logMethod{
		Print:   logger.Print,
		Printf:  logger.Printf,
		Println: logger.Println,
	}
}

// loglevel: {10: trace, 20: debug, 30: info, 40: warn, 50: err, 60: fatal}
func SetLogger(logLevel int) logger {
	var trace = emptyLog()
	var debug = emptyLog()
	var info = emptyLog()
	var warn = emptyLog()
	var err = emptyLog()
	var fatal = emptyLog()

	flags := log.Ldate | log.Ltime | log.Lshortfile
	if logLevel <= 60 {
		fatal = extractLog(log.New(os.Stderr, "FATAL: ", flags))
	}
	if logLevel <= 50 {
		err = extractLog(log.New(os.Stderr, "ERROR: ", flags))
	}
	if logLevel <= 40 {
		warn = extractLog(log.New(os.Stderr, "WARN: ", flags))
	}
	if logLevel <= 30 {
		info = extractLog(log.New(os.Stdout, "INFO: ", flags))
	}
	if logLevel <= 20 {
		debug = extractLog(log.New(os.Stdout, "DEBUG: ", flags))
	}
	if logLevel <= 10 {
		trace = extractLog(log.New(os.Stdout, "TRACE: ", flags))
	}

	return logger{Trace: trace, Debug: debug, Info: info, Warn: warn, Error: err, Fatal: fatal}
}

// loglevel: {10: trace, 20: debug, 30: info, 40: warn, 50: err, 60: fatal}
func Logger() logger {
	return SingletonGet("loggerInstance").(logger)
}

// Logger Trace Println
func LTP(line ...interface{}) {
	Logger().Trace.Println(line...)
}

// Logger Debug Println
func LDP(line ...interface{}) {
	Logger().Debug.Println(line...)
}

// Logger Info Println
func LIP(line ...interface{}) {
	Logger().Info.Println(line...)
}

//	Logger Warn Println
func LWP(line ...interface{}) {
	Logger().Warn.Println(line...)
}
