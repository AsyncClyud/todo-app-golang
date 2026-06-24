package router

import (
	"net/http"
	"todo-list/internal/handler"
)

func Router(Frontend_path string) {

	Fs := http.FileServer(http.Dir(Frontend_path))

	http.Handle("/", Fs)

	http.HandleFunc("/get_tasks", func(w http.ResponseWriter, r *http.Request) {
		handler.GetTasksHandler(w, *r)
	})

	http.HandleFunc("/add_task", func(w http.ResponseWriter, r *http.Request) {
		handler.AddTasksHandler(w, *r)
	})

	http.HandleFunc("/delete_task", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteTasksHandler(w, *r)
	})

	http.HandleFunc("/update_task", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateTasksHandler(w, *r)
	})

}
