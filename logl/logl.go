// Package logl is a simple logging library.
//
// Usage:
//
// var opts LoggerOptions = LoggerOptions{
// 	Dest: os.Open("some/file"), // any io.Writer is supported; default to stdout
// 	Source: "my-app",
// 	Separator: " // ", // default to \\t
// 	Level: logl.Info,
// 	BufferLength: 1024 // default to 4096, min 1024
// }
//
// logger := logl.NewLogger(opts)
//
// logger.Log("my super duper event fired!", time.Now(), logl.Info)
//
package logl

import (
	"bytes"
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
	Dest         io.Writer
	Source       string
	Separator    string
	Level        LogLevel
	BufferLength uint
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
	if lo.BufferLength == 0 {
		lo.BufferLength = 4096
	}
	if lo.BufferLength < 1024 {
		lo.BufferLength = 1024
	}
	return &Logger{buffer: make([]byte, lo.BufferLength), dest: lo.Dest, source: lo.Source, separator: lo.Separator, level: lo.Level}
}

func (l *Logger) Log(s string, timestamp time.Time, level LogLevel) error {
	if l.level < level {
		return nil
	}

	var writeError error

	timestamp = timestamp.UTC()

	l.mu.Lock()
	defer l.mu.Unlock()

	if len(s) > cap(l.buffer)-len(l.buffer) {
		writeError = l.FlushBuffer()
	}

	writePrefix(l, timestamp)

	l.buffer = append(l.buffer, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buffer = append(l.buffer, '\n')
	}

	return writeError
}

func (l *Logger) FlushBuffer() error {
	_, err := l.dest.Write(bytes.Trim(l.buffer, "\x00"))
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
