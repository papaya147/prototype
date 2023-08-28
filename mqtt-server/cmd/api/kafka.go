package main

import (
	"crypto/tls"
	"log"

	"github.com/IBM/sarama"
)

func createTLSConfig() *tls.Config {
	tlsConfig := &tls.Config{}
	return tlsConfig
}

func CreateProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	config.Net.SASL.Enable = true
	config.Net.SASL.User = "ETGQLH5KJD2KBETH"
	config.Net.SASL.Password = "0Bb2Jmigq4CcJwY1iTNE+JqjMrv10GKHOiDHI22bbyBEa5MbU3bVVfZeGPkbDb6K"
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	config.Net.TLS.Enable = true
	config.Net.TLS.Config = createTLSConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll

	brokerList := []string{"pkc-9q8rv.ap-south-2.aws.confluent.cloud:9092"}

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func ProduceMessage(producer sarama.AsyncProducer, topic string, key string, value string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	producer.Input() <- msg

	go func() {
		for {
			select {
			case err := <-producer.Errors():
				log.Println("Kafka          : failed to produce message:", err.Err)
			case success := <-producer.Successes():
				log.Println("Kafka          : message produced successfully on partition", success.Partition)
			}
		}
	}()
}
