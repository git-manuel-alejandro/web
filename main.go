package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Txt string `json:"txt"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Txt: "hello world from golang !!!"})

}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	json.NewEncoder(w).Encode(Message{Txt: id})

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/users/{id}", UserHandler)
	http.Handle("/", r)

	fmt.Println("server running ...")

	http.ListenAndServe(":3000", r)
}
