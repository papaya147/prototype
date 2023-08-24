package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var broker string
var port string
var clientID string
var telemetryTopic string
var acknowledgeTopic string
var username string
var password string

func main() {
	// load environment variables from the .env file
	LoadEnv()

	producer, err := CreateProducer()
	if err != nil {
		log.Panic("producer creation failed")
	} else {
		log.Println("producer created")
	}
	defer producer.Close()

	client, err := Connect()
	if err != nil {
		log.Panicf("error connecting to MQTT broker: %s", err)
	}
	log.Println("connected to MQTT broker")

	messageChannel := make(chan JSONPayload)

	err = Subscribe(client, messageChannel)
	if err != nil {
		log.Panic("error subscribing to MQTT topic")
	}
	defer Unsubscribe(client)
	defer Disconnect(client, 250)

	var msgBytes []byte

	for {
		select {
		case msgJSON := <-messageChannel:
			msgBytes, _ = json.Marshal(msgJSON)
			handleMessageReceived(msgBytes, producer, client)
		case <-time.After(time.Second * 5): // Wait for 5 seconds if no message is received
			fmt.Println("MQTT, telemetry: no message received in the last 5 seconds")
		}
	}
}

func handleMessageReceived(msgString []byte, producer sarama.AsyncProducer, client mqtt.Client) {
	// get message and display it
	log.Printf("MQTT, telemetry: received message: %s\n", string(msgString))

	// push message to kafka
	ProduceMessage(producer, "mqtt-sink", string(msgString))

	// send ack to MQTT
	err := Publish(client, msgString)
	if err != nil {
		log.Panicln("MQTT, ack      : error publishing MQTT message:", err)
	} else {
		log.Println("MQTT, ack      : published message", string(msgString))
	}
}
