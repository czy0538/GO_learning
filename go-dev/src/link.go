package main

type node struct {
	data int
	next *node
}
type linkelist *node

func NewLinklist() linkelist {
	return linkelist{0, nil}
}
func main() {
	head := NewLinklist()
}
