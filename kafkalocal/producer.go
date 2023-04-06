package kafkalocal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/loyalty-application/go-gin-backend/models"
)

var p *kafka.Producer = CreateProducer()
var server = os.Getenv("KAFKA_BOOTSTRAP_SERVER")
var topic = os.Getenv("KAFKA_TOPIC")

func CreateProducer() (producer *kafka.Producer) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}
	return p

}

// func KafkaProduce(t models.Transaction) {
//   key := t.CardId
//   data,_ := json.Marshal(t)

//   delivery_chan := make(chan kafka.Event, 10000)
//   p.Produce(&kafka.Message{
//     TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
//     Key:            []byte(key),
//     Value:          data,
//   },delivery_chan)

//   e := <-delivery_chan
//    m := e.(*kafka.Message)
//   if m.TopicPartition.Error != nil {
//     fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
//   } else {
//     fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
//         *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
//   }
//   // close(delivery_chan)
// }

func ProduceMessage(t models.Transaction) {

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)

	key := t.CardId
	data, _ := json.Marshal(t)
	fmt.Println("producing message")
	fmt.Println("producing to" + topic)

	delivery_chan := make(chan kafka.Event, 10000)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          data,
	}, delivery_chan)
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

	// Wait for all messages to be delivered
	// p.Flush(15 * 1000)
}
