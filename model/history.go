package model

import (
	"context"
	"weekly1/utils"
)

type Historymodel struct {
	IdHistory     int
	IdTransaction int `json:"idTransaction"`
}

func CreateHistory(history Historymodel) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), `INSERT INTO history (id_transaction)VALUES(
		$1
	)`, history.IdTransaction)
	defer conn.Conn().Close(context.Background())

	return err
}
