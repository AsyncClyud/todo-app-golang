package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"todo-list/internal/database"
)

type Task struct {
	Task_Name   string
	Description string
}

func GetTaskHandler(w http.ResponseWriter, r http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDataBase("user=myuser password=Hydra1234 dbname=cleardatabase sslmode=disable")
	defer db.Close()

	tasks := database.GetTasks(db)

	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		panic(err)
	}

}

func AddTaskHandler(w http.ResponseWriter, r http.Request) sql.Result {

	var task database.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		panic(err)
	}

	db := database.ConnectDataBase("user=myuser password=Hydra1234 dbname=cleardatabase sslmode=disable")
	defer db.Close()

	response := database.AddTasks(task, db)

	return response

}
