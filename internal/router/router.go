package router

import (
	"net/http"
	"todo-list/internal/handler"
)

func Router(Frontend_path string) {

	Fs := http.FileServer(http.Dir(Frontend_path))

	http.Handle("/", Fs)

	http.HandleFunc("/get_tasks", func(w http.ResponseWriter, r *http.Request) {
		handler.GetTaskHandler(w, *r)
	})

	http.HandleFunc("/add_tasks", func(w http.ResponseWriter, r *http.Request) {
		handler.AddTaskHandler(w, *r)
	})

}
