package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sofyan48/svc_payment/src/config"
	"github.com/sofyan48/svc_payment/src/worker"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.ConfigEnvironment(*environment)
	// worker.LoadCron()
	worker.LoadWorker()
}
