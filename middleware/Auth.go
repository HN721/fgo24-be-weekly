package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func createToken(username string) (string, error) {
	godotenv.Load()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
func AuthMiddleware() {

}
