package server

import "time"

type Config struct {
	Hostname string
	Port                string
	PersistenceFilePath string
	PersistenceWriteDuration time.Duration
}