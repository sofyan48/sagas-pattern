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
	"github.com/sofyan48/svc_user/src/app/v1/worker/consumer/controller"
	"github.com/sofyan48/svc_user/src/app/v1/worker/entity"
	"github.com/sofyan48/svc_user/src/utils/kafka"
)

// V1ConsumerEvents ...
type V1ConsumerEvents struct {
	Kafka      kafka.KafkaLibraryInterface
	Controller controller.ControllerEventInterface
	ready      chan bool
}

// V1ConsumerEventsHandler ...
func V1ConsumerEventsHandler() *V1ConsumerEvents {
	return &V1ConsumerEvents{
		Kafka:      kafka.KafkaLibraryHandler(),
		Controller: controller.ControllerEventHandler(),
		ready:      make(chan bool),
	}
}

// V1ConsumerEventsInterface ...
type V1ConsumerEventsInterface interface {
	Consume(topics []string, signals chan os.Signal)
}

// Consume ...
func (consumer *V1ConsumerEvents) Consume(topics string, group string) {
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
func (consumer *V1ConsumerEvents) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *V1ConsumerEvents) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *V1ConsumerEvents) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		eventData := &entity.StateFullFormatKafka{}
		json.Unmarshal(message.Value, eventData)
		switch eventData.Action {
		case "users":
			consumer.Controller.UserLoad(eventData)
		case "login":
			consumer.Controller.LoginLoad(eventData)
		default:
			fmt.Println("OK")
		}
		log.Println("EV Receive: ", message.Timestamp, " | Topic: ", message.Topic)
		session.MarkMessage(message, "")
	}
	return nil
}
