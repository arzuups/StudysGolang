/*You need to run main.go to print to the terminal.*/
package functions

func TotalVariadic(numbers ...int) int {
    sum := 0
    for i := 0; i < len(numbers); i++ {
        sum = sum + numbers[i]
    }
  return sum
}
