package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

const (
	TxTypeDeposit    = "deposit"
	TxTypeWithdrawal = "withdrawal"
	TxTypeBalance    = "balance_enquiry"
)

//payment is the model of payments
type Transaction struct {
	gorm.Model
	WalletID      string
	Currency      string
	Description   string
	Amount        int
	Timestamp     time.Time
	TransactionID string
	UserID        uuid.UUID
}

type FundRequest struct {
	UserID   int
	Currency string
	Amount   uint
}

type FlutterRequestBody struct {
	TxRef          string         `json:"tx_ref"`
	Amount         string         `json:"amount"`
	Currency       string         `json:"currency"`
	RedirectUrl    string         `json:"redirect_url"`
	PaymentOptions string         `json:"payment_options"`
	Customer       Customer       `json:"customer"`
	Customizations Customizations `json:"customizations"`
}
type Customer struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Customizations struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Transactions struct {
	ID            uint   `json:"id"`
	TransactionID uint   `json:"transaction_id" gorm:"column:transaction_id; type:int"`
	Amount        int    `json:"amount" gorm:"column:amount; type:int"`
	Trxref        string `json:"trx_ref" gorm:"column:trx_ref; type:string"`
	UserID        int    `json:"user_id" gorm:"column:user_id; type:int"`
}

type WithdrawalRequest struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
