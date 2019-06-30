package main

import (
	"ed/input"
	"ed/output"
	"ed/sum"
)

func main() {
	var a, n, d, s int
	a, n, d = input.Input(a, n, d)
	s = sum.Sum(a, n, d)
	output.Output(s)
}
