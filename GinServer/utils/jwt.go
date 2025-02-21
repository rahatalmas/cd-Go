package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var sECRET_KEY_JWT = []byte("secret_key_jwt")

//const pRIVATE_KEY_JWT = "private_key_jwt"

func AccessToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                         // Subject user identifier
		"iss": "gin-server",                     // Issuer
		"aud": "user",                           // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})
	tokenString, err := claims.SignedString(sECRET_KEY_JWT)
	if err != nil {
		return "", err
	}
	//fmt.Printf("Token claims added: %+v\n", claims)
	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return sECRET_KEY_JWT, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if username, ok := claims["sub"].(string); ok {
			return username, nil
		}
		return "", fmt.Errorf("invalid token claims")
	}

	return "", fmt.Errorf("invalid token")
}
