package model

import (
	"context"
	"weekly1/utils"

	"github.com/jackc/pgx/v5"
)

type Historymodel struct {
	IdHistory       int    `json:"idHistory" form:"idHistory"`
	IdTransaction   int    `json:"idTransaction" form:"idTransaction" binding:"required"`
	Amount          int    `json:"amount" form:"amount" binding:"required"`
	Status          string `json:"status" form:"status" binding:"required"`
	DateTransaction string `json:"dateTransaction" form:"dateTransaction" binding:"required"`
	Method          string `json:"method" form:"method" binding:"required"`
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
func GetAllHistory() ([]Historymodel, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	querys := `
		SELECT 
			h.id AS id_history,
			t.id AS id_transaction,
			t.amount,
			t.status,
			TO_CHAR(t.date_transaction, 'YYYY-MM-DD') AS date_transaction,
			t.method
		FROM history h
		JOIN transaction t ON h.id_transaction = t.id
	`

	rows, err := conn.Query(context.Background(), querys)
	if err != nil {
		return nil, err
	}

	histories, err := pgx.CollectRows[Historymodel](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	return histories, nil
}
