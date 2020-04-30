package kafka

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
)

// KafkaLibrary ...
type KafkaLibrary struct {
}

// ProducersMessageFormat ...
type StateFullFormat struct {
	UUID      string                 `json:"__uuid" bson:"__uuid"`
	Action    string                 `json:"__action" bson:"__action"`
	Data      map[string]interface{} `json:"data" bson:"data"`
	Offset    int64                  `json:"offset" bson:"offset"`
	CreatedAt *time.Time             `json:"created_at" bson:"created_at"`
}

// HistoryStatefull ...
type HistoryStatefull struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Code        uint   `json:"code" bson:"code"`
}

// KafkaLibraryHandler ...
func KafkaLibraryHandler() *KafkaLibrary {
	return &KafkaLibrary{}
}

// KafkaLibraryInterface ...
type KafkaLibraryInterface interface {
	GetStateFull() *StateFullFormat
	SendEvent(topic string, payload *StateFullFormat) (*StateFullFormat, int64, error)
	InitConsumer(group string) (sarama.ConsumerGroup, error)
}

// GetStateFull ...
func (kafka *KafkaLibrary) GetStateFull() *StateFullFormat {
	return &StateFullFormat{}
}

// Init ...
func (kafka *KafkaLibrary) init(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}

func (kafka *KafkaLibrary) initConsumerConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	switch os.Getenv("KAFKA_STRATEGY") {
	case "STICKY":
		kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "ROUND_ROBIN":
		kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "RANGE":
		kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	}

	kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
