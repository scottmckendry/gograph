package main

import (
	"fmt"
	"goGraph/models"
)

func main() {
	currentUser, err := models.GetUser()
	if err != nil {
		fmt.Printf("Error getting user: %v", err)
	}

	fmt.Printf("User: %v\n", currentUser.DisplayName)
	fmt.Printf("Email: %v\n", currentUser.Email)
	fmt.Printf("Id: %v\n", currentUser.Id)

	emails, err := models.GetEmails()
	if err != nil {
		fmt.Printf("Error getting emails: %v", err)
	}

	for _, email := range emails {
		fmt.Printf("Subject: %v\n", email.Subject)
		fmt.Printf("Body: %v\n", email.BodyPreview)
		fmt.Printf("Sent: %v\n", email.Sent)
		fmt.Printf("Recieved: %v\n", email.Recieved)
		fmt.Println()
	}
}