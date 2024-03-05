package guard

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// func Guard(cookie) {

// }

func VerifyToken(cookieToken string) bool {
	result, err := jwt.Parse(cookieToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !result.Valid {
		return false
	}

	return true
}
