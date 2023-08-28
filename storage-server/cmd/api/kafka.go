package main

import (
	"crypto/tls"

	"github.com/IBM/sarama"
)

func createTLSConfig() *tls.Config {
	tlsConfig := &tls.Config{}
	return tlsConfig
}

func CreateConsumer() (sarama.Consumer, error) {
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

	consumer, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func CreatePartitionConsumer(consumer sarama.Consumer) (sarama.PartitionConsumer, error) {
	partitionConsumer, err := consumer.ConsumePartition("mqtt-sink-topic", int32(partition), sarama.OffsetNewest)
	if err != nil {
		return nil, err
	}
	return partitionConsumer, nil
}
