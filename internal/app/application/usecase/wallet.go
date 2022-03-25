package usecase

import (
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres/model"
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/repository"
	"github.com/UfiairENE/send-wire-pay/internal/app/domain/valueobject"
)

//create acount is the use case for creating an wallet
func Createwallet(User valueobject.Wallet) model.Wallet {
	return repository.CreateWallet(User)
}

func GetBalance(Wallet valueobject.Wallet) model.Wallet {
	return repository.GetBalance(Wallet)
}
