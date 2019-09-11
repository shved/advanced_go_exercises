package logl

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestLogWriting(t *testing.T) {
	logEntryPattern, _ := regexp.Compile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z\t\S+\t\S+\n`)

	f, err := os.Create("/tmp/logl_test_log")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	opts := LoggerOptions{
		Dest:   f,
		Source: "test-app",
		Level:  Info,
	}
	logger := NewLogger(opts)
	logger.Log("message1", Error)
	logger.Log("message2", Error)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := fmt.Sprintln(scanner.Text())
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		if !logEntryPattern.MatchString(line) {
			t.Fatalf("given %v log entry doesnt match expected pattern", line)
		}
	}
}

func TestLogLevel(t *testing.T) {
	f, err := os.Create("/tmp/logl_test_log")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	opts := LoggerOptions{
		Dest:   f,
		Source: "test-app",
		Level:  Error,
	}
	logger := NewLogger(opts)
	logger.Log("message1", Info)
	logger.Log("message2", Debug)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := fmt.Sprintln(scanner.Text())
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		if len(line) > 0 {
			t.Fatal("log level doesnt filter message")
		}
	}

}

func TestNewLoggerDefaultValues(t *testing.T) {
	l := NewLogger(LoggerOptions{})

	if l.dest != os.Stdout {
		t.Fatalf("expected logger to have default destination set to os.Stdout, but was set to %v", l.dest)
	}

	if l.separator != "\t" {
		t.Fatalf("expected logger to have default separator set to \\t, but was set to %v", l.separator)
	}
}
