package main

import (
	"flag"
	"github.com/gookit/slog"
	"github.com/nepriyatelev/success_test.git/internal/app"
	"github.com/nepriyatelev/success_test.git/internal/logger"
	"os"
)

var (
	clients int
	logLvl  string
)

func init() {
	flag.IntVar(&clients, "clients", 1, "Количество клиентов для тестирования (по умолчанию 1)")
	flag.StringVar(&logLvl, "log", "info", "Уровень логирования (debug, info, warn, error, fatal, panic)")
}

func main() {
	flag.Parse()
	logger.SetLogLevel(logLvl)
	slog.Info("Тестирование начато")
	slog.Info("Количество клиентов: ", clients)
	if clients < 1 {
		slog.Info("Количество клиентов должно быть больше 0")
		os.Exit(1)
	}

	app.RunApp(clients)
}
