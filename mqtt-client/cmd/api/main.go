package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var broker string
var port string
var clientID string
var telemetryTopic string
var acknowledgeTopic string
var username string
var password string

type JSONPayload struct {
	Error        bool   `json:"error,omitempty"`
	ErrorMessage string `json:"message,omitempty"`
	Data         Data   `json:"data,omitempty"`
	Time         int64  `json:"time,omitempty"`
}

type Data struct {
	BatteryTemp int     `json:"battery,omitempty"`
	Speed       int     `json:"speed,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
}

func main() {
	// load environment variables from the .env file
	LoadEnv()

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

	go publishOnMain(client)

	var msgBytes []byte

	for {
		select {
		case msgJSON := <-messageChannel:
			msgBytes, _ = json.Marshal(msgJSON)
			log.Println("MQTT, ack      : message received: ", string(msgBytes))
		case <-time.After(time.Second * 5): // Wait for 5 seconds if no message is received
			log.Println("MQTT, ack      : no message received in the last 5 seconds")
		}
	}
}

func publishOnMain(client mqtt.Client) {
	for {
		jsonPayload := JSONPayload{
			Error:        false,
			ErrorMessage: "",
			Data: Data{
				BatteryTemp: rand.Intn(11) + 15,
				Speed:       rand.Intn(11) + 20,
				Latitude:    rand.Float64() * 10,
				Longitude:   rand.Float64() * 10,
			},
			Time: time.Now().Unix(),
		}

		msg, _ := json.Marshal(jsonPayload)
		err := Publish(client, msg)
		if err != nil {
			log.Printf("MQTT, telemetry: error publishing message: %s", err)
		} else {
			log.Printf("MQTT, telemetry:published message: %s\n", msg)
		}
		time.Sleep(time.Second)
	}
}
