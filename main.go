package main

import (
	"html/template"
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sirsean/pad/config"
	"github.com/sirsean/pad/api"
)

func main() {
	log.Printf("Starting Pad")

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/pad", api.Create).Methods("POST")
	router.HandleFunc("/pad/{id}", api.Consume).Methods("GET")

	http.Handle("/", router)

	port := config.Get().Host.Port
	http.ListenAndServe(port, nil)
}

var indexTemplate = template.Must(template.ParseFiles(fmt.Sprintf("%s/template/index.html", config.Get().Host.Path)))

func index(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, nil)
}
