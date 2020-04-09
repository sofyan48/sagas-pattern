package worker

import (
	"os"

	"github.com/sofyan48/svc_payment/src/app/v1/consumer"
)

// LoadWorker ...
func LoadWorker() {
	signals := make(chan os.Signal, 1)
	V1ConsumerWorker := consumer.V1OrderEventsHandler()
	V1ConsumerWorker.Consume([]string{os.Getenv("APP_NAME")}, signals)
}
