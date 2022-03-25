package usecase

import (
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/repository"
	"github.com/UfiairENE/send-wire-pay/internal/app/domain/valueobject"
)

//create acount is the use case for creating an account
func Createaccount(User valueobject.NewUser) interface{} {
	return repository.Create(User)
}
