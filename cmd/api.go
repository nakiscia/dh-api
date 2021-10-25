package main

import (
	"dh-api/server"
	"time"
)

func main() {

	s:= server.NewServer(&server.Config{
		Hostname:            "",
		Port:                "8181",
		PersistenceFilePath: "./key-record.txt",
		PersistenceWriteDuration: time.Duration(1000),
	})

	s.Run()
}