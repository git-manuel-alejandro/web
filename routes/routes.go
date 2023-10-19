package routes

import (
	"encoding/json"
	"html/template"
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

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	data := map[string]string{
		"id": id,
	}

	template, err := template.ParseFiles("templates/users.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, data)

	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)

	}

}

func Nosotros(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/nosotros.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(w, nil)

	}

}
