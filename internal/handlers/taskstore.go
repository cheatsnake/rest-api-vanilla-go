package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/cheatsnake/rest-api-vanilla-go/internal/taskstore"
	helpers "github.com/cheatsnake/rest-api-vanilla-go/tools"
)

type Server struct {
	Store *taskstore.TaskStore
}

type taskBody struct {
	Name     string   `json:"name"`
	Body     string   `json:"body"`
	Tags     []string `json:"tags"`
	Deadline int      `json:"deadline"`
}

func (s *Server) TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		taskId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
		if err != nil {
			helpers.HandleError(w, http.StatusBadRequest, "error parsing task id")
			return
		}

		task, err := s.Store.GetTaskById(taskId)
		if err != nil {
			helpers.HandleError(w, http.StatusNotFound, "task not found")
			return
		}

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

		newTask := s.Store.CreateTask(task.Name, task.Body, task.Tags, task.Deadline)
		response, _ := json.Marshal(newTask)
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	}

	if r.Method == http.MethodPut {
		var task taskBody

		taskId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
		if err != nil {
			helpers.HandleError(w, http.StatusBadRequest, "error parsing task id")
			return
		}

		err = json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			helpers.HandleError(w, http.StatusBadRequest, "failed parsing body")
			return
		}

		updatedTask, err := s.Store.UpdateTaskById(taskId, task.Name, task.Body, task.Tags, task.Deadline)
		if err != nil {
			helpers.HandleError(w, http.StatusNotFound, err.Error())
			return
		}

		response, _ := json.Marshal(updatedTask)
		w.Write(response)
	}

	if r.Method == http.MethodDelete {
		taskId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
		if err != nil {
			helpers.HandleError(w, http.StatusBadRequest, "error parsing task id")
			return
		}

		err = s.Store.DeleteTaskById(taskId)
		if err != nil {
			helpers.HandleError(w, http.StatusNotFound, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		tasks := s.Store.GetAllTasks()

		response, _ := json.Marshal(tasks)
		w.Write(response)
	}

	if r.Method == http.MethodDelete {
		s.Store.DeleteAllTasks()
		w.WriteHeader(http.StatusOK)
	}
}
