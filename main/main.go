package main

import (
	"bellmanzadeh/customrender"
	"bellmanzadeh/handler"
	"log"
	"net/http"
)

func main() {

	handler.Init()
	customrender.Init()

	server := http.NewServeMux()
	server.HandleFunc("/", handler.MainPageHandler)
	server.HandleFunc("/solveTask", handler.ResultHandler)
	server.HandleFunc("/favicon.ico", faviconHandler)

	log.Println("Server Started...")
	err := http.ListenAndServe("localhost:8080", server)
	if err != nil {
		log.Fatal(err)
	}
}

func faviconHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "resources/favicon.ico")
}
