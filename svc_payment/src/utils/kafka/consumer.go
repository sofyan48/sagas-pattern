package kafka

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

// InitConsumer ...
func (kafka *KafkaLibrary) InitConsumer() (sarama.Consumer, error) {
	configKafka := kafka.init("", "")
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")
	fmt.Println(kafkaHost)
	return sarama.NewConsumer([]string{kafkaHost + ":" + kafkaPort}, configKafka)
}
