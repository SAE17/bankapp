package payment

import "bank/pkg/bank/types"

func Max(payments []types.Payment) types.Payment {
	max := payments[0]
	for _, payment := range payments {
		if payment.Amount > max.Amount {
			max = payment
		}
	}
	return max
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
