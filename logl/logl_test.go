package logl

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLogWriting(t *testing.T) {
	timestamp1, _ := time.Parse(time.RFC3339, "2019-01-02T15:04:05Z")
	timestamp2, _ := time.Parse(time.RFC3339, "2019-02-02T15:04:05Z")

	expectedEntry1 := "2019-01-02T15:04:05Z\ttest-app\tmessage1"
	expectedEntry2 := "2019-02-02T15:04:05Z\ttest-app\tmessage2"

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
	logger.Log("message1", timestamp1, Error)
	if len(logger.buffer) > 0 {
		t.Fatal("logger buffer doesnt flush after log entry writing")
	}
	logger.Log("message2", timestamp2, Error)
	if len(logger.buffer) > 0 {
		t.Fatal("logger buffer doesnt flush after log entry writing")
	}

	dat, _ := ioutil.ReadFile("/tmp/logl_test_log")
	lines := strings.Split(string(dat), "\n")

	if lines[0] != expectedEntry1 {
		t.Fatalf("expected %q entry doesnt match %q entry", expectedEntry1, lines[0])
	}

	if lines[1] != expectedEntry2 {
		t.Fatalf("expected %q entry doesnt match %q entry", expectedEntry2, lines[1])
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
	logger.Log("message1", time.Now(), Info)
	logger.Log("message2", time.Now(), Debug)

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
