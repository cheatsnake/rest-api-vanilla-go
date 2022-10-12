package taskstore

import (
	"errors"
	"sync"
	"time"
)
type TaskStore struct {
	sync.Mutex

	Tasks  map[int]Task
	NextId int
}

type Task struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	Deadline  int      `json:"deadline"`
	CreatedAt int      `json:"createdAt"`
}

var Store *TaskStore

func New() *TaskStore {
	ts := &TaskStore{}
	ts.Tasks = make(map[int]Task)
	ts.NextId = 1
	return ts
}

func init() {
	Store = New()
}

func (ts *TaskStore) CreateTask(name, body string, tags []string, deadline int) Task {
	ts.Lock()
	defer ts.Unlock()

	newTask := Task{
		Id:        ts.NextId,
		Name:      name,
		Body:      body,
		Deadline:  deadline,
		CreatedAt: int(time.Now().Unix()),
	}
	newTask.Tags = make([]string, len(tags))
	copy(newTask.Tags, tags)

	return newTask
}

func (ts *TaskStore) GetTaskById(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	task, ok := ts.Tasks[id]
	if !ok {
		return Task{}, errors.New("task not found")
	}

	return task, nil
}

func (ts *TaskStore) GetAllTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	tasks := make([]Task, 0, len(ts.Tasks))
	for _, task := range ts.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (ts *TaskStore) UpdateTaskById(id int, name, body string, tags []string, deadline int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	_, err := ts.GetTaskById(id)
	if err != nil {
		return Task{}, nil
	}

	updatedTask := Task{
		Id:        id,
		Name:      name,
		Body:      body,
		Deadline:  deadline,
		CreatedAt: ts.Tasks[id].CreatedAt,
	}

	updatedTask.Tags = make([]string, len(tags))
	copy(updatedTask.Tags, tags)

	ts.Tasks[id] = updatedTask

	return ts.Tasks[id], nil
}

func (ts *TaskStore) DeleteTaskById(id int) error {
	ts.Lock()
	defer ts.Unlock()

	_, ok := ts.Tasks[id]
	if !ok {
		return errors.New("task not found")
	}

	delete(ts.Tasks, id)
	return nil
}

func (ts *TaskStore) DeleteAllTasks() {
	ts.Lock()
	defer ts.Unlock()

	ts.Tasks = make(map[int]Task)
}
