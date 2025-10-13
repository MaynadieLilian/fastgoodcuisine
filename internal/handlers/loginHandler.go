package handlers

import (
	"fastgoodcuisine/internal"
	"log"
	"net/http"
	"os"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "web/templates/LoginPage.html")
	case http.MethodPost:
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := internal.Login(email, password)
		if err != nil {
			if os.IsNotExist(err) {
				http.Error(w, "No account found with this email", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Incorrect email or password", http.StatusUnauthorized)
			return
		}

		log.Printf("connected : %s", user.Email)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
