package card

import (
	"bank/pkg/bank/types"

	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 20_000_00}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}
func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Active: false}, 30_000_00)
	fmt.Println(result.Active)
	// Output: false
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 40_000_00, Active: true}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 4000000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_00)
	fmt.Println(card.Balance)
	// Output:
	// 2000000
}
func ExampleDeposit_limit() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 50_000_01)
	fmt.Println(card.Balance)
	// Output:
	// 2000000
}
func ExampleDeposit_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 20_000_00)
	fmt.Println(card.Balance)
	// Output:
	// 4000000
}
func ExampleDeposit_negative() {
	card := types.Card{Balance: -20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output:
	// -1000000
}

func ExampleAddBonus() {
	card := types.Card{Balance: 20_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output:
	// 2002465
}
func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 10_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output:
	// 1000000
}
func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -10_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output:
	// -1000000
}
func ExampleAddBonus_limitAbove() {
	card := types.Card{Balance: 5_000_000_00, MinBalance: 4_000_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output:
	// 500000000
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  false,
		},
	}
	fmt.Println(Total(cards))
	// Output: 2000000
}


