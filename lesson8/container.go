package main

import (
	"fmt"

	"go.uber.org/dig"
)

type Database struct {
	Name string
}

func NewDatabase() *Database {
	return &Database{
		Name: "MyDatabase",
	}
}

type Server struct {
	Database *Database
}

func NewServer(db *Database) *Server {
	return &Server{
		Database: db,
	}
}

func (s *Server) Start() {
	fmt.Printf("Server is using database: %s\n", s.Database.Name)
}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewDatabase)
	container.Provide(NewServer)

	return container
}
