package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payment := []types.Payment{
		{
			ID:     1,
			Amount: 100,
		},
		{
			ID:     2,
			Amount: 1000,
		},
		{
			ID:     3,
			Amount: 10000,
		},
	}
	max := Max(payment)
	fmt.Println(max.ID)
	//Output: 3
}
