//Go Basic Lesson-5
package main

import "fmt"

func main() {

  //FUNCTIONS
  /*writeHello()//It will work when we write it this way.
  writeHello()
  writeHello()*/

  var number1 int = 43
  var number2 int = 32
  sumWrite(number1, number2)

  var number3 int = 54
  var number4 int = 94
  sumWrite(number3, number4)
  
}

/*func writeHello() {
  fmt.Println("Hello 1")//To get the output of this function we need to go to the main function and type the function name. Otherwise our function will not work
  fmt.Println("Hello 2")
  fmt.Println("Hello 3")
}*/

func sumWrite(number1 int , number2 int) {
  sum := number1 + number2
  fmt.Println(sum)
}
