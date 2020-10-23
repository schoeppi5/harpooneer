package logging

// Logger represents an instance of a logger
// An implementation of this interface is passed to all plugins
type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Print(...interface{})
	Printf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Panic(...interface{})
	Panicf(string, ...interface{})
}
