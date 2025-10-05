package handlers

import (
	"fastgoodcuisine/internal"
	"fastgoodcuisine/internal/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/RegisterPage.html"))
	if r.Method != http.MethodPost {
		err := t.Execute(w, nil)
		if err != nil {
			log.Printf("template execution error: %s", err)
		}
		return
	}
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password_confirm")
	if passwordConfirm != password {
		data := struct {
			Error    string
			Email    string
			Username string
		}{
			Error:    "Les mots de passe ne correspondent pas",
			Email:    email,
			Username: username,
		}
		err := t.Execute(w, data)
		if err != nil {
			log.Printf("template execution error: %s", err)
		}
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to generate password hash: %s", err)
	}
	user := model.User{
		Email:          email,
		Username:       username,
		HashedPassword: hashedPassword,
	}
	err = internal.Register(user)
	if err != nil {
		fmt.Printf("error while registering user: %s", err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
