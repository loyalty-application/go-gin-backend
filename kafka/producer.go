package kafka

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/loyalty-application/go-gin-backend/models"
)

func KafkaProduce(t models.Transaction) {
	server := os.Getenv("KAFKA_BOOTSTRAP_SERVER")
	topic := os.Getenv("KAFKA_TOPIC")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})

	
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	key := t.CardId
	data,_ := json.Marshal(t)

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          data,
	}, nil)

	// Wait for all messages to be delivered
	p.Flush(15 * 1000)
	p.Close()
}
