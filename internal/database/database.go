package database

import (
	"database/sql"
	jsonformat "todo-list/internal/json_format"

	_ "github.com/lib/pq"
)

type Task struct {
	Id          int
	Task_Name   string
	Description string
}

type Insert_Task struct {
	Task_Name   string
	Description string
}

type Modify_Task struct {
	Id int
}

func ConnectDataBase(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func GetTasks(db *sql.DB) string {
	tasks := []Task{}

	rows, err := db.Query("SELECT * FROM Tasks")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		task := Task{}
		err := rows.Scan(
			&task.Id, &task.Task_Name, &task.Description,
		)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, task)
	}

	return jsonformat.FormatRawDataIntoJSON(tasks)

}

func AddTasks(task Insert_Task, db *sql.DB) sql.Result {
	response, err := db.Exec("INSERT INTO Tasks(Task_Name, Description) VALUES($1, $2)", task.Task_Name, task.Description)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return response

}

func DeleteTasks(id Modify_Task, db *sql.DB) sql.Result {
	response, err := db.Exec("DELETE FROM Tasks WHERE Id = $1", id.Id)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return response
}

func UpdateTasks(task Task, db *sql.DB) sql.Result {
	response, err := db.Exec("UPDATE Tasks SET Task_Name = $1, Description = $2 WHERE Id = $3", task.Task_Name, task.Description, task.Id)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return response
}
