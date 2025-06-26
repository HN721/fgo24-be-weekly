package model

import (
	"context"
	"weekly1/utils"
)

type Balance struct {
	IdUser int    `json:"idUser"`
	Saldo  int    `json:"saldo"`
	Date   string `json:"json"`
}

func TopUp(balanceUser Balance) error {
	conn, err := utils.DBConnect()

	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), `INSERT INTO pin (saldo,last_updated, id_user) VALUES ($1, $2,$3)`, balanceUser.Saldo, balanceUser.Date, balanceUser.IdUser)

	return err
}
