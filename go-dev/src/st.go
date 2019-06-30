package main

import "fmt"

type Sp struct {
	Name string
	Age  int
	Sex  byte //0-boy 1-girl
}

//func (s Sp) String() string {
//	return fmt.Sprintf("%s(%c,%v)", s.Name, s.Sex, s.Age)
//}
func main() {
	Sp1 := Sp{"旭日", 45, 'm'}
	//Sp2 := Sp{"阳刚", 46, 'm'}
	fmt.Printf("%s(%c,%v)", Sp1.Name, Sp1.Sex, Sp1.Age)

}
