package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/flowck/news_service_ddd_golang/internal/common/logs"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port      int16  `envconfig:"PORT"`
	DebugMode string `envconfig:"FLAG_DEBUG_MODE"`
}

func main() {
	cfg, err := getConfig()
	if err != nil {
		panic(err)
	}

	logger := logs.New(cfg.DebugMode == "enabled")
	logger.Info("news_service")

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	<-done
	logger.Info("Preparing to shutdown gracefully")

	// ctxTerm, cancelCtxTerm := context.WithTimeout(ctx, time.Second*15)
	// defer cancelCtxTerm()

	logger.Info("The service has been terminated")
}

func getConfig() (*Config, error) {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// func fatalOnErr(logger *logs.Logger, err error) {
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// }
