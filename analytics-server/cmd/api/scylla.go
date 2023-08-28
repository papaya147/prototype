package main

import (
	"github.com/gocql/gocql"
	"github.com/gocql/gocql/scyllacloud"
)

const (
	connectionBundlePath = "./connect-bundle-mqtt-storage-test.yml"
)

func CreateScyllaSession() (*gocql.Session, error) {
	cluster, err := scyllacloud.NewCloudCluster(connectionBundlePath)
	if err != nil {
		return nil, err
	}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("us-east-1")

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

type Telemetry struct {
	Time        int64   `db:"time"`
	BatteryTemp int     `db:"battery_temp"`
	Speed       int     `db:"speed"`
	Latitude    float64 `db:"latitude"`
	Longitude   float64 `db:"longitude"`
}

func SelectQuery(session *gocql.Session) ([]Telemetry, error) {
	query := "SELECT * FROM mqtt_storage.telemetry"
	iter := session.Query(query).Iter()

	var result []Telemetry
	var telemetry Telemetry

	for iter.Scan(&telemetry.Time, &telemetry.BatteryTemp, &telemetry.Speed, &telemetry.Latitude, &telemetry.Longitude) {
		result = append(result, telemetry)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return result, nil
}
