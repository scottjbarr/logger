package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	timeFormat = "2006-01-02T15:04:05.0000-07:00"
	lineFormat = "%v  %-5s  [%s]  %v\n"

	// LevelDebug will log everything.
	LevelDebug = 0

	// LevelInfo will log messages at INFO and above.
	LevelInfo = 1

	// LevelWarn will log messages at WARN and above.
	LevelWarn = 2

	// LevelError will only log ERROR messages.
	LevelError = 3
)

var levels = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
}

// Logger handles writing structured logs, for a specific program.
//
// For example, given the name, "Foo", logs can look like this...
//
// 2016-12-06T20:27:52.3368+10:00  INFO   [Foo]  Some useful information.
//
// The default Level is 0, which is DEBUG.
type Logger struct {
	Name  string
	Out   io.Writer
	Level int64
}

// New returns a new named Logger, which will be default write to os.Stdout.
func New(name string) Logger {
	return Logger{
		Name: name,
		Out:  os.Stdout,
	}
}

// Debug logs a message at DEBUG level.
func (l Logger) Debug(message string, args ...interface{}) {
	l.write(LevelDebug, message, args)
}

// Info logs a message at INFO level.
func (l Logger) Info(message string, args ...interface{}) {
	l.write(LevelInfo, message, args)
}

// Warn logs a message at WARN level.
func (l Logger) Warn(message string, args ...interface{}) {
	l.write(LevelWarn, message, args)
}

// Error logs a message at ERROR level.
func (l Logger) Error(message string, args ...interface{}) {
	l.write(LevelError, message, args)
}

func (l Logger) write(level int64, message string, args []interface{}) {
	if level < l.Level {
		return
	}

	ts := time.Now().Format(timeFormat)
	data := fmt.Sprintf(message, args...)
	levelName := levels[level]

	line := fmt.Sprintf(lineFormat, ts, levelName, l.Name, data)
	fmt.Fprintf(l.Out, line)
}
