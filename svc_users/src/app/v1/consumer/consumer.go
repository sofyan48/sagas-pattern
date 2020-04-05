package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sofyan48/svc_user/src/utils/kafka"
)

// V1ConsumerEvents ...
type V1ConsumerEvents struct {
	Kafka kafka.KafkaLibraryInterface
}

// V1ConsumerEventsHandler ...
func V1ConsumerEventsHandler() *V1ConsumerEvents {
	return &V1ConsumerEvents{
		Kafka: kafka.KafkaLibraryHandler(),
	}
}

// V1ConsumerEventsInterface ...
type V1ConsumerEventsInterface interface {
	Consume(topics []string, signals chan os.Signal)
}

// Consume ...
func (consumer *V1ConsumerEvents) Consume(topics []string, signals chan os.Signal) {
	StateFullData := consumer.Kafka.GetStateFull()
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
			fmt.Println(partition)
			go consumeMessage(csm, topic, partition, chanMessage)
		}
	}
	log.Println("Event is Started....")

ConsumerLoop:
	for {
		select {
		case msg := <-chanMessage:
			json.Unmarshal(msg.Value, StateFullData)
			log.Println("New Event from , Event: ", StateFullData.Action)
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
