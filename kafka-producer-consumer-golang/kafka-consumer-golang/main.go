package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "golang",   // consumer-group
		"auto.offset.reset": "earliest", // from-begining
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	topic_name := "helloworld"
	err2 := consumer.Subscribe(topic_name, nil)
	if err2 != nil {
		panic(err2)
	}

	for {
		m, err3 := consumer.ReadMessage(1 * time.Second)
		if err3 == nil {
			fmt.Printf("Receive message: %s\n", m.Value)
		}
	}
}
