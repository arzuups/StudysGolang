/*You need to run main.go to print to the terminal.*/
package defer_statement

import (
	"fmt"
)

type product struct {
    productName string
    unitPrice int
}

func (p product) Save() {
    fmt.Println("Product saved :" , p.productName)	
    defer Log()
}

func Log() {
    fmt.Println("Logged!")	
}

func Demo3() {
    p := product{productName: "Telephone" , unitPrice: 17000}
    p = product{productName: "Powerbank" , unitPrice: 350}
	
    fmt.Println("Operation successful!")
    fmt.Println(p.productName)
    defer p.Save()
}
