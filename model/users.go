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
type PinUser struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
	Pin    int `json:"pin"`
}
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results,omiempty"`
	Token   any    `json:"token,omiempty"`
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
func FindAllUser(search string) []Profile {
	conn, err := utils.DBConnect()
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	if err != nil {

	}
	rows, err := conn.Query(
		context.Background(), `SELECT * FROM users WHERE username ILIKE $1`, fmt.Sprintf("%%%s%%", search))
	users, err := pgx.CollectRows[Profile](rows, pgx.RowToStructByName)
	if err != nil {

	}
	return users
}
func CreatePin(pinUser PinUser) error {
	conn, err := utils.DBConnect()
	defer conn.Conn().Close(context.Background())
	if err != nil {
		fmt.Println("error ini")
		return err
	}
	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO pin (pin, id_user) VALUES ($1, $2)`, pinUser.Pin, pinUser.UserId,
	)
	fmt.Println("err")
	return err

}
