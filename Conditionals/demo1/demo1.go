package conditionals

import (
	"fmt"
)

func Demo1() {
	var account float64 = 2000
	var requested float64 = 400

	if requested > account {
		fmt.Println("Not enough money in the account.")
	}

	if requested <= account {
		fmt.Println("Your money is being prepared...")
		account = account - requested
	}

	fmt.Println("Finish! Money in the account : " + fmt.Sprintf("%v", account))
	//fmt.Printf("Finish! Money in the account : %v", account) --> fmt.Printf is written as
}
