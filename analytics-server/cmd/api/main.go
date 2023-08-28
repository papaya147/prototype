package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gocql/gocql"
)

var port = "8080"

type Config struct {
	Session *gocql.Session
}

func main() {
	session := createScyllaSession()
	defer session.Close()

	app := Config{
		Session: session,
	}

	log.Println("starting the analytics server...")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	} else {
		log.Print("server started on port: ", port)
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
