package main

import (
	"log"
	"net/http"

	"github.com/cheatsnake/rest-api-vanilla-go/internal/handlers"
	"github.com/cheatsnake/rest-api-vanilla-go/internal/taskstore"
)

func main() {
	server := handlers.Server{Store: taskstore.New()}

	http.HandleFunc("/task/", server.TaskHandler)
	http.HandleFunc("/tasks/", server.TasksHandler)

	log.Fatal(http.ListenAndServe(":5005", nil))
}
