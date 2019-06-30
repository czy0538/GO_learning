package main

import "fmt"

func main() {
	var scoref float64
	var levelb byte
	fmt.Println("input scoref")
	fmt.Scanln(&scoref)
	switch int(scoref-50) / 10 {
	case 5:
		fallthrough
	case 4:
		levelb = 'A'
	case 3:
		levelb = 'B'
	case 2:
		levelb = 'C'
	case 1:
		levelb = 'D'
	default:
		levelb = 'E'

	}
	fmt.Printf("the result is %c", levelb)
}
