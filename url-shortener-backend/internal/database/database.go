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
	GetAll() ([]model.DataModel, error)
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
	que := "SELECT * FROM urlshrink WHERE shorturl = $1"
	var data model.DataModel
	var id int
	err := s.db.QueryRow(que, shorturl).Scan(&id, &data.Url, &data.ShortUrl, &data.Clicked)

	if err != nil {
		log.Printf("failed to fetch og url of %v", shorturl)
		log.Printf("error: %v", err)
		return "bhaskaraa45.me"
	}

	s.UpdateClick(data.Clicked+1, shorturl)

	return data.Url
}

func (s *service) UpdateClick(click int, shorturl string) bool {
	que := "UPDATE urlshrink SET clicked = $1 WHERE shorturl = $2"
	_, err := s.db.Exec(que, click, shorturl)

	if err != nil {
		log.Printf("Failed to upadte click for shorturl: %v", shorturl)
		log.Printf("Error: %v", err)
		return false
	}
	return true
}

func (s *service) GetAll() ([]model.DataModel, error) {
	que := "SELECT * FROM urlshrink ORDER BY id ASC"
	rows, err := s.db.Query(que)

	if err != nil {
		log.Printf("Error while getting all data: %v", err)
		return nil, err
	}

	defer rows.Close()

	var alldata []model.DataModel
	for rows.Next() {
		var data model.DataModel
		var id int
		err := rows.Scan(&id, &data.Url, &data.ShortUrl, &data.Clicked)
		if err != nil {
			log.Printf("error scanning data row: %v", err)
			return nil, err
		}
		alldata = append(alldata, data)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error iterating over data rows: %v", err)
		return nil, err
	}

	return alldata, nil
}
