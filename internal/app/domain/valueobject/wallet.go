package valueobject

type Currency struct {
	NGN string
	GHS string
	USD string
}

type Wallet struct {
	ID       int
	Currency string
	Amount   uint
}
