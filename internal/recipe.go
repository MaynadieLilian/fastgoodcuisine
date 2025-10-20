package internal

import (
	"encoding/json"
	"fastgoodcuisine/internal/model"
	"log"
	"os"
	"sync"
)

var muRecipe sync.Mutex

func CreateRecipe(recipe model.Recipe) error {
	muRecipe.Lock()
	defer muRecipe.Unlock()
	recipes, err := GetRecipe()
	if err != nil {
		return err
	}
	recipes = append(recipes, recipe)
	file, err := os.Create("recipe.json")
	if err != nil {
		log.Printf("Error creating recipe file: %s", err)
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(recipes)
	if err != nil {
		log.Printf("Error encoding file : %s", err)
		return err
	}
	return nil
}

func GetRecipe() ([]model.Recipe, error) {
	var recipes []model.Recipe
	mu.Lock()
	defer mu.Unlock()
	if fileData, err := os.ReadFile("recipe.json"); err == nil {
		err = json.Unmarshal(fileData, &recipes)
		if err != nil {
			return nil, err
		}
	}
	return recipes, nil
}
