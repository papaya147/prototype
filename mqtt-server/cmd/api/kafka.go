package main

import (
	"log"

	"github.com/IBM/sarama"
)

func CreateProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	brokerList := []string{"kafka:9092"}

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func ProduceMessage(producer sarama.AsyncProducer, topic string, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	producer.Input() <- msg

	go func() {
		for {
			select {
			case err := <-producer.Errors():
				log.Println("failed to produce message:", err.Err)
			case success := <-producer.Successes():
				log.Println("message sent successfully:", success.Offset)
			}
		}
	}()
}
