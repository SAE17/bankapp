package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	car := types.Card{Balance: 5_000_000_00, MinBalance: 4_000_000_00, Active: true}
	card.AddBonus(&car, 3, 30, 365)
	fmt.Println(car.Balance)
}
