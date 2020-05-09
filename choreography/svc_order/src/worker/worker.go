package worker

import (
	"os"

	"github.com/sofyan48/svc_order/src/app/v2/worker/consumer"
)

// LoadWorker ...
func LoadWorker() {
	// V1ConsumerWorker := consumer.V1OrderEventsHandler()
	// V1ConsumerWorker.Consume(os.Getenv("APP_NAME"), os.Getenv("APP_NAME"))
	V2ConsumerWorker := consumer.V2OrderEventsHandler()
	V2ConsumerWorker.Consume(os.Getenv("APP_NAME"), os.Getenv("APP_NAME"))
}
