package main

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	}

	p, err := kafka.NewProducer(config)

	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := "helloworld"

	for i := 0; i < 10; i++ {
		fmt.Printf("send message to kafka %d \n", i)
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(fmt.Sprintf("Hello %d", i)),
		}

		err = p.Produce(msg, nil)
		if err != nil {
			panic(err)
		}
	}

	p.Flush(5 * 1000)
}
