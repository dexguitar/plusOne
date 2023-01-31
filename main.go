package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello there!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
