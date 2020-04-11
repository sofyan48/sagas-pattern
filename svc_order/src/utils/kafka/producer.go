package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

// InitProducer ...
func (kafka *KafkaLibrary) initProducer() (sarama.SyncProducer, error) {
	configKafka := kafka.initProducerConfig("", "")
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")
	return sarama.NewSyncProducer([]string{kafkaHost + ":" + kafkaPort}, configKafka)
}

// SendEvent ...
func (kafka *KafkaLibrary) SendEvent(topic string, payload *StateFullFormat) (*StateFullFormat, int64, error) {
	now := time.Now()
	fixPayload := &StateFullFormat{}
	fixPayload.Action = payload.Action
	fixPayload.CreatedAt = &now
	fixPayload.Data = payload.Data
	fixPayload.UUID = payload.UUID
	producers, err := kafka.initProducer()
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	data, err := json.Marshal(fixPayload)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}
	_, offset, err := producers.SendMessage(kafkaMsg)
	if err != nil {
		log.Println("KAFKA: ", err)
		return nil, 0, err
	}
	return fixPayload, offset, nil
}
