package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	razorpayPayment := payment{gateway: Razorpay{}}
	razorpayPayment.makePayment(260.50)

	stripePayment := payment{gateway: Stripe{}}
	stripePayment.makePayment(150.75)

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}

// Interfaces describe behavior. Any type with a matching method set satisfies
// the interface automatically.
type paymentGateway interface {
	pay(amount float64)
}

type payment struct {
	gateway paymentGateway
}

func (p payment) makePayment(amount float64) {
	p.gateway.pay(amount)
}

type Razorpay struct{}

func (Razorpay) pay(amount float64) {
	fmt.Println("Making payment using Razorpay:", amount)
}

type Stripe struct{}

func (Stripe) pay(amount float64) {
	fmt.Println("Making payment using Stripe:", amount)
}
