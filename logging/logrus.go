package logging

import (
	"io"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

// NewLogrus returns a properly configured logrus logger fullfilling the Logger interface
func NewLogrus(level logrus.Level, stdout io.Writer, stderr io.Writer, file io.Writer) *logrus.Logger {
	logger := logrus.New()
	if level == logrus.DebugLevel {
		logger.SetReportCaller(true)
	}

	// All logs should go to a file
	logger.SetOutput(file)

	// stderr recieves only approriate logs (as well as stdout)
	logger.AddHook(&writer.Hook{
		Writer: stderr,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	})
	logger.AddHook(&writer.Hook{
		Writer: stdout,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})

	return logger
}
