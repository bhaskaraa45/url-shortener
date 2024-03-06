package guard

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(cookieToken string) (bool, int) {
	result, err := jwt.Parse(cookieToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !result.Valid {
		return false, 0
	}

	mp := result.Claims.(jwt.MapClaims)
	subject, ok := mp["sub"]

	if !ok {
		return false, 0
	}

	exp, ok := mp["exp"].(float64)
	if !ok {
		return false, 0
	}

	expirationTime := time.Unix(int64(exp), 0)
	sub := fmt.Sprintf("%v", subject)
	sub_int, err := strconv.Atoi(sub)

	if err != nil {
		return false, 0
	}

	return !expirationTime.Before(time.Now()), sub_int

}
