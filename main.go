package main

import (
	"encoding/json"
	"fmt"
	"goGraph/auth"
	"io"
	"net/http"
)

type User struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"mail"`
}

func main() {

	token, err := auth.GetToken()
	if err != nil {
		fmt.Printf("Error getting token: %v\n", err)
		return
	}

	// API call to Graph
	req, _ := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making API call: %v\n", err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("Error parsing response body:", err)
		return
	}

	fmt.Println(response.Status)
	fmt.Println(user.DisplayName)
}
