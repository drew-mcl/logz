package logz

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
)

// Define log levels as constants
const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
)

// Define the level mapping
var levelMapping = map[string]int{
	"TRACE": TRACE,
	"DEBUG": DEBUG,
	"INFO":  INFO,
	"WARN":  WARN,
	"ERROR": ERROR,
}

// Current logging level
var level = DEBUG
var lock sync.Mutex

// Setting custom output to std out over print as its easier to test
var Output io.Writer = os.Stdout

// SetLevel function to set the current logging level
func SetLevel(logLevel string) {
	if val, exists := levelMapping[strings.ToUpper(logLevel)]; exists {
		level = val
	} else {
		fmt.Fprintf(Output, "Invalid level: %s\n", logLevel)
	}
}

// DisableColors disables colorized output
func DisableColors() {
	color.NoColor = true
}

// Trace log function
func Trace(v ...interface{}) {
	if level <= DEBUG {
		fmt.Fprintf(color.Output, "[%s] %v\n", color.WhiteString("TRACE"), strings.Join(formatArgs(v...), " "))
	}
}

// Debug log function
func Debug(v ...interface{}) {
	if level <= DEBUG {
		fmt.Fprintf(color.Output, "[%s] %v\n", color.MagentaString("DEBUG"), strings.Join(formatArgs(v...), " "))
	}
}

// Info log function
func Info(v ...interface{}) {
	if level <= INFO {
		fmt.Fprintf(color.Output, "[%s] %v\n", color.CyanString("INFO"), strings.Join(formatArgs(v...), " "))
	}
}

// Info log function with Success: in green before the log message
func InfoWithSuccess(v ...interface{}) {
	if level <= INFO {
		fmt.Fprintf(color.Output, "[%s] %s: %v\n", color.CyanString("INFO"), color.GreenString("Success"), strings.Join(formatArgs(v...), " "))
	}
}

// Warning log function
func Warn(v ...interface{}) {
	if level <= WARN {
		fmt.Fprintf(color.Output, "[%s] %v\n", color.YellowString("WARN"), strings.Join(formatArgs(v...), " "))
	}
}

// Error log function
func Error(v ...interface{}) {
	if level <= ERROR {
		fmt.Fprintf(color.Output, "[%s] %v\n", color.RedString("ERROR"), strings.Join(formatArgs(v...), " "))
	}
}

// Helper function to format log arguments as strings
func formatArgs(v ...interface{}) []string {
	formattedArgs := make([]string, len(v))
	for i, arg := range v {
		formattedArgs[i] = fmt.Sprintf("%v", arg)
	}
	return formattedArgs
}
