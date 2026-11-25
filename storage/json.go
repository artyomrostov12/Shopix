package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"shopix/model"
)

var users []model.Registration

func SavedUsers(users []model.Registration) {
	file, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("users.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись прошла успешно")
}

func LoadUsers() ([]model.Registration, error) {

	data, err := os.ReadFile("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	fmt.Println("Данные успешно загружены")

	return users, nil
}
