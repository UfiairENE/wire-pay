package repository

import (
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres"
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres/model"
	"github.com/UfiairENE/send-wire-pay/internal/app/domain/valueobject"
)

//create user
func Create(u valueobject.NewUser) interface{} {
	db := postgres.Connection()
	user := model.User{FirstName: u.FirstName, LastName: u.LastName, Email: u.Email}
	result := db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	wallet := CreateWalletOnsignUp(db, user.ID)

	resp := struct {
		User   model.User
		Wallet model.Wallet
	}{
		User:   user,
		Wallet: wallet,
	}
	return resp
}
