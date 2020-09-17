package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleTotal() {
	cards := []types.Card{
	    {
		    Balance: 10_000_00,
		    Active: false,
		},
		{
		    Balance: 10_000_00,
		    Active: true,
		},
		{
		    Balance: 10_000_00,
		    Active: true,
	    },
	}
	
	fmt.Println(Total(cards))
	// Output: 2000000
}
