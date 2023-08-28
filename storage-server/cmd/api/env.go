package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load .env file")
	}

	partition, _ = strconv.Atoi(os.Getenv("PARTITION"))
}
