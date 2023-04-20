package channels

import (
	"fmt"
	"time"
)

func EvenNumbers(evenNumberCn chan int) {
	sum := 0
	for i := 0; i <= 10; i += 2 {
		sum = sum + i
		fmt.Println("Even number addition works!")
		time.Sleep(1 * time.Second)
	}
	evenNumberCn <- sum
}

func OddNumbers(oddNumberCn chan int) {
	sum := 0
	for i := 1; i <= 10; i += 2 {
		sum = sum + i
		fmt.Println("Single number addition is working!")
		time.Sleep(1 * time.Second)
	}
	oddNumberCn <- sum
}
