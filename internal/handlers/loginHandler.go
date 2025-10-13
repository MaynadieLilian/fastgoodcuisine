package handlers

import (
	"fastgoodcuisine/internal"
	"html/template"
	"log"
	"net/http"
	"os"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/LoginPage.html"))
	if r.Method != http.MethodPost {
		err := t.Execute(w, nil)
		if err != nil {
			log.Printf("template execution error: %s", err)
		}
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := internal.Login(email, password)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusUnauthorized)
			log.Printf("No account found with this email: %s", email)
			data := struct {
				Error string
			}{
				Error: "No account found with this email",
			}
			err = t.Execute(w, data)
			if err != nil {
				log.Printf("template execution error: %s", err)
			}
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("Incorrect email or password %s", err)
		data := struct {
			Error string
			Email string
		}{
			Error: "Incorrect email or password",
			Email: email,
		}
		err = t.Execute(w, data)
		if err != nil {
			log.Printf("template execution error: %s", err)
		}
		return
	}
	log.Printf("connected : %s", user.Email)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
