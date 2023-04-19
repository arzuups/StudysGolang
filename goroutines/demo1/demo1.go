package goroutines

import (
	"fmt"
	"time"
)

func EvenNumbers() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Even number : ", i)
		time.Sleep(1 * time.Second)
	}

}
