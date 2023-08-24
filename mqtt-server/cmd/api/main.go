package main

import (
	"log"
	"os"
	"os/signal"

	"syscall"
)

var broker string
var port string
var clientID string
var topic string
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

	_, err = Subscribe(client, producer)
	if err != nil {
		log.Panicf("error subscribing to MQTT topic: %s", err)
	}
	log.Println("subscribed to MQTT topic")

	// Setup signal handling to gracefully exit
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Println("exiting...")

	// Unsubscribe and disconnect
	Unsubscribe(client)
	log.Println("unsubscribed from MQTT topic")

	Disconnect(client, 250)
	log.Println("disconnected from MQTT broker")
}
