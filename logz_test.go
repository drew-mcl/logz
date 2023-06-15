package logz

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestSetLevelInvalid(t *testing.T) {
	var buf bytes.Buffer
	originalOutput := Output
	Output = &buf                              // redirect output to buf
	defer func() { Output = originalOutput }() // restore original output after the test

	SetLevel("INVALID")
	expected := "Invalid level: INVALID\n"
	assert.Contains(t, buf.String(), expected, "Expected error message")
}

func TestSetLevelLowerCase(t *testing.T) {
	SetLevel("info")
	assert.Equal(t, INFO, level, "Expected level to be set to INFO")
}

func TestLogFunctions(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf

	// Set log level to INFO
	level = INFO

	// Call log functions
	Debug("Debug message")
	Info("Info message")
	Warn("Warning message")
	Error("Error message")

	expected := "[INFO ] Info message\n[WARN ] Warning message\n[ERROR] Error message\n"
	assert.Contains(t, buf.String(), expected, "Expected log output")
}

func TestInfoWithSuccess(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf

	// Set log level to INFO
	level = INFO

	InfoWithSuccess("Operation completed")

	expected := "[INFO ] " + color.GreenString("Success") + ": Operation completed\n"
	assert.Contains(t, buf.String(), expected, "Expected log output")
}

func TestDisableColors(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf
	defer func() { Output = os.Stdout }()

	DisableColors()
	Info("Test message")

	expected := "[INFO ] Test message\n"
	assert.Contains(t, buf.String(), expected, "Expected log output without color codes")
}

func TestSetLevelCaseInsensitive(t *testing.T) {
	SetLevel("info")
	assert.Equal(t, INFO, level, "Expected level to be set to INFO")

	SetLevel("INFO")
	assert.Equal(t, INFO, level, "Expected level to be set to INFO")

	SetLevel("InFo")
	assert.Equal(t, INFO, level, "Expected level to be set to INFO")
}

func TestInvalidLevelDoesNotChangeLevel(t *testing.T) {
	SetLevel("INFO")
	assert.Equal(t, INFO, level, "Expected level to be set to INFO")

	SetLevel("INVALID")
	assert.Equal(t, INFO, level, "Expected level to still be INFO after invalid SetLevel call")
}

func TestLogMultipleArguments(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf
	defer func() { Output = os.Stdout }()

	level = INFO
	Info("Test", "message", 123)

	expected := "[INFO ] Test message 123\n"
	assert.Contains(t, buf.String(), expected, "Expected log output with multiple arguments")
}

func benchmarkLogLatency(logFunc func(...interface{}), b *testing.B) {
	for i := 0; i < b.N; i++ {
		logFunc("Log message")
	}
}
func BenchmarkCustomLog(b *testing.B) {
	Info("Warm up")
	b.ResetTimer()
	benchmarkLogLatency(Info, b)
}

func BenchmarkStandardLog(b *testing.B) {
	log.Println("Warm up")
	b.ResetTimer()
	benchmarkLogLatency(log.Println, b)
}
