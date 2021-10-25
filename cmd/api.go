package main

import (
	"dh-api/server"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}
	s := server.NewServer(&server.Config{
		Hostname:                 "",
		Port:                     port,
		PersistenceFilePath:      "./key-record.txt",
		PersistenceWriteDuration: time.Duration(1000),
	})

	s.Run()
}
