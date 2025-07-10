package main

import (
	"fmt"
	"net/http"

	"CacheCluster/controller"
)

func main() {
	http.HandleFunc("/", controller.CallOperation)

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
