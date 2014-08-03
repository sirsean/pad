package api

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sirsean/pad/model"
)

var store model.Store

func init() {
	store = model.NewStore()
}

func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p model.Pad
	decoder.Decode(&p)

	store.Add(&p)
	log.Printf("Added: %v", p)

	response, _ := json.Marshal(p)
	w.Write(response)
}

func Consume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("Consuming pad: %v", id)

	p, ok := store.Consume(id)
	log.Printf("got pad: %v", p)

	if ok {
		// redirect to the callback
		url := p.CallbackUrl()
		log.Printf("Redirecting to: %v", url)
		http.Redirect(w, r, url, 301)
	} else {
		// pad not found
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("NO PAD"))
	}
}
