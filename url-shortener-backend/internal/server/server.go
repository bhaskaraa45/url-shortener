package server

import (
	"os"
	"strconv"

	"url-shortener-backend/internal/database"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   database.Service
	*gin.Engine
}

func NewServer() *Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	server := &Server{
		port:   port,
		db:     database.New(),
		Engine: gin.Default(),
	}
	server.RegisterRoutes()

	return server
}
