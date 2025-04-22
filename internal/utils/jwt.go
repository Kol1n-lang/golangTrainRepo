package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	"train-http/config"
)

type Claims struct {
	UserEmail string `json:"userEmail"`
	jwt.StandardClaims
}

func GenerateToken(userEmail string) (string, error) {
	cfg, _ := config.LoadConfig("config.toml")
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &Claims{
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}
func DecodeJWT(tokenString string, secret string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		log.Println(time.Now().Unix())
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func ValidateJWT(tokenString string) bool {
	cfg, _ := config.LoadConfig("config.toml")
	secret := cfg.JWTSecret()
	_, err := DecodeJWT(tokenString, secret)
	return err == nil
}
