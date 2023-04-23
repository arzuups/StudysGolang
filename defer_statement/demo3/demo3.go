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
}
