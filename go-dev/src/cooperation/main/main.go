package main

import "cooperation/input"
import "cooperation/process"
import "cooperation/output"

func main() {
	var a, n, d, r int
	a, n, d = input.Input(a, n, d)
	r = process.Process(a, n, d)
	output.Output(r)
}
