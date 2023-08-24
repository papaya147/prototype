package main

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Config() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%s", broker, port))
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	return opts
}

func Connect() (mqtt.Client, error) {
	opts := Config()
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}

func Publish(client mqtt.Client, message []byte) error {
	token := client.Publish(acknowledgeTopic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

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

var jsonPayload JSONPayload

func Subscribe(client mqtt.Client, messageChannel chan JSONPayload) error {
	if token := client.Subscribe(telemetryTopic, 0, func(client mqtt.Client, msg mqtt.Message) {
		json.Unmarshal(msg.Payload(), &jsonPayload)
		messageChannel <- jsonPayload
	}); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func Unsubscribe(client mqtt.Client) {
	client.Unsubscribe(telemetryTopic)
}

func Disconnect(client mqtt.Client, millis uint) {
	client.Disconnect(millis)
}
