package main

import (
	"fmt"
	"sync"
)

func main() {
	names := []string{"Mark", "Tom", "Alice", "Mike", "Robert"}
	var wg sync.WaitGroup
	for _, name := range names {
		wg.Add(1)
		go func(who string) {
			fmt.Println("hello ", who)
			wg.Done()
		}(name)

	}
	wg.Wait()

}
