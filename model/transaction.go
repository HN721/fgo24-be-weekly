package model

import (
	"context"
	"weekly1/utils"

	"github.com/jackc/pgx/v5"
)

type Transactions struct {
	Amount int    `json:"amount"`
	Status string `json:"status"`
	Date   string `json:"date"`
	Method string `json:"method"`
	IdUser int    `json:"id_user"`
}

func CreateTransactions(id int, trans Transactions) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	trans.IdUser = id
	_, err = conn.Exec(context.Background(), `INSERT INTO transaction (amount,status,date_transaction,method,id_user)VALUES(
		$1, $2,$3,$4,$5
	)`, trans.Amount, trans.Status, trans.Date, trans.Method, trans.IdUser)
	defer conn.Conn().Close(context.Background())

	return err
}
func GetTransactionByUserId(userId int) ([]Transactions, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(context.Background(),
		`SELECT amount, status, date_transaction, method, id_user 
		FROM transaction 
		WHERE id_user = $1`, userId)
	if err != nil {
		return nil, err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	transactions, err := pgx.CollectRows[Transactions](rows, pgx.RowToStructByName)

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
