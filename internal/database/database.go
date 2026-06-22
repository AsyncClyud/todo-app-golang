package database

import (
	"database/sql"
	jsonformat "todo-list/internal/json_format"

	_ "github.com/lib/pq"
)

type Task struct {
	Task_Name   string
	Description string
}

type GetTask struct {
	Id int
	Task
}

func ConnectDataBase(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func GetTasks(db *sql.DB) string {
	tasks := []GetTask{}

	rows, err := db.Query("SELECT * FROM Tasks WHERE Id < 11")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		task := GetTask{}
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

func AddTasks(task Task, db *sql.DB) sql.Result {

	response, err := db.Exec("INSERT INTO Tasks(Task_Name, Description) VALUES($1, $2)", task.Task_Name, task.Description)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return response

}
