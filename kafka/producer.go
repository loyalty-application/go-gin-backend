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

  key := t.CardId
  data,_ := json.Marshal(t)
  
  delivery_chan := make(chan kafka.Event, 10000)
  p.Produce(&kafka.Message{
    TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
    Key:            []byte(key),
    Value:          data,
  },delivery_chan)

  e := <-delivery_chan
   m := e.(*kafka.Message)
  if m.TopicPartition.Error != nil {
    fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
  } else {
    fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
        *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
  }
  close(delivery_chan)
}
