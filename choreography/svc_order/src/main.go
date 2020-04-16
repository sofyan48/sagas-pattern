package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sofyan48/svc_order/src/config"
	"github.com/sofyan48/svc_order/src/worker"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.ConfigEnvironment(*environment)
	worker.LoadWorker()
}
