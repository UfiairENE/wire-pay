package main

import (
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter"
	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres"
)

func main() {
	postgres.StartDB()
	r := adapter.Router()
	r.Run("")
}
