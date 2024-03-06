package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

var userId = 0

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

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token cookie not found"})
			c.Abort()
		return

		//no guard waala routes

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

		res, sub := guard.VerifyToken(cookieToken)
		if !res {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		userId = sub

		c.Next()
	})

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	r.POST("/add", s.handlePostData)
	r.GET("/getAll", s.handleGetAll)
	r.POST("/verify", auth.HandleLogin)
	r.GET("/logout", auth.HandleLogout)
	r.POST("/check", s.handleCheckAvailability)
	r.DELETE("/delete/:id", s.handleDelete)
	r.GET("/:shorturl", s.handleShortUrlClick)

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

	shortUrl := data.ShortUrl

	if len(shortUrl) == 0 {
		temp, err := services.GenerateShortId()
		shortUrl = temp
		if err != nil {
			resp := internal.CustomResponse(err.Error(), http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, resp)
			return
		}
	}

	res, id := s.db.AddData(data.Url, shortUrl, 0, userId)

	if res {
		var resp model.DataModel
		resp.Id = id
		resp.Clicked = 0
		resp.Url = data.Url
		resp.ShortUrl = shortUrl
		resp.UserID = userId
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
	c.Redirect(http.StatusMovedPermanently, og)
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

func (s *Server) handleCheckAvailability(c *gin.Context) {
	var requestData model.DataModel
	err := json.NewDecoder(c.Request.Body).Decode(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}

	shorturl := requestData.ShortUrl

	fmt.Println(shorturl)

	if len(shorturl) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No data found"})
		return
	}

	exists := s.db.UrlAvaliable(shorturl)

	if exists {
		c.JSON(http.StatusConflict, internal.CustomResponse("already in use", http.StatusConflict))
		return
	}

	c.JSON(http.StatusOK, internal.CustomResponse("it's avaliable", http.StatusOK))
}

func (s *Server) handleDelete(c *gin.Context) {
	// var requestData model.DataModel
	// err := json.NewDecoder(c.Request.Body).Decode(&requestData)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
	// 	return
	// }
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if id < 1 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	data, err := s.db.GetData(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if data.UserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}
	res := s.db.DeleteData(id)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal sever error"})
		return
	}
	c.JSON(http.StatusOK, internal.CustomResponse("Successfully deleted", http.StatusOK))
}
