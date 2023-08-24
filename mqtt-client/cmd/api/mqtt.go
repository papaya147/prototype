package main

import (
	"fmt"
	"log"

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
	token := client.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

func Subscribe(client mqtt.Client) (mqtt.Token, error) {
	token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("received message: %s from topic: %s", msg.Payload(), msg.Topic())

	})
	token.Wait()
	if token.Error() != nil {
		return nil, token.Error()
	}
	return token, nil
}

func Unsubscribe(client mqtt.Client) {
	client.Unsubscribe(topic)
}

func Disconnect(client mqtt.Client, millis uint) {
	client.Disconnect(millis)
}
