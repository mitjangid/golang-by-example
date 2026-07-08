package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	changeOrderStatus(StatusReceived)
	changeOrderStatus(StatusPrepared)

	changeOrderStatusString(StatusStringReceived)
	changeOrderStatusString(StatusStringOutForDelivery)

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}

// Go does not have a dedicated enum keyword. Use typed constants instead.
type OrderStatus int
type OrderStatusString string

const (
	StatusReceived OrderStatus = iota
	StatusConfirmed
	StatusPrepared
	StatusOutForDelivery
	StatusDelivered
)

const (
	StatusStringReceived       OrderStatusString = "Received"
	StatusStringConfirmed      OrderStatusString = "Confirmed"
	StatusStringPrepared       OrderStatusString = "Prepared"
	StatusStringOutForDelivery OrderStatusString = "Out For Delivery"
	StatusStringDelivered      OrderStatusString = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println(status)
}

func changeOrderStatusString(status OrderStatusString) {
	fmt.Println(status)
}
