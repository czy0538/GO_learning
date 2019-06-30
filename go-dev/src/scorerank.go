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
			continue //结束本次循环，重新录入
		} else {
			break //退出循环
		}
	}
	if scoref >= 90 {
		levelb = 'A'
	} else if scoref >= 80 {
		levelb = 'B'
	} else if scoref >= 70 {
		levelb = 'C'
	} else if scoref >= 60 {
		levelb = 'D'
	} else {
		levelb = 'E'
	}
	fmt.Printf("The rank of %v is %c\n", scoref, levelb)
}
