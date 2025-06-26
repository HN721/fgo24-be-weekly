package controller

import (
	"net/http"
	"os"
	"time"
	"weekly1/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateToken(username string, id int) (string, error) {
	godotenv.Load()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       username,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
func LoginController(ctx *gin.Context) {
	var input model.Profile

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Input tidak valid",
		})
		return
	}

	user, err := model.Login(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	token, err := CreateToken(user.Name, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Message: "Gagal membuat token",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Login berhasil",
		Results: user,
		Token:   token,
	})
}
func RegisterController(ctx *gin.Context) {
	var input model.Profile
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Input Tidak Valid",
		})
		return
	}
	err = model.Register(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Input Tidak Valid",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Sucessfully",
		Results: input,
	})
}
