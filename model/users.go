package model

import (
	"context"
	"fmt"
	"weekly1/utils"

	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Images   string `json:"images"`
	Password string `json:"password"`
}
type PublicProfile struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Images string `json:"images"`
}
type ChangePasswordInput struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
type PinUser struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
	Pin    int `json:"pin"`
}
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results,omitempty"`
	Token   string `json:"token,omitempty"`
}

func Register(user Profile) error {
	conn, err := utils.DBConnect()
	defer conn.Conn().Close(context.Background())
	if err != nil {
		return err
	}
	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO users (name, email, images, password) VALUES ($1, $2, $3, $4)`,
		user.Name, user.Email, user.Images, user.Password,
	)
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
func FindAllUser(search string) ([]Profile, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	rows, err := conn.Query(
		context.Background(),
		`SELECT id, name, email, images,password FROM users WHERE name ILIKE $1`,
		fmt.Sprintf("%%%s%%", search),
	)
	if err != nil {
		fmt.Println("QUERY ERROR:", err)
		return nil, err
	}

	users, err := pgx.CollectRows[Profile](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreatePin(id int, pinUser PinUser) error {
	conn, err := utils.DBConnect()
	defer conn.Conn().Close(context.Background())
	if err != nil {
		fmt.Println("error ini")
		return err
	}
	pinUser.UserId = id
	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO pin (pin, id_user) VALUES ($1, $2)`, pinUser.Pin, pinUser.UserId,
	)
	fmt.Println("err")
	return err

}
func UpdateProfile(userID int, input PublicProfile) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	_, err = conn.Exec(context.Background(), `
		UPDATE users 
		SET name = $1, email = $2, images = $3
		WHERE id = $4
	`, input.Name, input.Email, input.Images, userID)

	return err
}
func ChangePassword(userID int, oldPass, newPass string) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	var currentPassword string
	err = conn.QueryRow(context.Background(),
		`SELECT password FROM users WHERE id = $1`, userID).Scan(&currentPassword)
	if err != nil {
		return fmt.Errorf("user not found or error reading password")
	}

	if currentPassword != oldPass {
		return fmt.Errorf("old password is incorrect")
	}

	_, err = conn.Exec(context.Background(),
		`UPDATE users SET password = $1 WHERE id = $2`, newPass, userID)
	if err != nil {
		return fmt.Errorf("failed to update password")
	}

	return nil
}
