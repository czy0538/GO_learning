package main

import "fmt"

func main() {
	scoresl := make([]int, 0, 5)
	scoresl = append(scoresl, 89)
	scoresl = append(scoresl, 90)
	scoresl = append(scoresl, 88)
	scoresl = append(scoresl, 87)
	scoresl = append(scoresl, 86)
	fmt.Printf("before:%p:%v:len=%v:cap=%v\n", scoresl, scoresl, len(scoresl), cap(scoresl))
	scoresl = append(scoresl, 101)
	scoresl = append(scoresl, 101)
	scoresl = append(scoresl, 101)
	scoresl = append(scoresl, 101)
	scoresl = append(scoresl, 101)
	fmt.Printf("after1:%p:%v:len=%v:cap=%v\n", scoresl, scoresl, len(scoresl), cap(scoresl))
	scoresl = append(scoresl, 102)
	fmt.Printf("after2:%p:%v:len=%v:cap=%v", scoresl, scoresl, len(scoresl), cap(scoresl))

}
