package main

import (
	"log"
	"net/http"

	"github.com/cheatsnake/rest-api-vanilla-go/internal/server"
	"github.com/cheatsnake/rest-api-vanilla-go/internal/taskstore"
)

func main() {
	server := server.New(taskstore.New())

	http.HandleFunc("/task/", server.TaskHandler)
	http.HandleFunc("/tasks/", server.TasksHandler)

	log.Fatal(http.ListenAndServe(":5005", nil))
}
