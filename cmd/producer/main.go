package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer := NewKafkaProducer()
	Publish("Oi Luciano Romão", "teste", producer, []byte(nil))
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{"bootstrap.servers": "fc_kafka-kafka-1:9092"}

	p, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Println(err.Error())
	}

	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: topic, kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(message)

	if err != nil {
		return err
	}

	return nil
}
