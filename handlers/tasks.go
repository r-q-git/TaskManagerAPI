package handlers

import (
	"TaskManager/data"
	"TaskManager/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TaskHandlers struct {
	Storage data.TaskRepositaryModel
}

func (th *TaskHandlers) Health(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode("Health Check"); err != nil {
		http.Error(w, "Not able to get the data json", http.StatusBadRequest)
		return
	}
}

func (th *TaskHandlers) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := th.Storage.GetTasks()
	if err != nil {
		log.Printf("tasks can't be retrieved: %v", err)
		http.Error(w, "No Task Found!", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "Not able to get the data json", http.StatusBadRequest)
		return
	}
}

func (th *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.CreateTask
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Not Able to Decode the Body", http.StatusBadRequest)
		return
	}
	var effectedRows = th.Storage.CreateTask(task)
	if effectedRows > 0 {
		w.Write([]byte("Task Added Successfully!"))
	} else {
		http.Error(w, "Task not added!", http.StatusBadRequest)
	}
}

func (th *TaskHandlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/task/"):]
	id_num, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id can't be found", http.StatusBadRequest)
		return
	}
	var task models.CreateTask
	w.Header().Set("Content-Type", "application/json")
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Not Able to Decode the Body", http.StatusBadRequest)
		return
	}

	updated_task, err := th.Storage.UpdateTask(id_num, task.Title, task.Description, task.Status)
	if err != nil {
		http.Error(w, "Not Able to Update the Task", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(updated_task); err != nil {
		http.Error(w, "Not Able to Decode the Updated Task", http.StatusBadRequest)
		return
	}
}
