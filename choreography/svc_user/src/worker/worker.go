package worker

import (
	"os"

	"github.com/sofyan48/svc_user/src/app/v1/worker/consumer"
)

// LoadWorker ...
func LoadWorker() {
	V1ConsumerWorker := consumer.V1ConsumerEventsHandler()
	V1ConsumerWorker.Consume(os.Getenv("APP_NAME"), os.Getenv("APP_NAME"))
}
