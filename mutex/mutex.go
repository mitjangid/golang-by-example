package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	article := post{views: 0}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go article.incrementViews(&wg)
	}

	wg.Wait()
	fmt.Println(article.views)

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}

// A mutex protects shared data so only one goroutine can modify it at a time.
type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer wg.Done()

	p.mu.Lock()
	defer p.mu.Unlock()

	p.views++
}
