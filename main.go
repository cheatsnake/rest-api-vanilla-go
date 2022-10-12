package main

import (
	"log"
	"net/http"

	"github.com/cheatsnake/rest-api-vanilla-go/handlers"
	"github.com/cheatsnake/rest-api-vanilla-go/models/taskstore"
)

var Store *taskstore.TaskStore

func init() {
	Store = taskstore.New()
}

func main() {
	http.HandleFunc("/task/", handlers.TaskHandler)
	http.HandleFunc("/tasks/", handlers.TasksHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
