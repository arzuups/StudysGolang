package loops

import "fmt"

func Workshop2() {
	number := 0
	fmt.Println("Please enter a number:")
	fmt.Scanln(&number)

	prime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			prime = false
		}
	}

	if prime == true {
		fmt.Println("This is a prime number.")
	} else {
		fmt.Println("This is not a prime number!")
	}
}
