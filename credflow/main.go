package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/uber-common/cadence-samples/new_samples/worker"
)

func main() {
	worker.StartWorker()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)
	fmt.Println("Cadence worker for Credflow started, press ctrl+c to terminate...")
	<-done
}
