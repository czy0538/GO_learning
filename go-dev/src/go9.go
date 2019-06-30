package main

import "fmt"
import "time"

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello,%s.\n", name)
		}()
		time.Sleep(10 * time.Nanosecond)
	}
}
