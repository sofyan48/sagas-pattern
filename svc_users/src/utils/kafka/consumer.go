package kafka

import (
	"os"

	"github.com/Shopify/sarama"
)

// InitConsumer ...
func (kafka *KafkaLibrary) InitConsumer() (sarama.Consumer, error) {
	configKafka := kafka.init("", "")
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")
	return sarama.NewConsumer([]string{kafkaHost + ":" + kafkaPort}, configKafka)
}
