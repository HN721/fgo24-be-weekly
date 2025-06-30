package model

import (
	"context"
	"fmt"
	"weekly1/utils"
)

type Profile struct {
	Id       int    `json:"id" form:"id" `
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Images   string `json:"images" form:"images" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
type PublicProfile struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Images string `json:"images"`
}
type ChangePasswordInput struct {
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required,min=6"`
}
type PinUser struct {
	Id     int `json:"id"`
	UserId int `json:"userId" form:"userId" binding:"required"`
	Pin    int `json:"pin" form:"pin" binding:"required,numeric"`
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
func FindAllUser(search string, limit, offset int) ([]Profile, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
		SELECT id, name, email, images 
		FROM users 
		WHERE name ILIKE '%' || $1 || '%' 
		ORDER BY id 
		LIMIT $2 OFFSET $3
	`

	rows, err := conn.Query(context.Background(), query, search, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []Profile
	for rows.Next() {
		var u Profile
		if err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Images); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
func FindOneUserById(id int) (Profile, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return Profile{}, err
	}
	defer conn.Conn().Close(context.Background())

	var user Profile
	query := `SELECT id, name, email, images FROM users WHERE id = $1`
	err = conn.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Name, &user.Email, &user.Images)
	if err != nil {
		return Profile{}, err
	}

	return user, nil
}

func FindOneUserByEmail(email string) (Profile, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return Profile{}, err
	}
	defer conn.Conn().Close(context.Background())

	var user Profile
	query := `SELECT id, name, email, images FROM users WHERE email = $1`
	err = conn.QueryRow(context.Background(), query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Images)
	if err != nil {
		return Profile{}, err
	}

	return user, nil
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
