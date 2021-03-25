package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 10_000_00,
		},
		{
			ID:     2,
			Amount: 10_000_00,
		},
		{
			ID:     3,
			Amount: 5_000_00,
		},
	}
	res := Max(payments)
	fmt.Println(res.ID)
	// Output: 1
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			ID:      1000,
			PAN:     "5058 xxxx xxxx 0001",
			Balance: 200_00,
			Active:  false,
		},
		{
			ID:      1001,
			PAN:     "5058 xxxx xxxx 0002",
			Balance: 10_00,
			Active:  true,
		},
		{
			ID:      1002,
			PAN:     "5058 xxxx xxxx 0003",
			Balance: 1000_00,
			Active:  false,
		},
		{
			ID:      1003,
			PAN:     "5058 xxxx xxxx 0004",
			Balance: 0,
			Active:  true,
		},
	}
	sources := PaymentSources(cards)
	for _, source := range sources {
		fmt.Println(source.Number)
	}
	// Output:
	// 5058 xxxx xxxx 0002
}
