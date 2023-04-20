/*You need to run main.go to print to the terminal.*/
package interfaces

import (
	"fmt"
)

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{creditPaymentTotal: 100000, rate: 10, address: "Istanbul"}
	credit2 := Mortgage{creditPaymentTotal: 50000, rate: 16, address: "Izmır"}
	credit3 := Car{creditPaymentTotal: 80000, rate: 8, carInfo: "Konya"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Total payment:", total)
}
