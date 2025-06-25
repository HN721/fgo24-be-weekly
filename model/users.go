package model

import (
	"context"
	"fmt"
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
func Login(user Profile) (Profile, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return Profile{}, err
	}
	defer conn.Conn().Close(context.Background())

	var dbUser Profile

	err = conn.QueryRow(context.Background(),
		`SELECT id, name, email, images, password FROM users WHERE email = $1`,
		user.Email,
	).Scan(&dbUser.Id, &dbUser.Name, &dbUser.Email, &dbUser.Images, &dbUser.Password)

	if err != nil {
		return Profile{}, fmt.Errorf("email tidak ditemukan")
	}

	if user.Password != dbUser.Password {
		return Profile{}, fmt.Errorf("password salah")
	}

	return dbUser, nil
}
