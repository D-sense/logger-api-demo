//Handles the application logging processes
package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// LogEvent stores messages to log later, from our standard interface
type LogEvent struct {
	id      int
	message string
}

// MainLogger enforces specific log message formats
type MainLogger struct {
	*logrus.Logger
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	///log.SetLevel(log.WarnLevel)
}

// NewLogger initializes the standard logger
func NewLogger() *MainLogger {
	f, err := os.OpenFile("logrus.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	var baseLogger = logrus.New()
	var standardLogger = &MainLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	mw := io.MultiWriter(os.Stdout, f)
	standardLogger.SetOutput(mw)

	return standardLogger
}

// Variables to store our log messages as new Events
var (
	invalidArgMessage      = LogEvent{1, "Invalid arg: %s"}
	invalidArgValueMessage = LogEvent{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = LogEvent{3, "Missing arg: %s"}
)

// Standard "InvalidArg error message
func (l *MainLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

// Standard "InvalidArgValue error message
func (l *MainLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// "Standard "MissingArg error message"
func (l *MainLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}
