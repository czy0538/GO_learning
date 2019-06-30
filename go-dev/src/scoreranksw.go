package main

import "fmt"

func main() {
	var scoref float64 = 0.0
	var levelb byte
	for {
		fmt.Println("Please input a float to scoref")
		fmt.Scanln(&scoref)
		if scoref > 100 || scoref < 0 {
			fmt.Println("the value is out of scoref's fields")
			continue
		} else {
			break
		}
	}
	switch {
	case scoref <= 100 && scoref >= 90:
		levelb = 'A'
	case scoref < 90 && scoref >= 80:
		levelb = 'B'
	case scoref < 80 && scoref >= 70:
		levelb = 'C'
	case scoref < 70 && scoref >= 60:
		levelb = 'D'
	default:
		levelb = 'E'
	}
	fmt.Printf("The rank of %v is %c\n", scoref, levelb)
}
