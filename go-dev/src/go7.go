package main

import "fmt"
import "runtime"

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello,%s.\n", name)
		}()
	}
	runtime.Gosched()
}
