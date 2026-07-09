package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	demonstrateUnbufferedChannel()
	demonstrateReceiveFromGoroutine()
	demonstrateDoneChannel()
	demonstrateBufferedChannel()
	demonstrateSelect()

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}

func demonstrateUnbufferedChannel() {
	messageChannel := make(chan string)

	// Unbuffered sends block until another goroutine receives the value.
	go func() {
		messageChannel <- "ping"
	}()

	fmt.Println(<-messageChannel)
}

func demonstrateReceiveFromGoroutine() {
	result := make(chan int)
	go addNumbers(result, 10, 20)
	fmt.Println("sum from goroutine:", <-result)
}

func addNumbers(result chan<- int, firstNumber, secondNumber int) {
	result <- firstNumber + secondNumber
}

func demonstrateDoneChannel() {
	done := make(chan bool)
	go runTask(done)
	<-done
}

func runTask(done chan<- bool) {
	defer func() { done <- true }()
	fmt.Println("Processing task")
}

func demonstrateBufferedChannel() {
	emailChannel := make(chan string, 3)
	emailChannel <- "test1@example.com"
	emailChannel <- "test2@example.com"
	emailChannel <- "test3@example.com"
	close(emailChannel)

	for email := range emailChannel {
		fmt.Println("Email:", email)
	}
}

func demonstrateSelect() {
	numberChannel := make(chan int)
	textChannel := make(chan string)

	go func() {
		numberChannel <- 10
	}()

	go func() {
		textChannel <- "done"
	}()

	for i := 0; i < 2; i++ {
		select {
		case number := <-numberChannel:
			fmt.Println("Received from number channel:", number)
		case text := <-textChannel:
			fmt.Println("Received from text channel:", text)
		}
	}
}
