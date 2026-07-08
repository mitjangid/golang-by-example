package main

import (
	"fmt"
	"time"
)

type order struct {
	id        int
	amount    float64
	status    string
	createdAt time.Time
	customer  // Embedded struct: order gets customer fields directly.
}

type customer struct {
	name    string
	address string
	pincode int
	email   string
	phone   int
}

func main() {
	start := time.Now()

	// Structs group related fields into one custom type.
	currentOrder := order{
		id:     23,
		amount: 240.50,
		status: "Order Placed",
	}

	currentOrder.createdAt = time.Now()

	currentOrder.updateStatusByValue("Received")
	fmt.Println("After value receiver:", currentOrder.status)

	currentOrder.updateStatus("Received")
	fmt.Println("After pointer receiver:", currentOrder.status)

	createdOrder := newOrder(12, 23.50, "Initiated")
	fmt.Println(createdOrder.id, createdOrder.amount, createdOrder.status)

	// Anonymous structs are useful for short-lived grouped data.
	language := struct {
		name    string
		origin  string
		country string
	}{"Hindi", "Sanskrit", "Bharat"}
	fmt.Println(language)

	createdCustomer := newCustomer("Amit", "Bikaner", 123456, "test@gmail.com", 9876543210)
	fmt.Println(createdCustomer.name, createdCustomer.email)

	createdOrder.customer = *createdCustomer
	fmt.Println(createdOrder.customer.email)
	fmt.Println(createdOrder.email) // Embedded fields can be promoted.

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}

func (o order) updateStatusByValue(status string) {
	// Value receivers work on a copy of the struct.
	o.status = status
}

func (o *order) updateStatus(status string) {
	// Pointer receivers can modify the original struct.
	o.status = status
}

func newOrder(id int, amount float64, status string) *order {
	createdOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &createdOrder
}

func newCustomer(name, address string, pincode int, email string, phone int) *customer {
	sampleCustomer := customer{
		name:    name,
		address: address,
		pincode: pincode,
		email:   email,
		phone:   phone,
	}
	return &sampleCustomer
}
