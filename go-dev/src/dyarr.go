package main

import (
	"fmt"
)

func main() {
	var dyarr = make([][]int, 0, 100)
	var oned = make([]int, 0, 100)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			oned = append(oned, i*j)
		}
		dyarr = append(dyarr, oned[i*4:])
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", dyarr[i][j])
		}
		fmt.Println("")
	}
}

/*var dyarr = make([][]int, 0, 100)
for i := 0; i < 3; i++ {
	oned := make([]int, 0, 100)
	for j := 0; j < 3; j++ {
		oned = append(oned, i*j)
	}
	dyarr = append(dyarr, oned)
}
for i := 0; i < 3; i++ {
	for j := 0; j < 3; j++ {
		fmt.Printf("%d ", dyarr[i][j])
	}
	fmt.Println("")
}*/

/*#include <stdio.h>
void main(){
	int BeerBottle = 1000;
	int EmptyBottle = 0;
	for(int i=0;i<BeerBottle;i++){
		EmptyBottle = EmptyBottle+1;
		if(EmptyBottle == 3){
			BeerBottle = BeerBottle+1;
			EmptyBottle = 0;
		}
	}
	printf("EmptyBottle = %d  ",EmptyBottle);
	printf("BeerBottle = %d   ",BeerBottle);
}*/
