package util

import (
	"io"
	"log"
	"os"
	"path/filepath"

	u "github.com/Truth1984/awadau-go"
)

type loggerStruct struct {
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
	PrintEH func(error, ...interface{})
}

type ConfigLogger struct {
	Level         int
	LogDir        string
	TraceFileName string
	DebugFileName string
	InfoFileName  string
	WarnFileName  string
	ErrorFileName string
	FatalFileName string
}

// loglevel: {10: trace, 20: debug, 30: info, 40: warn, 50: err, 60: fatal}
var Logger loggerStruct = SetLogger(ConfigLogger{Level: 40})
var Trace = func(...interface{}) {}
var Debug = func(...interface{}) {}
var Info = func(...interface{}) {}
var Warn = func(...interface{}) {}
var Error = func(...interface{}) {}
var Fatal = func(...interface{}) {}
var WarnEH = func(error, ...interface{}) {}
var ErrorEH = func(error, ...interface{}) {}
var FatalEH = func(error, ...interface{}) {}

// add errorhanle func with flag

func LogMap(param map[string]interface{}, value map[string]interface{}) string {
	res, err := u.JsonToString(u.Map("param", param, "value", value), "")
	if err != nil {
		log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile).Println("LogMap", err)
	}
	return res
}

func emptyLog() logMethod {
	return logMethod{
		Print:   func(...interface{}) {},
		Printf:  func(string, ...interface{}) {},
		Println: func(...interface{}) {},
		PrintEH: func(error, ...interface{}) {},
	}
}

func _eh(logger *log.Logger, exitAfterEH bool) func(error, ...interface{}) {
	if exitAfterEH {
		return func(err error, line ...interface{}) {
			if err != nil {
				logger.Println(append([]interface{}{err}, line...)...)
				os.Exit(1)
			}
		}
	} else {
		return func(err error, line ...interface{}) {
			if err != nil {
				logger.Println(append([]interface{}{err}, line...)...)
			}
		}
	}
}

func extractLog(logger *log.Logger, dir string, file string, exitAfterEH bool) logMethod {
	if file != "" {
		logpath := filepath.Join(dir, file)
		logFile, err := os.OpenFile(logpath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		mw := io.MultiWriter(os.Stdout, logFile)
		logger.SetOutput(mw)
	}
	return logMethod{
		Print:   logger.Print,
		Printf:  logger.Printf,
		Println: logger.Println,
		PrintEH: _eh(logger, exitAfterEH),
	}
}

// loglevel: {10: trace, 20: debug, 30: info, 40: warn, 50: err, 60: fatal}
func SetLogger(conf ConfigLogger) loggerStruct {
	var trace = emptyLog()
	var debug = emptyLog()
	var info = emptyLog()
	var warn = emptyLog()
	var err = emptyLog()
	var fatal = emptyLog()

	logLevel := conf.Level

	flags := log.Ldate | log.Ltime | log.Lshortfile
	if logLevel <= 60 {
		fatal = extractLog(log.New(os.Stderr, "FATAL: ", flags), conf.LogDir, conf.FatalFileName, true)
		Fatal = fatal.Println
		FatalEH = fatal.PrintEH
	}
	if logLevel <= 50 {
		err = extractLog(log.New(os.Stderr, "ERROR: ", flags), conf.LogDir, conf.ErrorFileName, false)
		Error = err.Println
		ErrorEH = err.PrintEH
	}
	if logLevel <= 40 {
		warn = extractLog(log.New(os.Stderr, "WARN: ", flags), conf.LogDir, conf.WarnFileName, false)
		Warn = warn.Println
		WarnEH = warn.PrintEH
	}
	if logLevel <= 30 {
		info = extractLog(log.New(os.Stdout, "INFO: ", flags), conf.LogDir, conf.InfoFileName, false)
		Info = info.Println
	}
	if logLevel <= 20 {
		debug = extractLog(log.New(os.Stdout, "DEBUG: ", flags), conf.LogDir, conf.DebugFileName, false)
		Debug = debug.Println
	}
	if logLevel <= 10 {
		trace = extractLog(log.New(os.Stdout, "TRACE: ", flags), conf.LogDir, conf.TraceFileName, false)
		Trace = trace.Println
	}

	return loggerStruct{Trace: trace, Debug: debug, Info: info, Warn: warn, Error: err, Fatal: fatal}
}
