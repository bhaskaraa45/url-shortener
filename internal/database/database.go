package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"url-shortener-backend/internal/model"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string
	AddData(url string, shorturl string, clicked int) bool
	GetOGUrl(shorturl string) string
}

type service struct {
	db *sql.DB
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() Service {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) AddData(url string, shorturl string, clicked int) bool {
	que := "INSERT INTO urlshrink (url, shorturl, clicked) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(que, url, shorturl, clicked)
	if err != nil {
		log.Printf("Failed to add data, err: %v", err)
		return false
	}
	return true
}

func (s *service) GetOGUrl(shorturl string) string {
	que := "SELECT url FROM urlshrink WHERE shorturl = $1"
	var data model.DataModel
	rows, err := s.db.Query(que, shorturl)
	
	if err != nil {
		log.Printf("failed to fetch og url of %v", shorturl)
		log.Printf("error: %v", err)
		return "bhaskaraa45.me"
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&data.ShortUrl)
		if err != nil {
			log.Printf("failed to scan og url of %v", shorturl)
			log.Printf("error: %v", err)
			return "bhaskaraa45.me"
		}
	}
	
	if err := rows.Err(); err != nil {
        log.Printf("failed to iterate over rows for og url of %v", shorturl)
        log.Printf("error: %v", err)
        return "bhaskaraa45.me"
    }

	return data.ShortUrl
}
