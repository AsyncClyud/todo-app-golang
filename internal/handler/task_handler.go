package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"todo-list/internal/database"
)

func GetTasksHandler(w http.ResponseWriter, r http.Request) {

	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDataBase("user=myuser password=Hydra1234 dbname=cleardatabase sslmode=disable")
	defer db.Close()

	tasks := database.GetTasks(db)

	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		panic(err)
	}

}

func AddTasksHandler(w http.ResponseWriter, r http.Request) sql.Result {

	var task database.Insert_Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		panic(err)
	}

	db := database.ConnectDataBase("user=myuser password=Hydra1234 dbname=cleardatabase sslmode=disable")
	defer db.Close()

	response := database.AddTasks(task, db)
	return response

}

func DeleteTasksHandler(w http.ResponseWriter, r http.Request) sql.Result {
	var id database.Modify_Task
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		panic(err)
	}

	db := database.ConnectDataBase("user=myuser password=Hydra1234 dbname=cleardatabase sslmode=disable")
	defer db.Close()

	response := database.DeleteTasks(id, db)
	return response
}

func UpdateTasksHandler(w http.ResponseWriter, r http.Request) sql.Result {
	var task database.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		panic(err)
	}

	db := database.ConnectDataBase("user=myuser password=Hydra1234 dbname=cleardatabase sslmode=disable")
	defer db.Close()

	response := database.UpdateTasks(task, db)
	return response
}
