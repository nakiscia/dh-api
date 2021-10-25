package server

import (
	"dh-api/core/repository"
	"dh-api/core/service"
	"fmt"
	"net/http"
)

type Server struct {
	Config *Config
}

func NewServer(config *Config) Server{
	return Server{Config: config}
}

func (s *Server) String() string {
	return fmt.Sprintf("%s:%s", s.Config.Hostname, s.Config.Port)
}

func (s *Server) init(){
	fileRepository := repository.NewFileRepository(s.Config.PersistenceFilePath)
	keyService := service.NewKeyService(fileRepository,s.Config.PersistenceWriteDuration)

	handler := NewDefaultHandler(keyService)
	handler.HandleRequests()
}

func (s *Server) Run(){
	s.init()
	fmt.Printf(fmt.Sprintf("Server started at %s", s.String()))
	err := http.ListenAndServe(s.String(), nil)
	if err != nil{
		panic("Error while running the server..")
	}
}