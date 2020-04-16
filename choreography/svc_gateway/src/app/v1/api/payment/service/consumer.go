package service

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

// Consume ...
func (service *PaymentService) Consume(topics string, group string) {
	ctx, cancel := context.WithCancel(context.Background())
	client, err := service.Kafka.InitConsumer(group)
	if err != nil {
		log.Panicf("Error creating service group client: %v", err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := client.Consume(ctx, strings.Split(topics, ","), service); err != nil {
			log.Panicf("Error from service: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
		service.ready = make(chan bool)
	}()

	<-service.ready // Await till the service has been set up
	log.Println("Service Ready!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

// Setup ...
func (service *PaymentService) Setup(sarama.ConsumerGroupSession) error {
	close(service.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (service *PaymentService) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a service loop of ConsumerGroupClaim's Messages().
func (service *PaymentService) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Println("EV Receive: ", message.Timestamp, " | Topic: ", message.Topic)
		session.MarkMessage(message, "")
	}
	return nil
}
