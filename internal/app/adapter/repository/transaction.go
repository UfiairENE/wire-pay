package repository

import (
	"time"

	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres"
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres/model"
	"github.com/google/uuid"
)

const (
	minimumDepositAmount    = 50 // least possible amount that can be deposited into an account
	minimumWithdrawalAmount = 10 // least possible amount that can be withdrawn from an account
)

type Repository interface {
	Add(model.Transaction) (model.Transaction, error)
	GetTransactions(userId uuid.UUID, from time.Time, limit int) (*[]model.Transaction, error)
}

func Add(tx model.Transaction) (model.Transaction, error) {
	db := postgres.Connection()
	result := db.Create(&tx)
	if result.Error != nil {
		panic(result.Error)
	}
	return tx, nil
}

func GetTransactions(UserId uuid.UUID, from time.Time, limit int) (*[]model.Transaction, error) {
	var transactions []model.Transaction
	db := postgres.Connection()
	result := db.Where(
		model.Transaction{},
	).Where(
		"timestamp <= ?", from,
	).Order("timestamp desc").Limit(limit).Find(&transactions)

	if result.Error != nil {
		panic(result.Error)
	}

	return &transactions, nil
}
