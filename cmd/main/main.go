package main

import (
	"fmt"
	"net/http"
	"todo-list/internal/router"
)

func main() {

	router.Router("frontend")

	fmt.Println("Server is started...\nGo to http://localhost:2222")
	err := http.ListenAndServe("localhost:2222", nil)
	if err != nil {
		fmt.Println(err)
	}
}
