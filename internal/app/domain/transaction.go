package domain

import "time"

//payment is the model of payments
type Transacton struct {

	WalletID      string    
	Currency      string    
	CardType      string    
	Description   string    
	Amount        int   
	Time          time.Time
	TransactionID string    
}

