package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/sofyan48/svc_payment/src/app/v2/worker/consumer/controller"
	"github.com/sofyan48/svc_payment/src/app/v2/worker/entity"
	"github.com/sofyan48/svc_payment/src/utils/kafka"
)

// V1OrderEvents ...
type V1OrderEvents struct {
	Kafka      kafka.KafkaLibraryInterface
	Controller controller.ControllerEventInterface
	ready      chan bool
}

// V1OrderEventsHandler ...
func V1OrderEventsHandler() *V1OrderEvents {
	return &V1OrderEvents{
		Kafka:      kafka.KafkaLibraryHandler(),
		Controller: controller.ControllerEventHandler(),
		ready:      make(chan bool),
	}
}

// V1OrderEventsInterface ...
type V1OrderEventsInterface interface {
	Consume(topics []string, signals chan os.Signal)
}

// Consume ...
func (consumer *V1OrderEvents) Consume(topics string, group string) {
	ctx, cancel := context.WithCancel(context.Background())
	client, err := consumer.Kafka.InitConsumer(group)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(topics, ","), consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Service Ready!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
		os.Exit(0)
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

// Setup ...
func (consumer *V1OrderEvents) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *V1OrderEvents) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *V1OrderEvents) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		eventData := &entity.StateFullFormatKafka{}
		json.Unmarshal(message.Value, eventData)
		switch eventData.Action {
		case "payment_save":
			consumer.Controller.PaymentSave(eventData)
		case "payment_order":
			consumer.Controller.PaymentPaidOrder(eventData)
		default:
			fmt.Println("OK")
		}
		log.Println("EV Receive: ", message.Timestamp, " | Topic: ", message.Topic, " | Action: ", eventData.Action)
		session.MarkMessage(message, "")
	}
	return nil
}
