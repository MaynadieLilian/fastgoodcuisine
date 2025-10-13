package server

import (
	"fastgoodcuisine/internal/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RegisterHandler)
	r.HandleFunc("/login", handlers.LoginHandler)

	fs := http.FileServer(http.Dir("./web/styles"))
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", fs))
	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("Server error: %s", err)
	}
}
