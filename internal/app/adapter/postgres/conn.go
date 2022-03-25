package postgres

import (
	"fmt"
	"log"

	"github.com/UfiairENE/send-wire-pay/internal/app/adapter/postgres/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var postgresdb *gorm.DB

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Connection gets connection of postgresql database
func Connection() (db *gorm.DB) {
	return postgresdb
}

func StartDB() (db *gorm.DB) {
	// dsn := viper.GetString("DB_SOURCE")
	dsn := "postgres://jvtkaahtroqcbg:476b802ed25ea70a1be17da685811594d6aac0fb986b637e3c760c32c0a899fa@ec2-50-19-32-96.compute-1.amazonaws.com:5432/d32s70iousao4l"
	fmt.Println("Check string", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(model.User{}, model.Wallet{}, model.Transactions{})
	postgresdb = db
	return db
}
