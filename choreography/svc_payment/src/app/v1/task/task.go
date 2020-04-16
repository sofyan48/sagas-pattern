package task

import (
	"log"
	"os"
	"time"

	"github.com/sofyan48/svc_payment/src/app/v1/repository"
	"github.com/sofyan48/svc_payment/src/utils/kafka"
	"github.com/sofyan48/svc_payment/src/utils/logger"
)

// TaskCron ...
type TaskCron struct {
	Repository repository.PaymentRepositoryInterface
	Kafka      kafka.KafkaLibraryInterface
	Logger     logger.LoggerInterface
}

// TaskCronHandler ...
func TaskCronHandler() *TaskCron {
	return &TaskCron{
		Repository: repository.PaymentRepositoryHandler(),
		Kafka:      kafka.KafkaLibraryHandler(),
		Logger:     logger.LoggerHandler(),
	}
}

// TaskCronInterface ...
type TaskCronInterface interface {
	Every2Minutes()
}

// Every2Minutes ...
func (cron *TaskCron) Every2Minutes() {
	log.Println("Task Running")
	waitingValue := os.Getenv("CRON_WAITING_VALUE")
	for {
		data, err := cron.Repository.GetPaymentByStatus(waitingValue)
		if err != nil {
			log.Println("Check Status | Cron Failed")
		}
		for _, i := range data {
			// sending to status order
			if i.ChangeTotal != 0 {
				cron.updateCompleteOrder(i.UUID, i.UUIDOrder, "")
			}
			cron.updateCompleteOrder(i.UUID, i.UUIDOrder, "Complete")

		}
		time.Sleep(2 * time.Minute)
	}
}

// Every1Day ...
func (cron *TaskCron) Every1Day() {

}

func (cron *TaskCron) updateCompleteOrder(UUID, UUIDOrder, status string) {
	// sending order prepare
	now := time.Now()
	payloadPayment := cron.Kafka.GetStateFull()
	payloadPayment.Action = "order_update"
	payloadPayment.CreatedAt = &now
	payloadPayment.UUID = UUID
	payloadPayment.Data = map[string]interface{}{
		"uuid_order":     UUIDOrder,
		"payment_status": status,
	}
	resultOrder, _, err := cron.Kafka.SendEvent("order", payloadPayment)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		cron.Logger.Save(UUID, "failed", loggerData)
		return
	}
	paymentLog := map[string]interface{}{
		"code":     "200",
		"messages": "Order Status Update",
		"result":   resultOrder,
	}
	cron.Logger.Save(UUID, "success", paymentLog)
}
