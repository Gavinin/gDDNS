package cmd

import (
	"gDDNS/internal/log"
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

	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTSTP)
	<-shutdown
	log.Log.Info("shutting down")

}
