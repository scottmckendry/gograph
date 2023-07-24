package models

import (
	"encoding/json"
	"goGraph/auth"
	"io"
)

type User struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"mail"`
}

func GetUser() (User, error) {
	// API call to Graph
	response, err := auth.MakeRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)
	if err != nil {
		return User{}, err
	}
	defer response.Body.Close()

	// Unmarshal response body into User struct
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return User{}, err
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, err
	}

	return user, nil
}
