package model

import (
	"context"
	"weekly1/utils"
)

type Profile struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Images   string `json:"images"`
	Password string `json:"password"`
}

func Register(user Profile) error {
	conn, err := utils.DBConnect()
	defer conn.Conn().Close(context.Background())
	if err != nil {

	}
	_, err = conn.Exec(context.Background(), `INSERT INTO users (name,email,images,pasword)`, user.Name, user.Email, user.Images, user.Password)
	return err
}
