package main

import (
	"fmt"
	"net/http"
	"web/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users/{id}", routes.UsersHandler)
	r.HandleFunc("/home", routes.Home)
	r.HandleFunc("/nosotros", routes.Nosotros)
	http.Handle("/", r)

	fmt.Println("server running ...")

	http.ListenAndServe(":3000", r)
}
