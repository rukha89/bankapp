package card

import (
	"bank/pkg/bank/types"
)

const bonuslimit = 5_000_00

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, operation := range cards {
		if operation.Balance > 0 {
			if operation.Active {
				sum += operation.Balance
			}
		}
	}
	return sum
}
