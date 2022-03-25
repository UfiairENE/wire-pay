package repository

import (
	"strings"

	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres"
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres/model"
	"github.com/UfiairENE/send-wire-pay/internal/app/domain/valueobject"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateWalletOnsignUp(db *gorm.DB, userID uint) model.Wallet {
	wallet := model.Wallet{User_id: int(userID)}
	result := db.Create(&wallet)
	if result.Error != nil {
		panic(result.Error)
	}
	return wallet
}

//create user
func CreateWallet(w valueobject.Wallet) model.Wallet {
	db := postgres.Connection()
	var wallet model.Wallet
	currencySymbol := strings.ToLower(w.Currency)
	switch currencySymbol {
	case "ngn":
		wallet = model.Wallet{NGNAmount: w.Amount, USDAmount: 0, GHSAmount: 0, User_id: w.ID}
	case "ghs":
		wallet = model.Wallet{NGNAmount: 0, USDAmount: 0, GHSAmount: w.Amount, User_id: w.ID}
	case "usd":
		wallet = model.Wallet{NGNAmount: 0, USDAmount: w.Amount, GHSAmount: 0, User_id: w.ID}
	default:
		panic("not implemented")
	}

	result := db.FirstOrCreate(&wallet)
	if result.Error != nil {
		panic(result.Error)
	}
	return wallet
}

// // // Get gets wallet balance
func GetBalance(w valueobject.Wallet) model.Wallet {
	db := postgres.Connection()
	var wallet model.Wallet
	result := db.First(&wallet, "User_id = ?", w.ID)
	if result.Error != nil {
		panic(result.Error)
	}
	return wallet

}

// UpdateBalance
func UpdateBalance(amount uint, userID uuid.UUID) (model.Wallet, error) {
	var wallet model.Wallet
	db := postgres.Connection()
	result := db.Model(model.Wallet{}).Where(model.Wallet{}).Updates(model.Wallet{}).Scan(&wallet)
	if result.Error != nil {
		panic(result.Error)
	}

	return wallet, nil
}
