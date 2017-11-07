package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/liteByte/mascotas-backend/config"
	"time"
)

type CustomClaims struct {
	Username string
	jwt.StandardClaims
}

type Token struct {
	Username string
}

func CreateToken(username string) string {
	claims := CustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "TODO change to application name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(config.GetConfig().JWT_SECRET))

	return tokenString
}

func GetTokenData(tokenString string) Token {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT_SECRET), nil
	})
	if err != nil {
		return Token{}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return Token{
			claims.Username,
		}
	} else {
		return Token{}
	}
}
