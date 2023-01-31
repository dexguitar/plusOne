package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello there!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", homeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
