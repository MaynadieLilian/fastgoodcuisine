package handlers

import (
	"fastgoodcuisine/internal"
	"fastgoodcuisine/internal/model"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/HomePage.html"))
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		time := r.FormValue("time")
		difficulty := r.FormValue("difficulty")
		ingredients := r.FormValue("ingredients")
		instructions := r.FormValue("instructions")
		recipe := model.Recipe{
			Title:        title,
			Time:         time,
			Difficulty:   difficulty,
			Ingredients:  ingredients,
			Instructions: instructions,
		}
		if err := internal.CreateRecipe(recipe); err != nil {
			log.Printf("Error creating recipe: %s", err)
			return
		}
		log.Println("Recipe created successfully")
		http.Redirect(w, r, "/homePage", http.StatusSeeOther)
		return
	}
	recipes, err := internal.GetRecipe()
	if err != nil {
		log.Printf("error getting recipes: %s", err)
		return
	}
	err = t.Execute(w, recipes)
	if err != nil {
		return
	}
}
