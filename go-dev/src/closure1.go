//closure1.go
package main

import "fmt"

func main() {
	func() {
		var sum int = 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
		fmt.Println("the res is ", sum)
	}()
}
