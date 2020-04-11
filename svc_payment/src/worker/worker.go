package worker

import (
	"os"

	"github.com/sofyan48/svc_payment/src/app/v1/consumer"
	"github.com/sofyan48/svc_payment/src/app/v1/task"
)

// LoadWorker ...
func LoadWorker() {
	V1ConsumerWorker := consumer.V1OrderEventsHandler()
	V1ConsumerWorker.Consume(os.Getenv("APP_NAME"), os.Getenv("APP_NAME"))
}

// LoadCron ...
func LoadCron() {
	V1Cron := task.TaskCronHandler()
	go V1Cron.Every2Minutes()
}
