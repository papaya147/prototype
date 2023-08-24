package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load .env file")
	}

	broker = os.Getenv("HIVEMQ_BROKER_URL")
	port = os.Getenv("HIVEMQ_BROKER_PORT")
	clientID = os.Getenv("MQTT_CLIENTID")
	telemetryTopic = os.Getenv("MQTT_TELEMETRY_TOPIC")
	acknowledgeTopic = os.Getenv("MQTT_ACK_TOPIC")
	username = os.Getenv("HIVEMQ_USERNAME")
	password = os.Getenv("HIVEMQ_PASSWORD")

	return nil
}
