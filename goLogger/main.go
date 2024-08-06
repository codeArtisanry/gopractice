// help me write a golang logger for info, warning, error, debug, trace and critical levels with interface
// and also write a test case for the logger
package main

import "fmt"



const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

type Logger interface {
  Trace(msg string)
  Info(msg string)
  Warning(msg string)
  Error(msg string)
  Debug(msg string)
  Critical(msg string)
}

type ConsoleLogger struct{}

func (l *ConsoleLogger) Trace(msg string) {
	fmt.Println("TRACE:", msg)
}

func (l *ConsoleLogger) Debug(msg string) {
	fmt.Println("DEBUG:", msg)
}

func (l *ConsoleLogger) Info(msg string) {
	fmt.Println("INFO:", msg)
}

func (l *ConsoleLogger) Warning(msg string) {
	fmt.Println("WARNING:", msg)
}

func (l *ConsoleLogger) Error(msg string) {
	fmt.Println("ERROR:", msg)
}

func (l *ConsoleLogger) Critical(msg string) {
	fmt.Println("CRITICAL:", msg)
}

func main() {
	logger := &ConsoleLogger{}

	logger.Trace("This is a trace message.")
	logger.Debug("This is a debug message.")
	logger.Info("This is an info message.")
	logger.Warning("This is a warning message.")
	logger.Error("This is an error message.")
	logger.Critical("This is a critical message.")
}

