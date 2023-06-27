package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ardanlabs/service/foundation/logger"
)

func main() {
	log := logger.New(os.Stdout)

	if err := run(log); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(log *logger.Logger) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	log.Info("shutdown", "status", "shutdown started")
	defer log.Info("shutdown", "status", "shutdown complete")

	return nil
}
