package model

type Amount interface {
	TableName() string
}

// wallet is model of wallets
type Wallet struct {
	ID        uint
	USDAmount uint `gorm:"type:int; not null"`
	NGNAmount uint `gorm:"type:int; not null"`
	GHSAmount uint `gorm:"type:int; not null"`
	User_id   int  `gorm:"column:user_id; type:int"`
}

//TableName gets table name of wallet
func (Wallet) TableName() string {
	return "amount"
}
