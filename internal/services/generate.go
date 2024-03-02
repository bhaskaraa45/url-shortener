package services

import (
	"log"

	"github.com/teris-io/shortid"
)

func GenerateShortId() (string, error) {
	id, err := shortid.Generate()
	if err != nil {
		log.Printf("error while generating shortid: %v",err)
		return "", err
	}
	return id , nil
}
