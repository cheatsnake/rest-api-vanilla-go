package handlers

import (
	"encoding/json"
	"github.com/cheatsnake/rest-api-vanilla-go/helpers"
	"github.com/cheatsnake/rest-api-vanilla-go/models/taskstore"
	"net/http"
	"strconv"
	"strings"
)

type taskBody struct {
	Name      string   `json:"name"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	Deadline  int      `json:"deadline"`
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		taskId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
		if err != nil {
			helpers.HandleError(w, http.StatusBadRequest, "error parsing task id")
			return
		}

		task, err := taskstore.Store.GetTaskById(taskId)
		if err != nil {
			helpers.HandleError(w, http.StatusNotFound, "task not found")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(task)
		w.Write(response)
	}

	if r.Method == http.MethodPost {
		var task taskBody

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			helpers.HandleError(w, http.StatusBadRequest, err.Error())
			return
		}

		newTask := taskstore.Store.CreateTask(task.Name, task.Body, task.Tags, task.Deadline)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(newTask)
		w.Write(response)
	}
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tasks := taskstore.Store.GetAllTasks()

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(tasks)
		w.Write(response)
	}
}
