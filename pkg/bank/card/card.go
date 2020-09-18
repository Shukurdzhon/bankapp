package card

import (
	"bank/pkg/bank/types"
)

//IssueCard generates a new card
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:         1000,
		PAN:        "5058 xxxx xxxx 0001",
		Balance:    0,
		Currency:   currency,
		Active:     true,
		Color:      color,
		Name:       name,
		MinBalance: 0,
	}
	return card
}

const withdrawLimit = 20_000_00

// Withdraw allows to take out money from card
func Withdraw(card *types.Card, amount types.Money) {
	if amount < 0 {
		return
	}

	if amount > withdrawLimit {
		return
	}

	if !(*card).Active {
		return
	}

	if amount > (*card).Balance {
		return
	}

	(*card).Balance -= amount
}

const depositLimit = 50_000_00

// Deposit allows to add money to card
func Deposit(card *types.Card, amount types.Money) {
	if !(*card).Active {
		return
	}

	if amount > depositLimit {
		return
	}

	if amount < 0 {
		return
	}

	(*card).Balance += amount
}

// AddBonus adds cashback in a period
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !(*card).Active {
		return
	}

	if (*card).Balance < 0 {
		return
	}

	bonus := int((*card).MinBalance) * percent / 100 * daysInMonth / daysInYear

	if bonus > 5_000_00 {
		return
	}

	(*card).Balance += types.Money(bonus)
}

func sum(operations []int64) int64 {
	sum := int64(0)
	for _, operation := range operations {
		sum += operation
	}
	return sum
}


//Total sums all amount of money on cards
//Negative Balance and Inactive cards are IGNORED
func Total(cards []types.Card) types.Money {
	var operations []int64
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
		operations = append(operations, int64(card.Balance))
	}

	sum := sum(operations)
	return types.Money(sum)
}

//PaymentSources gives slice of valid card 
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var result []types.PaymentSource
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}

        var validCard = types.PaymentSource {
			Type: "card",
			Number: string(card.PAN),
			Balance: card.Balance,
		}

		result = append(result, validCard)
	}
	return result
}
