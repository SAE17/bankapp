package types

// Money представляет собой денежную сумму в минимальных единицах
type Money int64

// Currency представляет код валюты
type Currency string

// Код валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN представляет номер карты
type PAN string

// Card представляет информацию о платежной карте

type Card struct {
	ID         int
	PAN        PAN
	Balance    int64
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance int64
}

type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
