package main

import (
	"log"
	"net/http"
	hd "post/internals/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := "8082"
	http.HandleFunc("/", hd.MainHandler)
	log.Printf("Server (portAPI) started at http://localhost:%v\n", port)
	http.ListenAndServe(":"+port, nil)
}
