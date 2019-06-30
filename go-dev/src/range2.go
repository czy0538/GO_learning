package main

import "fmt"

func main() {
	var seasons = []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, val := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, val)
	}
}
