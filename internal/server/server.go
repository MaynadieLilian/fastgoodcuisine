package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
