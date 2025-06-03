package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


var jwtSecret = []byte(os.Getenv("PASSWORD_HASH_KEY"))

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(userId string) (string, error){
  claims := &Claims{
	UserId: userId,
	RegisteredClaims: jwt.RegisteredClaims{
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	},
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  return token.SignedString(jwtSecret)
}

func VerifyJwtToken(tokenString string) (*Claims, error){
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error){
		return jwtSecret, nil
	})

	if err != nil{
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid{
		return claims, nil
	}

	return nil, errors.New("Invalid Token")
}
