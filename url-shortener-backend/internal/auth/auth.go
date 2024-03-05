package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"url-shortener-backend/internal"
	"url-shortener-backend/internal/database"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type TokenInfo struct {
	Email string `json:"email"`
}

var client *auth.Client

func FirebaseClient(c *auth.Client) {
	client = c
}

type Token struct {
	IdToken string `json:"idToken"`
}

func Verify(idToken string) (bool, string, int) {
	ctx := context.Background()
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("Error verifying ID token: %v\n", err)
		return false, "", 0
	}

	// Get user email from the verified token
	userEmail := token.Claims["email"].(string)
	fmt.Printf("User Email: %s\n", userEmail)

	isExists, id := database.New().UserExists(userEmail)

	if isExists {
		return true, userEmail, id
	}

	res, id_ := database.New().CreateUser(userEmail)

	return res, userEmail, id_

}

func HandleLogin(c *gin.Context) {
	var data Token
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		resp := internal.CustomResponse("ivalid json data!", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	idToken := data.IdToken

	if idToken == "" {
		c.JSON(http.StatusBadRequest, internal.CustomResponse("ID token not provided", http.StatusBadRequest))
		return
	}

	res, email, userId := Verify(data.IdToken)

	if !res {
		c.JSON(http.StatusUnauthorized, internal.CustomResponse("Failed to verify token", http.StatusUnauthorized))
		return
	}

	jwtToken := jwt.New(jwt.SigningMethodHS256)
	jwtToken.Claims = jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(3 * 24 * time.Hour).Unix(),
	}

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, internal.CustomResponse("Failed to create token", http.StatusInternalServerError))
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Domain:   "localhost",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().Add(3 * 24 * time.Hour),
		Value:    tokenString,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(c.Writer, &cookie)

	resp := make(map[string]string)
	resp["email"] = email
	c.JSON(http.StatusOK, resp)
}

func HandleLogout(c *gin.Context) {
	cookie := http.Cookie{
		Name:     "token",
		Domain:   "localhost",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		Value:    "",
		Expires:  time.Now().Add(0 * time.Second),
		SameSite: http.SameSiteNoneMode,
		MaxAge:   -1,
	}
	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusAccepted, internal.CustomResponse("successfully logged out!", http.StatusAccepted))
}
