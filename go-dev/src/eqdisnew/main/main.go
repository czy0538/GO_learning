package main

import (
	"eqdisnew/input"
	"eqdisnew/out"
	"eqdisnew/sum"
)

func main() {
	var a, n, d, s int
	a, n, d = input.Input(a, n, d)
	s = sum.Sum(a, n, d)
	out.Output(s)
}
