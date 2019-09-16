// Package logl is a simple logging library.
// Every logging event is set
//
// Usage:
//
// var opts LoggerOptions = LoggerOptions{
// 	Dest: os.Open("some/file"),
// 	Source: "my-app",
// 	Separator: "] [",
//	Level: logl.Info
// }
//
// logger := logl.NewLogger(opts)
//
// logger.Log("my super duper event fired!", logl.Info)
//
package logl

import (
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type LogLevel uint8

const (
	Error LogLevel = iota
	Info
	Debug
)

type LoggerOptions struct {
	Dest      io.Writer
	Source    string
	Separator string
	Level     LogLevel
}

type Logger struct {
	mu        sync.Mutex
	dest      io.Writer
	buffer    []byte
	source    string
	separator string
	level     LogLevel
}

// Returns a pointer to new logger object.
// If not set, dest will be set to stdout and separator will be set to \t (tab).
func NewLogger(lo LoggerOptions) *Logger {
	if lo.Dest == nil {
		lo.Dest = os.Stdout
	}
	if lo.Separator == "" {
		lo.Separator = "\t"
	}
	return &Logger{dest: lo.Dest, source: lo.Source, separator: lo.Separator, level: lo.Level}
}

func (l *Logger) Log(s string, timestamp time.Time, level LogLevel) error {
	if l.level < level {
		return nil
	}

	timestamp = timestamp.UTC()

	l.mu.Lock()
	defer l.mu.Unlock()

	writePrefix(l, timestamp)

	l.buffer = append(l.buffer, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buffer = append(l.buffer, '\n')
	}

	_, err := l.dest.Write(l.buffer)
	l.buffer = l.buffer[:0]
	return err
}

func writePrefix(l *Logger, timestamp time.Time) {
	parts := []string{timestamp.Format(time.RFC3339)}
	if len(l.source) > 0 {
		parts = append(parts, l.source)
	}
	prefix := strings.Join(parts, l.separator)
	l.buffer = append(l.buffer, prefix...)
	l.buffer = append(l.buffer, l.separator...)
}
