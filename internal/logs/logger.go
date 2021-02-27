package logs

import (
	"fmt"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

//InitLogger ...
func InitLogger() error {
	// NewDevelopment builds a development Logger that writes DebugLevel and above
	// logs to standard error in a human-friendly format.
	//
	// It's a shortcut for NewDevelopmentConfig().Build(...Option).
	l, err := zap.NewDevelopment()
	if err != nil {
		_ = fmt.Errorf("cannot create zap logger %s", err.Error())
		return err
	}
	// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
	// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
	// single application to use both Loggers and SugaredLoggers, converting
	// between them on the boundaries of performance-sensitive code.
	sugar = l.Sugar()
	return nil
}

//Log ..
func Log() *zap.SugaredLogger {
	return sugar
}
