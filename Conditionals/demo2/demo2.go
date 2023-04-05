package conditionals

import (
	"fmt"
)

func Demo2() {
	var account float64 = 2000
	var requested float64 = 400

	if requested > account {
		fmt.Println("Not enough money in the account.")
	} else if requested == account {
		fmt.Println("Your money is being prepared...")
	} else {
		fmt.Println("Your money is being prepared...")
	}
}
