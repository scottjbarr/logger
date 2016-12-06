package logger

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var validTime = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.\d{4}\+\d{2}:\d{2}`)

func TestLoggerWithArgs(t *testing.T) {
	data := []string{"INFO", " [Test]", "Don't forget your towel\n"}

	log := New("Test")
	b := []byte{}
	log.Out = bytes.NewBuffer(b)

	log.Info("Don't forget your %v", "towel")

	actual := fmt.Sprintf("%s", log.Out)

	parts := strings.SplitN(actual, "  ", 4)

	if !validTime.MatchString(parts[0]) {
		t.Fatalf("Expected '%s', to match time format", actual)
	}

	if parts[1] != data[0] {
		t.Fatalf("Expected '%s', to be %s", parts[1], data[0])
	}

	if parts[2] != data[1] {
		t.Fatalf("Expected '%s', to be %s", parts[2], data[1])
	}

	if parts[3] != data[2] {
		t.Fatalf("Expected '%s', to be %s", parts[3], data[2])
	}
}

func TestLoggerWithoutArgs(t *testing.T) {
	data := []string{"DEBUG", "[Test]", "Don't forget\n"}

	log := New("Test")
	b := []byte{}
	log.Out = bytes.NewBuffer(b)

	log.Debug("Don't forget")

	actual := fmt.Sprintf("%s", log.Out)

	parts := strings.SplitN(actual, "  ", 4)

	if !validTime.MatchString(parts[0]) {
		t.Fatalf("Expected '%s', to match time format", actual)
	}

	if parts[1] != data[0] {
		t.Fatalf("Expected '%s', to be %s", parts[1], data[0])
	}

	if parts[2] != data[1] {
		t.Fatalf("Expected '%s', to be %s", parts[2], data[1])
	}

	if parts[3] != data[2] {
		t.Fatalf("Expected '%s', to be %s", parts[3], data[2])
	}
}

func TestLogLevel(t *testing.T) {
	log := New("Toast")
	log.Level = LevelError
	b := []byte{}
	log.Out = bytes.NewBuffer(b)

	log.Warn("You will not see this")

	actual := fmt.Sprintf("%s", log.Out)

	if len(actual) > 0 {
		t.Fatalf("Expected Warn log message to not be logged")
	}
}
