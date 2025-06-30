package model

import (
	"context"
	"fmt"
	"time"
	"weekly1/utils"
)

type UserBalanceResponse struct {
	Name  string `json:"name"`
	Saldo int    `json:"saldo"`
}

type Balances struct {
	IdUser int    `json:"idUser" form:"idUser" `
	Saldo  int    `json:"saldo" form:"saldo" `
	Date   string `json:"date" form:"date" `
}
type TopUps struct {
	Amount int `json:"amount" form:"amount" `
	UserID int
	Date   time.Time `json:"date" form:"date" `
}

func Balance(id int, balanceUser Balances) error {
	conn, err := utils.DBConnect()
	defer conn.Conn().Close(context.Background())

	if err != nil {
		return err
	}
	balanceUser.IdUser = id
	_, err = conn.Exec(context.Background(), `INSERT INTO balance (saldo,last_updated, id_user) VALUES ($1, $2,$3)`, balanceUser.Saldo, balanceUser.Date, balanceUser.IdUser)

	return err
}
func TopUp(id int, topup TopUps) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(context.Background())
	topup.UserID = id
	_, err = tx.Exec(context.Background(),
		`INSERT INTO topup (amount, date_transactions, id_user)
         VALUES ($1, $2, $3)`,
		topup.Amount, topup.Date, topup.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to insert into topup: %w", err)
	}
	_, err = tx.Exec(context.Background(),
		`INSERT INTO balance (saldo, last_updated, id_user)
         SELECT $1, CURRENT_TIMESTAMP, $2
         WHERE NOT EXISTS (
             SELECT 1 FROM balance WHERE id_user = $2
         )`,
		topup.Amount, topup.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to insert initial balance: %w", err)
	}
	_, err = tx.Exec(context.Background(),
		`UPDATE balance
         SET saldo = saldo + $1,
             last_updated = CURRENT_TIMESTAMP
         WHERE id_user = $2`,
		topup.Amount, topup.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to update balance: %w", err)
	}
	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
func GetUserBalance(userID int) (string, int, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return "", 0, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer conn.Conn().Close(context.Background())

	var name string
	var saldo int

	err = conn.QueryRow(context.Background(), `
		SELECT u.name, b.saldo
		FROM users u
		JOIN balance b ON u.id = b.id_user
		WHERE u.id = $1
	`, userID).Scan(&name, &saldo)

	if err != nil {
		return "", 0, fmt.Errorf("failed to get user balance: %w", err)
	}

	return name, saldo, nil
}
