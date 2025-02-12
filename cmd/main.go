package cmd

import (
	"gDDNS/internal/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var projectName = ""
var version = ""
var gitRev = ""
var buildTime = ""

func main() {
	shutdown := make(chan os.Signal, 1)
	logger.InitLog()

	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTSTP)
	<-shutdown
	log.Info("shutting down")

}
