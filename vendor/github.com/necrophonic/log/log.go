// Package log defines a generic log implemetation with differing granularity levels.
//
// The logger is built on top of the standard golang log package.
package log

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Level setting constants
//   LevelTrace > LevelDebug > LevelWarn > LevelInfo > LevelError > LevelNone
const (
	LevelNone  = 0
	LevelError = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelDebug = 4
	LevelTrace = 5
)

//Logger is the default interface all logs must implement in this library
type Logger interface {
	Fatal(v ...interface{})
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(fv ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})

	Fatalf(format string, v ...interface{})
	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(fformat string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

var (
	level = LevelInfo
)

// DefaultLogger defines a simple logging package
type DefaultLogger struct {
	Level int
}

// Init initialises the logger. This must be called at least once before and of
// the functions can be called.
func Init(l int) error {
	if l < LevelNone || l > LevelTrace {
		return errors.New("unable to init logger: invalid level")
	}
	level = l
	return nil
}

// InitFromString initialises the logger with Init(). Takes a string rather than a numerical levvel
func InitFromString(l string) error {
	switch strings.ToLower(l) {
	case "error":
		Init(LevelError)
	case "info":
		Init(LevelInfo)
	case "warn":
		Init(LevelWarn)
	case "debug":
		Init(LevelDebug)
	case "trace":
		Init(LevelTrace)
	case "none":
		Init(LevelNone)
	default:
		return fmt.Errorf("unable to init logger: unknown level '%s'", l)
	}
	return nil
}

// Fatalf provides fatal level logging in the manner of fmt.Printf.
// Being fatal it will log, and then it will exit the current process.
func Fatalf(format string, v ...interface{}) {
	writeLog("[ERROR]", LevelError, format, v...)
	os.Exit(1)
}

// Fatal provides fatal level logging in the manner of fmt.Print.
// Being fatal it will log, and then it will exit the current process.
func Fatal(v ...interface{}) {
	writeLog("[ERROR]", LevelError, fmt.Sprint(v))
	os.Exit(1)
}

// Errorf provides error level logging in the manner of fmt.Printf
func Errorf(format string, v ...interface{}) {
	writeLog("[ERROR]", LevelError, format, v...)
}

// Error provides error level logging in the manner of fmt.Print
func Error(v ...interface{}) {
	Errorf(fmt.Sprint(v))
}

// Debugf provides debug level logging in the manner of fmt.Printf
func Debugf(format string, v ...interface{}) {
	writeLog("[DEBUG]", LevelDebug, format, v...)
}

// Debug provides debug level logging in the manner of fmt.Print
func Debug(v ...interface{}) {
	Debugf(fmt.Sprint(v))
}

// Warnf provides warning level logging in the manner of fmt.Printf
func Warnf(format string, v ...interface{}) {
	writeLog("[WARN]", LevelWarn, format, v...)
}

// Warn provides warning level logging in the manner of fmt.Print
func Warn(v ...interface{}) {
	Warnf(fmt.Sprint(v))
}

// Infof provides info level logging in the manner of fmt.Printf
func Infof(format string, v ...interface{}) {
	writeLog("[INFO]", LevelInfo, format, v...)
}

// Info provides info level logging the manner of fmt.Print
func Info(v ...interface{}) {
	Infof(fmt.Sprint(v))
}

// Tracef provides trace level logging in the manner of fmt.Printf
func Tracef(format string, v ...interface{}) {
	writeLog("[TRACE]", LevelTrace, format, v...)
}

// Trace provides trace level logging in the manner of fmt.Print
func Trace(v ...interface{}) {
	Tracef(fmt.Sprint(v))
}

func writeLog(prefix string, l int, format string, v ...interface{}) {
	if level >= l {
		if v != nil {
			log.Printf(prefix+" "+format, v...)
		} else {
			log.Println(prefix + " " + format)
		}
	}
}
