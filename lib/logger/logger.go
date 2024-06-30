package logger

import (
	"log"
	"os"
	"sync"
)

type LogLevel int

const (
	ERROR LogLevel = iota
	WARN
	INFO
	DEBUG
)

var levelStrings = map[LogLevel]string{
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
}

type Logger struct {
	level LogLevel
	mu    sync.Mutex
}

var logger *Logger

func init() {
	logger = &Logger{}
	logger.SetLevelFromEnv()
	log.SetFlags(log.LstdFlags)
}

// SetLevel sets the logging level
func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// SetLevelFromEnv sets the logging level from the environment variable LOG_LEVEL
func (l *Logger) SetLevelFromEnv() {
	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "ERROR":
		l.SetLevel(ERROR)
	case "WARN":
		l.SetLevel(WARN)
	case "INFO":
		l.SetLevel(INFO)
	case "DEBUG":
		l.SetLevel(DEBUG)
	default:
		l.SetLevel(INFO)
	}
}

func logMessage(level LogLevel, v ...interface{}) {
	if logger.level >= level {
		log.Println(append([]interface{}{levelStrings[level]}, v...)...)
	}
}

func logFormattedMessage(level LogLevel, format string, v ...interface{}) {
	if logger.level >= level {
		log.Printf(levelStrings[level]+": "+format, v...)
	}
}

// Error logs error messages
func Error(v ...interface{}) {
	logMessage(ERROR, v...)
}

// Errorf logs formatted error messages
func Errorf(format string, v ...interface{}) {
	logFormattedMessage(ERROR, format, v...)
}

// Warn logs warning messages
func Warn(v ...interface{}) {
	logMessage(WARN, v...)
}

// Warnf logs formatted warning messages
func Warnf(format string, v ...interface{}) {
	logFormattedMessage(WARN, format, v...)
}

// Info logs informational messages
func Info(v ...interface{}) {
	logMessage(INFO, v...)
}

// Infof logs formatted informational messages
func Infof(format string, v ...interface{}) {
	logFormattedMessage(INFO, format, v...)
}

// Debug logs debug messages
func Debug(v ...interface{}) {
	logMessage(DEBUG, v...)
}

// Debugf logs formatted debug messages
func Debugf(format string, v ...interface{}) {
	logFormattedMessage(DEBUG, format, v...)
}
