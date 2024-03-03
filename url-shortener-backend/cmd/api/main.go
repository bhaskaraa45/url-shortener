package main

import (
	"context"
	"fmt"
	"log"
	"url-shortener-backend/internal/auth"
	"url-shortener-backend/internal/server"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)


func main() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("credentials.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error initializing app: %v\n", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing Auth client: %v\n", err)
	}
	auth.FirebaseClient(client)
	log.Printf("Firebase Admin SDK initialized")

	server := server.NewServer()

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
