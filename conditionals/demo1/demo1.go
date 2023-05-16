/*You need to run main.go to print to the terminal.*/
package conditionals
 
import (
	"fmt"
)

func Demo1() {
	var account float64 = 2000
	var requested float64 = 400
	
	//if keyword followed by a condition
        // if the condition is true, the if block is executed
        // if the condition is invalid, the else block is executed
        //The expressions we call conditions here can be any question pattern that we can answer yes or no.
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
