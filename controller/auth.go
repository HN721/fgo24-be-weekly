package controller

import (
	"net/http"
	"os"
	"strconv"
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
			"id":       id,
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
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Message: "Cant Register data",
		})
		return
	}
	ctx.JSON(http.StatusCreated, model.Response{
		Success: true,
		Message: "Sucessfully ",
		Results: input,
	})
}
func PinController(ctx *gin.Context) {
	var pin model.PinUser

	err := ctx.ShouldBindJSON(&pin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid Cannot Bind Js",
		})
		return
	}
	userId, _ := ctx.Get("userID")
	err = model.CreatePin(userId.(int), pin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid Pin",
		})
		return
	}
	ctx.JSON(http.StatusCreated, model.Response{
		Success: true,
		Message: "Sucess create Pin",
	})

}

func GetUser(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "5")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	users, err := model.FindAllUser(search, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Message: "Failed to search users: " + err.Error(),
		})
		return
	}

	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, model.Response{
			Success: false,
			Message: "No user found",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Successfully found users",
		Results: users,
	})
}
func GetUserById(ctx *gin.Context) {
	userIdValue, _ := ctx.Get("userID")

	userIdStr, ok := userIdValue.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid userID format",
		})
		return
	}

	user, err := model.FindOneUserById(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "User found",
		Results: user,
	})
}

// GET /users/email?email=nanda@gmail.com
func GetUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Email is required",
		})
		return
	}

	user, err := model.FindOneUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "User found",
		Results: user,
	})
}

func UpdateProfileController(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Success: false,
			Message: "Unauthorized: user ID not found",
		})
		return
	}
	userID := userIDVal.(int)

	var input model.PublicProfile
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	err := model.UpdateProfile(userID, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Message: "Failed to update profile: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Profile updated successfully",
		Results: input,
	})
}
func ChangePasswordController(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}
	userID := userIDVal.(int)

	var input model.ChangePasswordInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	err := model.ChangePassword(userID, input.OldPassword, input.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success: true,
		Message: "Password changed successfully",
	})
}
