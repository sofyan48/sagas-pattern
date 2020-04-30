package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sofyan48/order/src/config"

	apiRouter "github.com/sofyan48/order/src/router"
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
	router := config.SetupEngine(env)
	apiRouter.LoadRouter(router)
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	router.Run(serverString)
}
