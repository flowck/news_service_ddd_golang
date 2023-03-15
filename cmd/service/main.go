package main

import (
	"github.com/flowck/news_service_ddd_golang/internal/common/logs"
)

var (
	Version = "" // injected during build
)

func main() {
	logger := logs.New(true)

	logger.Info("news_service")
}
