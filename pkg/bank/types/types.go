package types

// Money represents the amount of money in (diram, cents, kopeyka end etc.) 
type Money int64

//Currency represents the code of currency
type Currency string

//codes of currencies
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN represents the card number
type PAN string

//Card rpresents information about card
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

//Payment represents information about payment
type Payment struct {
	ID int
	Amount Money
}

//PaymentSource represents source of payment
type PaymentSource struct {
	Type       string // 'card'
	Number     string // номер вида '5058 xxxx xxxx 8888'
	Balance    Money // баланс в дирамах
}