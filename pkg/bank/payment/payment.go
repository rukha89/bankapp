package payment

import(
	"bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment{
	max := payments[0]
	for _, operation := range payments {
		if max.Amount < operation.Amount{
			max = operation
		}
	}
	return max

}