package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sofyan48/svc_order/src/config"
	"github.com/sofyan48/svc_order/src/router"
	"github.com/sofyan48/svc_order/src/worker"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	startApp(*environment)
}

func startApp(env string) {
	engine := config.SetupEngine(env)

	go worker.LoadWorker()

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	router.LoadRouter(engine)
	engine.Run(serverString)
}
