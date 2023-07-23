// Golang Basic Lesson-9
package main

import "fmt"

func main() {

	//ARRAYS
	//0 - 1 - 2 => Arrays starts zero.

	var cars [3]string = [3]string{"Qashqai", "Courier", "Bmw"}
	var models [3]string = [3]string{"Nıssan", "Ford", "Ford"}
	var fees [3]float32 = [3]float32{800.000, 560.000, 900.500}

	fmt.Println(cars, models, fees)
	fmt.Print("CAR FEATURES: ", cars, models, fees)
	fmt.Println(cars[2], models[2], fees[2])

	var authors [4]string = [4]string{"Yaşar KEMAL", "Halide Edib ADIVAR", "Oğuz ATAY", "Kemal TAHİR"}
	var authorBookNames [4]string = [4]string{"Sarı Sıcak", "Sinekli Bakkal", "Tutunamayanlar", "Devlet Ana"}
	var booksReleaseDates [4]int = [4]int{1952, 1935, 1972, 1967}

	fmt.Println(authors, authorBookNames, booksReleaseDates)
	fmt.Print("BOOK FEATURES: ")
	fmt.Println(authors[2], authorBookNames[2], booksReleaseDates[2])

}
