package card

import (
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

func Withdraw(card types.Card, amount int64) types.Card {
	if !card.Active {
		return card
	}

	if amount > card.Balance {
		return card
	}

	if amount < 0 {
		return card
	}

	if amount > 20_000_00 {
		return card
	}

	card.Balance -= amount
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	if amount > 50_000_00 {
		return
	}
	if !card.Active {
		return
	}
	if amount < 0 {
		return
	}
	card.Balance += int64(amount)

	return
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}
	if card.Balance <= 0 {
		return
	}

	bonus := ((card.MinBalance * int64(percent) / 100) * int64(daysInMonth)) / int64(daysInYear)

	if bonus > 5_000_00 {
		return
	}

	card.Balance += bonus

}

func Total(cards []types.Card) (sum types.Money) {

	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance < 0 {
			continue
		}
		sum += types.Money(card.Balance)
	}
	return
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymenysSource []types.PaymentSource

	for _, card := range cards {
		source := types.PaymentSource{
			Type: "card",
		}
		if !card.Active {
			continue
		}

		if card.Balance <= 0 {
			continue
		}
		source.Number = string(card.PAN)
		source.Balance = types.Money(card.Balance)
		paymenysSource = append(paymenysSource, source)
	}

	return paymenysSource

}
