package worker

import (
	"os"

	"github.com/sofyan48/svc_payment/src/app/v2/worker/consumer"
	"github.com/sofyan48/svc_payment/src/app/v2/worker/task"
)

// LoadWorker ...
func LoadWorker() {
	V2ConsumerWorker := consumer.V1OrderEventsHandler()
	V2ConsumerWorker.Consume(os.Getenv("APP_NAME"), os.Getenv("APP_NAME"))
}

// LoadCron ...
func LoadCron() {
	V1Cron := task.TaskCronHandler()
	go V1Cron.Every2Minutes()
}
