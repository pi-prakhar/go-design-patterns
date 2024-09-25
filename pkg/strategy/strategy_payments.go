package strategy

import "fmt"

// Strategy interface for different payment methods
type PaymentMethod interface {
	Pay(amount float64)
}

// Concrete strategy for Credit Card
type CreditCard struct{}

func (cc CreditCard) Pay(amount float64) {
	fmt.Println("Paying %.2f using Credit Card.\n", amount)
}

// Concrete strategy for PayPal
type PayPal struct{}

func (pp PayPal) Pay(amount float64) {
	fmt.Printf("Paying %.2f using PayPal.\n", amount)
}

// Concrete strategy for Cryptocurrency
type Crypto struct{}

func (c Crypto) Pay(amount float64) {
	fmt.Printf("Paying %.2f using Cryptocurrency.\n", amount)
}

// Context that handles the payment
type PaymentContext struct {
	paymentMethod PaymentMethod
}

func (pc *PaymentContext) SetPaymentMethod(pm PaymentMethod) {
	pc.paymentMethod = pm
}

func (pc *PaymentContext) Pay(amount float64) {
	if pc.paymentMethod == nil {
		fmt.Println("No payment method selected")
		return
	}
	pc.paymentMethod.Pay(amount)
}
