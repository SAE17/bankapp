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
