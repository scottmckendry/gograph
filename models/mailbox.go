package models

import (
	"encoding/json"
	"goGraph/auth"
	"io"
)

type graphResponse struct {
	Value []Email `json:"value"`
}

type Email struct {
	Id          string `json:"id"`
	Subject     string `json:"subject"`
	BodyPreview string `json:"bodyPreview"`
	Sent        string `json:"sentDateTime"`
	Recieved    string `json:"receivedDateTime"`
}

func GetEmails() ([]Email, error) {
	// API call to Graph
	response, err := auth.MakeRequest("GET", "https://graph.microsoft.com/v1.0/me/messages", nil)
	if err != nil {
		return []Email{}, err
	}
	defer response.Body.Close()

	// Unmarshal response body into User struct
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []Email{}, err
	}

	var emails graphResponse
	if err := json.Unmarshal(body, &emails); err != nil {
		return []Email{}, err
	}

	return emails.Value, nil
}
