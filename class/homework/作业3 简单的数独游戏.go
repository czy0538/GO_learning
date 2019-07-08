package main

import "fmt"

func main() {
	var answer [6] int = [6]int{14, 15, 16, 8, 18, 19}
	var x[4] int=[4]int{8,18,14,15}
	fmt.Println("目标数据",answer)
	var cb [3][3] int
	for cb[1][1]=1;cb[1][1]<10;cb[1][1]++{
		for cb[1][2]=1;cb[1][2]<10;cb[1][2]++{
			for cb[2][1]=1;cb[2][1]<10;cb[2][1]++{
				for cb[2][2]=1;cb[2][2]<10;cb[2][2]++{
					cb[0][0]=x[2]+x[0]-45+cb[1][1]+cb[1][2]+cb[2][1]+cb[2][2]
					cb[0][1]=x[3]-cb[1][1]-cb[2][1]
					cb[0][2]=45-x[2]-x[3]-cb[1][2]-cb[2][2]
					cb[1][0]=x[1]-cb[1][1]-cb[1][2]
					cb[2][0]=45-x[0]-x[1]-cb[2][1]-cb[2][2]
				}
			}
		}
	}

}
