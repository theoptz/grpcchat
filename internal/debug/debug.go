package debug

import "log"

// Logger interface
type Logger interface {
	Debug(string, ...interface{})
}

// NopLog provides logger that does nothing
type NopLog struct{}

// Debug does nothing
func (n *NopLog) Debug(layout string, args ...interface{}) {}

// Log provides basic logger that writes to stdout
type Log struct{}

// Debug writes output to stdout
func (l *Log) Debug(layout string, args ...interface{}) {
	log.Printf(layout, args...)
}

// New returns Logger
func New(debug bool) Logger {
	if !debug {
		return new(NopLog)
	}

	return new(Log)
}
