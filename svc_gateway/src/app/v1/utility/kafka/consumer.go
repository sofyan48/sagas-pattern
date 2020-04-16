package kafka

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

// InitConsumer ...
func (kafka *KafkaLibrary) InitConsumer(group string) (sarama.ConsumerGroup, error) {
	configKafka := kafka.initConsumerConfig("", "")
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")
	version, err := sarama.ParseKafkaVersion(os.Getenv("KAFKA_VERSION"))
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}
	configKafka.Version = version
	return sarama.NewConsumerGroup([]string{kafkaHost + ":" + kafkaPort}, group, configKafka)
}
