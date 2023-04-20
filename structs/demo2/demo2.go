/*You need to run main.go to print to the terminal.*/
package structs

import (
	"fmt"
)

type customer struct {
	firstName string
	lastName string
	age int
}

func (a customer) save() {
	fmt.Println("Added : ", a.firstName)
}

func (a customer) update() {
	fmt.Println("Updated  : ", a.firstName)
}

func Demo2() {
	a := customer{firstname: "Joseph" , lastname: "Smith" , age: 30}
	a.save()
	a.update
}
