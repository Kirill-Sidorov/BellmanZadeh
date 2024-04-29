package main

import "bellmanzadeh/data"

func main() {

	data.ParseJsonData()

	/*
			server := http.NewServeMux()
			server.HandleFunc("/", baseHandler)
			server.HandleFunc("/favicon.ico", faviconHandler)

			log.Println("Server Started...")
			err := http.ListenAndServe("localhost:8080", server)
			if err != nil {
				log.Fatal(err)
			}

			func baseHandler(writer http.ResponseWriter, request *http.Request) {

		}

		func faviconHandler(writer http.ResponseWriter, request *http.Request) {
			http.ServeFile(writer, request, "resources/favicon.ico")
		}
	*/
}
