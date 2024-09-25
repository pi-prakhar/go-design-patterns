package main

import (
	"fmt"
	"go-design-patterns/pkg/strategy"
)

func main() {
	context := strategy.Context{}

	context.SetStrategy(strategy.ConcreteStrategyA{})
	context.ExecuteStrategy() // Uses strategy A

	context.SetStrategy(strategy.ConcreteStrategyB{})
	context.ExecuteStrategy() // Switches to strategy B

	amount := 100.0
	// Some logic to choose payment method based on various conditions
	paymentType := "PayPal" // Assume user chose PayPal

	// WITHOUT THE CONTEXT
	// Decision logic in client code
	switch paymentType {
	case "CreditCard":
		creditCard := strategy.CreditCard{}
		creditCard.Pay(amount)
	case "PayPal":
		payPal := strategy.PayPal{}
		payPal.Pay(amount)
	case "Crypto":
		crypto := strategy.Crypto{}
		crypto.Pay(amount)
	default:
		fmt.Println("Invalid payment method")
	}
	//END

	// WITH THE CONTEXT
	// Centralized decision logic
	paymentContext := strategy.PaymentContext{}

	switch paymentType {
	case "CreditCard":
		paymentContext.SetPaymentMethod(strategy.CreditCard{})
	case "PayPal":
		paymentContext.SetPaymentMethod(strategy.PayPal{})
	case "Crypto":
		paymentContext.SetPaymentMethod(strategy.Crypto{})
	default:
		fmt.Println("Invalid payment method")
		return
	}

	//client code
	// Single point of payment execution
	paymentContext.Pay(amount)

	// Later in the code, you can change strategy and reuse the context
	paymentContext.SetPaymentMethod(strategy.Crypto{})
	paymentContext.Pay(50.0)

	// END
}
