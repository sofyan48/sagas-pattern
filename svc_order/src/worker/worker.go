package worker

import (
	"os"

	"github.com/sofyan48/svc_order/src/app/v1/consumer"
	"github.com/sofyan48/svc_order/src/registry"
)

// LoadWorker ...
func LoadWorker() {
	V1ConsumerWorker := consumer.V1OrderEventsHandler()
	V1ConsumerWorker.Consume(os.Getenv("APP_NAME"), os.Getenv("APP_NAME"))
}

// LoadRegistry ...
func LoadRegistry() {
	registry := registry.ServiceRegistryHandler()
	// load healthchecks
	go registry.Health("/healthchecks")
}
