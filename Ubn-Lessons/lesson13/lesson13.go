// Golang Basic Lesson-13
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomNumbers := rand.Intn(100) //It takes a random number between 0 and 100 but not 100.
	fmt.Println(randomNumbers)

	random1 := rand.Intn(50)
	fmt.Println(random1)

	time.Sleep(time.Second * 5) //Holds for the entered number and prints the next value.
	random2 := rand.Intn(20)
	fmt.Println(random2)

	fmt.Println("Please enter a name:")
	name := ""
	fmt.Scanf("%s", name) //Used to retrieve data from the user.
	fmt.Println(name)

	/*NOTES*/
	// %s == Used for String values.
	// %d == Used for İnt values.
	// %f == Used for float32(etc) values.

}
