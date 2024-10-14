package commonlib

import (
	"os"
	"os/signal"
	"syscall"
)

func KeepServerAlive() {

	// Waits for SIGTERM signal on a channel
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
