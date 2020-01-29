package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", ServeJobs)

	log.Fatal(http.ListenAndServe(":8000", router))
}
