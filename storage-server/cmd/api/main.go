package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
	"github.com/gocql/gocql"
)

var partition int

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
	LoadEnv()

	consumer := createKafkaConsumer()
	defer consumer.Close()

	partitionConsumer := createKafkaPartitionConsumer(consumer)
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	session := createScyllaSession()
	defer session.Close()

	err := CreateScyllaKeyspace(session)
	if err != nil {
		log.Panicln("error creating keyspace: ", err)
	} else {
		log.Println("keyspace created successfully!")
	}

	err = CreateScyllaTable(session)
	if err != nil {
		log.Panicln("error creating table: ", err)
	} else {
		log.Println("table created successfully!")
	}

	var jsonPayload JSONPayload
	var telemetry Telemetry

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("received message: %s\n", string(msg.Value))
			json.Unmarshal(msg.Value, &jsonPayload)

			telemetry = Telemetry{
				Time:        jsonPayload.Time,
				BatteryTemp: jsonPayload.Data.BatteryTemp,
				Speed:       jsonPayload.Data.Speed,
				Latitude:    jsonPayload.Data.Latitude,
				Longitude:   jsonPayload.Data.Longitude,
			}

			err := InsertQuery(session, &telemetry)
			if err != nil {
				log.Println("error inserting data: ", err)
			} else {
				log.Println("data inserted successfully!")
			}
		case <-signals:
			log.Println("interrupt signal received, stopping consumer...")
		}
	}
}

var counts = 0

func createScyllaSession() *gocql.Session {
	for {
		session, err := CreateScyllaSession()
		if err != nil {
			log.Println("scylla session not yet ready...")
			counts++
		} else {
			log.Println("scylla session connected to ScyllaDB!")

			return session
		}

		if counts > 5 {
			log.Println(err)
		}

		log.Println("backing off for 10 seconds...")
		time.Sleep(10 * time.Second)
		continue
	}
}

func createKafkaPartitionConsumer(consumer sarama.Consumer) sarama.PartitionConsumer {
	for {
		consumer, err := CreatePartitionConsumer(consumer)
		if err != nil {
			log.Println("partition consumer not yet ready...")
			counts++
		} else {
			log.Println("partition consumer connected to Kafka!")
			return consumer
		}

		if counts > 5 {
			log.Println(err)
		}

		log.Println("backing off for 2 seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func createKafkaConsumer() sarama.Consumer {
	for {
		consumer, err := CreateConsumer()
		if err != nil {
			log.Println("Kafka not yet ready...")
			counts++
		} else {
			log.Println("connected to Kafka!")
			return consumer
		}

		if counts > 5 {
			log.Println(err)
		}

		log.Println("backing off for 2 seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
