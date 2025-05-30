package models

import (
	"encoding/json"
	"errors"
	"os"
	"github.com/Desyncq/warden/services"
)

type Admin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (a *Admin) CheckPassword(password string) bool {
	hashingService := services.HashService{}
	isValid := hashingService.CheckPasswordHash(password, a.Password)
	return isValid
}

func LoadAdminData() (*Admin, error) {
	exists := fileExists("admin.json")
	if !exists {
		return nil, errors.New("admin file not created")
	}

	file, err := os.Open("admin.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var admin Admin
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&admin)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}