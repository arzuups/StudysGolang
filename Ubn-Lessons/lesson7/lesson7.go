// Golang Basic Lesson-7
package main

import "fmt"

func main() {

	//ELSE-İF == bool
	var isİtComing bool = true

	if isİtComing {
		fmt.Println("Yes,it is coming.")
	} else {
		fmt.Println("No, it is not coming.")
	}

	var myNumber int = 14
	youNumber := 20

	if myNumber <= youNumber {
		fmt.Println("myNumber is less than youNumber.")
	} else if myNumber == youNumber{
		fmt.Println("myNumber is equal than youNumber.")
	}else{
		fmt.Println("myNumber is greater than youNumber")
	}

	productPrice := 46
	myProduct := 50

	if productPrice >= myProduct {
		fmt.Println("Product price is greater than my product.")
	} else if productPrice == myProduct {
		fmt.Println("Product price is equal than my product.")
	} else {
		fmt.Println("Product price is less than my product.")

	}
}

