package main

import "fmt"

import "poi"

func main() {
	p := new(poi.Point)
	p.X = 3
	p.Y = 4
	fmt.Println(p)

}
