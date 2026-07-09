package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	demonstrateGoroutines()
	demonstrateWaitGroups()

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}

// Goroutines are lightweight concurrent functions managed by the Go runtime.
func printTask(taskName string, id int) {
	fmt.Printf("%s processed item %d\n", taskName, id)
}

func demonstrateGoroutines() {
	fmt.Println("Starting goroutines")

	count := 5
	for i := 0; i < count; i++ {
		go printTask("background task", i+1)

		func(itemID int) {
			fmt.Println("regular function call:", itemID)
		}(i)
	}

	// Sleep is only used here so the beginner demo can show goroutine output.
	// Prefer sync.WaitGroup or channels in real code.
	time.Sleep(50 * time.Millisecond)
}

func processWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("waitgroup item:", id)
}

func demonstrateWaitGroups() {
	// A WaitGroup waits until a collection of goroutines finishes.
	count := 10
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go processWithWaitGroup(i+1, &wg)
	}

	wg.Wait()
	fmt.Println("All waitgroup items finished")
}
