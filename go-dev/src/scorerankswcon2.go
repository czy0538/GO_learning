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
	switch int(scoref-50.0) / 10 { //以(scoref-50)的整数部分除以10作为分支条件
	case 5:	fallthrough
	case 4:	levelb = 'A'
	case 3:	levelb = 'B'
	case 2:	levelb = 'C'
	case 1:	levelb = 'D'
	default:	levelb = 'E'
	}
	fmt.Printf("The rank of %v is %c\n", scoref, levelb)
}
if initialzation;condition {
	//do something
}
if val := 10; val > max {
    // do something
}