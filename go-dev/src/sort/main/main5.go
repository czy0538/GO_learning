package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10)
	s2 := make([]int, 3, 10)
	for i, _ := range s1 {
		s1[i] = i
	}
	fmt.Println("before s1", s1)
	for i, _ := range s2 {
		s2[i] = i + 2
	}
	copy(s1, s2)
	fmt.Println("afeter s1", s1)
}
