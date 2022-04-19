package utils

import (
	"context"
	"fmt"
	"time"
	"todo/prisma"

	"todo/repository"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Secret string
	Client *prisma.PrismaClient
}

type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func (j *JWT) ParseToken(token string, ctx context.Context) (bool, error) {
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match

	tokenString, _ := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if claims, ok := tokenString.Claims.(*Claims); ok && tokenString.Valid {
		fmt.Printf("%v %v", claims.UserId, claims.StandardClaims.ExpiresAt)
		if claims.UserId == "" {
			return false, fmt.Errorf("User id is empty")
		} else {
			userRepository := repository.NewUserRepository(j.Client, ctx)
			if user, _ := userRepository.FindOne(claims.UserId); user != nil {
				return true, nil
			}
		}
	}

	// return user with the user id

	return false, nil
}

func (j *JWT) CreateToken(userID string) (string, error) {

	claims := &Claims{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)

	// create the encoded token string
	tokenString, _ := token.SignedString([]byte(j.Secret))

	return tokenString, nil
}
