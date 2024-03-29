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
	AddData(url string, shorturl string, clicked int, userId int) (bool, int)
	GetOGUrl(shorturl string) string
	GetAll(userId int) ([]model.DataModel, error)
	CreateUser(email string, name string) (bool, int)
	GetUser(userId int) (string, error)
	UserExists(email string) (bool, int)
	UrlAvaliable(shorturl string) bool
	GetData(id int) (model.DataModel, error)
	DeleteData(id int) bool
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

func (s *service) AddData(url string, shorturl string, clicked int, userId int) (bool, int) {
	que := "INSERT INTO urlshrink (url, shorturl, clicked, user_id) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := s.db.QueryRow(que, url, shorturl, clicked, userId).Scan(&id)
	if err != nil {
		log.Printf("Failed to add data, err: %v", err)
		return false, 0
	}
	return true,id
}

func (s *service) GetOGUrl(shorturl string) string {
	que := "SELECT * FROM urlshrink WHERE shorturl = $1"
	var data model.DataModel
	var id int
	err := s.db.QueryRow(que, shorturl).Scan(&id, &data.Url, &data.ShortUrl, &data.Clicked, &data.UserID)

	if err != nil {
		log.Printf("failed to fetch og url of %v", shorturl)
		log.Printf("error: %v", err)
		return "https://shrink.bhaskaraa45.me"
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

func (s *service) GetAll(userId int) ([]model.DataModel, error) {
	que := "SELECT * FROM urlshrink WHERE user_id = $1 ORDER BY id DESC"
	rows, err := s.db.Query(que, userId)

	if err != nil {
		log.Printf("Error while getting all data: %v", err)
		return nil, err
	}

	defer rows.Close()

	var alldata []model.DataModel
	for rows.Next() {
		var data model.DataModel
		err := rows.Scan(&data.Id, &data.Url, &data.ShortUrl, &data.Clicked, &data.UserID)
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

func (s *service) CreateUser(email string, name string) (bool, int) {
	que := "INSERT INTO users (email, name) VALUES ( $1, $2 ) RETURNING id"
	var id int
	err := s.db.QueryRow(que, email, name).Scan(&id)
	if err != nil {
		log.Printf("Failed to create user, err: %v", err)
		return false, 0
	}
	return true, id
}

func (s *service) GetUser(userId int) (string, error) {
	que := "SELECT * FROM users WHERE id = $1 "
	var email string
	var name string
	err := s.db.QueryRow(que, userId).Scan(&email, &name)
	if err != nil {
		log.Printf("Failed to get user, err: %v", err)
		return "", err
	}
	return email, nil
}

func (s *service) UserExists(email string) (bool, int) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	var exists bool

	err := s.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		log.Printf("error checking users email = %v, error: %v", email, err)
		return false, 0
	}

	if exists {
		var id int
		query = "SELECT id FROM users WHERE email = $1"
		err := s.db.QueryRow(query, email).Scan(&id)
		if err != nil {
			log.Printf("error checking users email = %v, error: %v", email, err)
			return false, 0
		}
		return true, id
	}

	return false, 0
}

func (s *service) UrlAvaliable(shorturl string) bool {
	query := "SELECT EXISTS(SELECT 1 FROM urlshrink WHERE shorturl = $1)"
	var exists bool

	err := s.db.QueryRow(query, shorturl).Scan(&exists)
	if err != nil {
		log.Printf("error checking url shorturl = %v, error: %v", shorturl, err)
		return true //if error then assume its not avlbl
	}
	return exists
}

func (s *service) GetData(id int) (model.DataModel, error) {
	query := "SELECT * FROM urlshrink WHERE id = $1 "
	var data model.DataModel
	err := s.db.QueryRow(query, id).Scan(&data.Id, &data.Url, &data.ShortUrl, &data.Clicked, &data.UserID)

	if err != nil {
		log.Printf("error fetching data of id = %v, error: %v", id, err)
		return data, err
	}

	return data, nil
}

func (s *service) DeleteData(id int) bool {
	query := "DELETE FROM urlshrink WHERE id = $1"
	_, err := s.db.Exec(query, id)
	if err != nil {
		log.Printf("error deleting data of id = %v, error: %v", id, err)
		return false
	}
	return true
}
