package logger

import "github.com/gookit/slog"

func SetLogLevel(logLvl string) {
	if logLvl == "debug" {
		slog.SetLogLevel(slog.DebugLevel)
	} else if logLvl == "info" {
		slog.SetLogLevel(slog.InfoLevel)
	} else if logLvl == "warn" {
		slog.SetLogLevel(slog.WarnLevel)
	} else if logLvl == "error" {
		slog.SetLogLevel(slog.ErrorLevel)
	} else if logLvl == "fatal" {
		slog.SetLogLevel(slog.FatalLevel)
	} else if logLvl == "panic" {
		slog.SetLogLevel(slog.PanicLevel)
	} else {
		slog.SetLogLevel(slog.InfoLevel)
		slog.Warn("Уровень логирования должен быть одним из: debug, info, warn, error, fatal, panic")
	}
}
