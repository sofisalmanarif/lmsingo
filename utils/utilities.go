package utilities

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(id int) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(id),                      //issuer contains the ID of the user.
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //Adds time to the token i.e. 24 hours.
	})

	token, err = claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, err

}
