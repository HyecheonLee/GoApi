package main

import (
	"GoApi/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/users/", handlers.UserRouter)
	http.HandleFunc("/users", handlers.UserRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
