package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener-backend/internal"
	"url-shortener-backend/internal/auth"
	"url-shortener-backend/internal/guard"
	"url-shortener-backend/internal/model"
	"url-shortener-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type Token struct {
	IdToken string `json:"idToken"`
}

var userId = 1 //TODO: change this

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		if c.FullPath() == "/verify" {
			c.Next()
			return
		}

		cookie, err := c.Request.Cookie("token")
		if err != nil || cookie == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token cookie not found"})
			c.Abort()
			return
		}
		cookieToken := "aa45"
		if cookie != nil {
			cookieToken = cookie.Value
		}

		res := guard.VerifyToken(cookieToken)
		if !res {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	})

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)
	r.POST("/add", s.handlePostData)
	r.GET("/:shorturl", s.handleShortUrlClick)
	r.GET("/getAll", s.handleGetAll)
	r.POST("/verify", auth.HandleLogin)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"
	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) handlePostData(c *gin.Context) {

	var data model.DataModel
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	fmt.Println(data)

	if err != nil {
		resp := internal.CustomResponse("ivalid json data!", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	shortUrl, err := services.GenerateShortId()
	if err != nil {
		resp := internal.CustomResponse(err.Error(), http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	res := s.db.AddData(data.Url, shortUrl, 0, userId)

	if res {
		// resp := internal.CustomResponse("data successfully added!", http.StatusOK)
		resp := make(map[string]string)
		resp["shorturl"] = shortUrl
		resp["url"] = data.Url
		resp["clicked"] = "0"
		c.JSON(http.StatusOK, resp)
	} else {
		resp := internal.CustomResponse("failed to add data!", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, resp)
	}
}

func (s *Server) handleShortUrlClick(c *gin.Context) {
	shorturl := c.Param("shorturl")

	og := s.db.GetOGUrl(shorturl)
	resp := make(map[string]string)
	resp["url"] = og
	resp["shorturl"] = shorturl
	// c.JSON(http.StatusOK, resp)
	c.Redirect(http.StatusSeeOther, og)
}

func (s *Server) handleGetAll(c *gin.Context) {
	res, err := s.db.GetAll(userId)
	if err != nil {
		resp := internal.CustomResponse("failed to get data!", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, resp)
	}
	c.JSON(http.StatusOK, res)
}

// func (s *Server) handleVerify(c *gin.Context) {
// 	var data Token
// 	err := json.NewDecoder(c.Request.Body).Decode(&data)
// 	if err != nil {
// 		resp := internal.CustomResponse("ivalid json data!", http.StatusBadRequest)
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}

// 	res, email := auth.Verify(data.IdToken)

// 	if !res {
// 		c.JSON(http.StatusUnauthorized, internal.CustomResponse("Failed to verify token", http.StatusUnauthorized))
// 		return
// 	}
// 	resp := make(map[string]string)
// 	resp["email"] = email
// 	c.JSON(http.StatusOK, resp)
// }
