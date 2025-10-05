package internal

import (
	"encoding/json"
	"fastgoodcuisine/internal/model"
	"os"
	"sync"
)

var mu sync.Mutex

func Register(user model.User) error {
	var users []model.User
	mu.Lock()
	defer mu.Unlock()
	if fileData, err := os.ReadFile("db.json"); err == nil {
		err = json.Unmarshal(fileData, &users)
		if err != nil {
			return err
		}
	}
	users = append(users, user)
	file, err := os.Create("db.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(users)
	if err != nil {
		return err
	}
	return nil
}
