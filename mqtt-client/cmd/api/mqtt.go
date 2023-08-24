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
	token := client.Publish(telemetryTopic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

var jsonPayload JSONPayload

func Subscribe(client mqtt.Client, messageChannel chan JSONPayload) error {
	token := client.Subscribe(acknowledgeTopic, 0, func(client mqtt.Client, msg mqtt.Message) {
		json.Unmarshal(msg.Payload(), &jsonPayload)
		messageChannel <- jsonPayload
	})
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

func Unsubscribe(client mqtt.Client) {
	client.Unsubscribe(acknowledgeTopic)
}

func Disconnect(client mqtt.Client, millis uint) {
	client.Disconnect(millis)
}
