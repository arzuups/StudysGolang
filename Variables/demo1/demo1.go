/*You need to run main.go to print to the terminal.*/
package variables

import (
	"fmt"
)

func Demo1() {

	//The "string" function is used for textual data.
	var text string = "Hello World..."
	    fmt.Print(text)
	    fmt.Print(text)
	    fmt.Print(text)
	    fmt.Println(text)

	//The "int" data type is used for natural numbers.
	//-5,-3,10,0,23456 etc.
        //var variablename variabletype = value -> 2nd variable definition method
	var age int = 20
	    fmt.Println(age)

	//The data types "float32" or "float64" are used for decimal numbers.
	var number float32 = 25.3
	    fmt.Println(number)

	//Represents true or false. If you do not specify any value, the default is zero.
	var stuatio bool

	var text1 string = "Olivia"
	var text2 string = "Michale"

	stuatio = text1 == text2

	    fmt.Println(stuatio)
	
	kdv3 =: 39.4
	     fmt.Println(kdv3)
	         fmt.Printf("data type : %T" , kdv3)

}
