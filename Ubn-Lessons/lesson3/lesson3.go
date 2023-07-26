// Golang Basic Lesson-3
package main

import "fmt"

func main() {
	//VARIABLES

	//string = "Hola!"
	//name := "Arzuups"
	var name string = "Arzuups"
	fmt.Println(name)

	//int = 1 2 3 4 5 6 7 8 9
	//age := 14
	var age int = 14
	fmt.Println(age)

	//float32 - float64
	//lengthTable := 103.54
	var lengthTable float32 = 101.4
	fmt.Println(lengthTable)

	//bool => true - false
	//aktifMi := true
	var aktifMi bool = false
	fmt.Println(aktifMi)

	//LİTTLE EXAMPLE//
	var myName string = "Arzu" + " " + "Demir" //If the "" sign is placed and a space is inserted between it, a space is inserted between the two words.
	var myAge int = 14
	var myHeight float32 = 1.70
	var isItTrue bool = true

	fmt.Println(myName)
	fmt.Println(myAge)
	fmt.Println(myHeight)
	fmt.Println(isItTrue)

}
