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
	}
}

func extractLog(logger *log.Logger, dir string, file string) logMethod {
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
		fatal = extractLog(log.New(os.Stderr, "FATAL: ", flags), conf.LogDir, conf.FatalFileName)
	}
	if logLevel <= 50 {
		err = extractLog(log.New(os.Stderr, "ERROR: ", flags), conf.LogDir, conf.ErrorFileName)
	}
	if logLevel <= 40 {
		warn = extractLog(log.New(os.Stderr, "WARN: ", flags), conf.LogDir, conf.WarnFileName)
	}
	if logLevel <= 30 {
		info = extractLog(log.New(os.Stdout, "INFO: ", flags), conf.LogDir, conf.InfoFileName)
	}
	if logLevel <= 20 {
		debug = extractLog(log.New(os.Stdout, "DEBUG: ", flags), conf.LogDir, conf.DebugFileName)
	}
	if logLevel <= 10 {
		trace = extractLog(log.New(os.Stdout, "TRACE: ", flags), conf.LogDir, conf.TraceFileName)
	}

	return loggerStruct{Trace: trace, Debug: debug, Info: info, Warn: warn, Error: err, Fatal: fatal}
}

func Trace(line ...interface{}) {
	Logger.Trace.Println(line...)
}

func Debug(line ...interface{}) {
	Logger.Debug.Println(line...)
}

func Info(line ...interface{}) {
	Logger.Info.Println(line...)
}

func Warn(line ...interface{}) {
	Logger.Warn.Println(line...)
}

func WarnEH(err error, line ...interface{}) {
	if err != nil {
		Logger.Warn.Println(append([]interface{}{err}, line...)...)
	}
}

func Error(line ...interface{}) {
	Logger.Error.Println(line...)
}

func ErrorEH(err error, line ...interface{}) {
	if err != nil {
		Logger.Error.Println(append([]interface{}{err}, line...)...)
	}
}

func Fatal(line ...interface{}) {
	Logger.Fatal.Println(line...)
	os.Exit(1)
}

func FatalEH(err error, line ...interface{}) {
	if err != nil {
		Logger.Fatal.Println(append([]interface{}{err}, line...)...)
		os.Exit(1)
	}
}
