package server

import "github.com/cheatsnake/rest-api-vanilla-go/internal/taskstore"

type Server struct {
	Store *taskstore.TaskStore
}

func New(s *taskstore.TaskStore) *Server {
	newServer := Server{s}
	return &newServer
}
