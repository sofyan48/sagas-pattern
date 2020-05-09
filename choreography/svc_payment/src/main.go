package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sofyan48/svc_payment/src/config"
	"github.com/sofyan48/svc_payment/src/router"
	"github.com/sofyan48/svc_payment/src/worker"
)

// func main() {
// 	environment := flag.String("e", "development", "")
// 	flag.Usage = func() {
// 		fmt.Println("Usage: server -e {mode}")
// 		os.Exit(1)
// 	}
// 	flag.Parse()
// 	config.ConfigEnvironment(*environment)
// 	// worker.LoadCron()
// 	worker.LoadWorker()
// }

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
	// go worker.LoadCron()
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	router.LoadRouter(engine)
	engine.Run(serverString)
}
