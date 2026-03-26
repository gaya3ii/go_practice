package main

import (
	"fmt"
	"sync"
)

func fetchPrice(source string, ch chan string) {
	// fmt.Println("Fetching price from", source)
	ch <- fmt.Sprintf("Price from %s: $100", source)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 3)
	sources := []string{"Amazon", "Flipkart", "Myntra"}

	for _, source := range sources {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			fetchPrice(s, ch)
		}(source)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Println(result)
	}
}
