package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Printf("Input Char Is : %v", string([]byte(input)[0]))

	// fmt.Printf("You entered: %v", []byte(input))
}
