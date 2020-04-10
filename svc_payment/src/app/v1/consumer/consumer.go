package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sofyan48/svc_payment/src/app/v1/entity"
	"github.com/sofyan48/svc_payment/src/app/v1/event"
	"github.com/sofyan48/svc_payment/src/utils/kafka"
	"github.com/sofyan48/svc_payment/src/utils/logger"
)

// V1OrderEvents ...
type V1OrderEvents struct {
	Kafka  kafka.KafkaLibraryInterface
	Event  event.PaymentEventInterface
	Logger logger.LoggerInterface
}

// V1OrderEventsHandler ...
func V1OrderEventsHandler() *V1OrderEvents {
	return &V1OrderEvents{
		Kafka:  kafka.KafkaLibraryHandler(),
		Event:  event.PaymentEventHandler(),
		Logger: logger.LoggerHandler(),
	}
}

// V1OrderEventsInterface ...
type V1OrderEventsInterface interface {
	Consume(topics []string, signals chan os.Signal)
}

// Consume ...
func (consumer *V1OrderEvents) Consume(topics []string, signals chan os.Signal) {
	// StateFullData := consumer.Kafka.GetStateFull()
	chanMessage := make(chan *sarama.ConsumerMessage, 256)
	csm, err := consumer.Kafka.InitConsumer()
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	for _, topic := range topics {
		partitionList, err := csm.Partitions(topic)
		if err != nil {
			log.Println("Unable to get partition got error ", err)
			continue
		}
		for _, partition := range partitionList {
			go consumeMessage(csm, topic, partition, chanMessage)
		}
	}
	log.Println("Event is Started....")

ConsumerLoop:
	for {
		select {
		case msg := <-chanMessage:
			eventData := &entity.StateFullFormatKafka{}
			json.Unmarshal(msg.Value, eventData)
			switch eventData.Action {
			case "payment_save":
				consumer.paymentSave(eventData)
			case "payment_order":
				consumer.paymentPaidOrder(eventData)
			default:
				fmt.Println("OK")
			}

		case sig := <-signals:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Println("Unable to consume partition got error ", partition, err)
		return
	}
	defer func() {
		if err := msg.Close(); err != nil {
			log.Println("Unable to close partition : ", partition, err)
		}
	}()
	for {
		msg := <-msg.Messages()
		c <- msg
	}

}

func (consumer *V1OrderEvents) paymentSave(paymentData *entity.StateFullFormatKafka) {
	result, err := consumer.Event.InsertDatabase(paymentData)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		consumer.Logger.Save(paymentData.UUID, "failed", loggerData)
		return
	}
	loggerData := map[string]interface{}{
		"code":   "200",
		"result": result,
	}
	consumer.Logger.Save(paymentData.UUID, "success", loggerData)

}

func (consumer *V1OrderEvents) paymentPaidOrder(paymentData *entity.StateFullFormatKafka) {
	result, err := consumer.Event.PaymentUpdateOrder(paymentData)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		fmt.Println(loggerData)
		consumer.Logger.Save(paymentData.UUID, "failed", loggerData)
		return
	}
	loggerData := map[string]interface{}{
		"code":   "200",
		"result": result,
	}
	consumer.Logger.Save(paymentData.UUID, "success", loggerData)

	// sending payment prepare
	now := time.Now()
	payloadPayment := consumer.Kafka.GetStateFull()
	payloadPayment.Action = "order_update"
	payloadPayment.CreatedAt = &now
	payloadPayment.UUID = result.UUID
	payloadPayment.Data = map[string]interface{}{
		"uuid_order":     result.UUIDOrder,
		"payment_status": "Process",
	}
	resultOrder, _, err := consumer.Kafka.SendEvent("order", payloadPayment)
	if err != nil {
		loggerData := map[string]interface{}{
			"code":  "400",
			"error": err,
		}
		consumer.Logger.Save(paymentData.UUID, "failed", loggerData)
		return
	}
	paymentLog := map[string]interface{}{
		"code":     "200",
		"messages": "Order Status Update",
		"result":   resultOrder,
	}
	consumer.Logger.Save(paymentData.UUID, "success", paymentLog)
}
