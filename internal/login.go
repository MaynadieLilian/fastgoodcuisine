package internal

import (
	"encoding/json"
	"fastgoodcuisine/internal/model"
	"golang.org/x/crypto/bcrypt"
	"os"
	"sync"
)

var muLogin sync.Mutex

func Login(email, password string) (*model.User, error) {
	muLogin.Lock()
	defer muLogin.Unlock()

	fileData, err := os.ReadFile("db.json")
	if err != nil {
		return nil, err
	}
	var users []model.User
	if err := json.Unmarshal(fileData, &users); err != nil {
		return nil, err
	}
	for _, u := range users {
		if u.Email == email {
			if err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(password)); err != nil {
				return nil, err
			}
			return &u, nil
		}
	}
	return nil, os.ErrNotExist
}
