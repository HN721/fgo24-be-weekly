package model

import (
	"context"
	"weekly1/utils"
)

type Transactions struct {
	amount  int
	status  string
	date    string
	method  string
	id_user int
}

func CreateTransactions(trans Transactions) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	_, err = conn.Exec(context.Background(), `INSERT INTO transaction (amount,status,date_transaction,method,id_user)VALUES(
		$1, $2,$3,$4,$5
	)`, trans.amount, trans.status, trans.date, trans.method, trans.id_user)

	return err
}
