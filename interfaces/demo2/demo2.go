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

func CalculateMonthlyPayment(credits []CreditCalculator) float64 { //birden fazla kredii olabillir diye array olarak interfacei gönderdik
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}
