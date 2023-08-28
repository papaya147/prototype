package main

import (
	"github.com/gocql/gocql"
	"github.com/gocql/gocql/scyllacloud"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
)

const (
	connectionBundlePath = "./connect-bundle-mqtt-storage-test.yml"
)

var stmts *statements = createStatements()

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

func CreateScyllaKeyspace(session *gocql.Session) error {
	query := `
		CREATE KEYSPACE IF NOT EXISTS mqtt_storage 
		WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};
	`
	return session.Query(query).Exec()
}

func CreateScyllaTable(session *gocql.Session) error {
	query := `
        CREATE TABLE IF NOT EXISTS mqtt_storage.telemetry (
            time BIGINT,
            battery_temp INT,
            speed INT,
            latitude DOUBLE,
            longitude DOUBLE,
            PRIMARY KEY (time)
        )
    `
	return session.Query(query).Exec()
}

type query struct {
	stmt  string
	names []string
}

type statements struct {
	del query
	ins query
	sel query
}

type Telemetry struct {
	Time        int64   `db:"time"`
	BatteryTemp int     `db:"battery_temp"`
	Speed       int     `db:"speed"`
	Latitude    float64 `db:"latitude"`
	Longitude   float64 `db:"longitude"`
}

func createStatements() *statements {
	metadata := table.Metadata{
		Name:    "mqtt_storage.telemetry",
		Columns: []string{"time", "battery_temp", "speed", "latitude", "longitude"},
		PartKey: []string{"time"},
	}
	tbl := table.New(metadata)

	deleteStmt, deleteNames := tbl.Delete()
	insertStmt, insertNames := tbl.Insert()

	// Normally a select statement such as this would use `tbl.Select()` to select by
	// primary key but now we just want to display all the records...
	selectStmt, selectNames := qb.Select(metadata.Name).Columns(metadata.Columns...).ToCql()

	return &statements{
		del: query{
			stmt:  deleteStmt,
			names: deleteNames,
		},
		ins: query{
			stmt:  insertStmt,
			names: insertNames,
		},
		sel: query{
			stmt:  selectStmt,
			names: selectNames,
		},
	}
}

func InsertQuery(session *gocql.Session, telemetry *Telemetry) error {
	err := gocqlx.Query(session.Query(stmts.ins.stmt), stmts.ins.names).BindStruct(telemetry).ExecRelease()
	return err
}
