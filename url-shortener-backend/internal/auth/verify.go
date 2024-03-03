package auth

import (
	"context"
	"fmt"
	"log"
	"url-shortener-backend/internal/database"

	"firebase.google.com/go/auth"
)

type TokenInfo struct {
	Email string `json:"email"`
}

var client *auth.Client

func FirebaseClient(c *auth.Client) {
	client = c
}

func Verify(idToken string) (bool, string) {
	ctx := context.Background()
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("Error verifying ID token: %v\n", err)
		return false, ""
	}

	// Get user email from the verified token
	userEmail := token.Claims["email"].(string)
	fmt.Printf("User Email: %s\n", userEmail)

	res := database.New().CreateUser(userEmail)

	return res , userEmail

}
