/*package main

import "fmt"

func main() {
	deadlockTest()
}
func deadlockTest() {

	ch := make(chan int)
	results := make(chan int)

	for i := 0; i < 2; i++ {
		go func() {
			// 把从channel里取得的数据，再传回去
			x := <-ch
			results <- x

		}()
	}

	// 向输入数据里传两个数据
	ch <- 1
	ch <- 2

	for re := range results {
		fmt.Printf("re:%v\n", re)
	}

}*/

package main

import "fmt"
import "time"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1                                // 1流入信道，堵塞当前线, 没人取走数据信道不会打开
		fmt.Println("This line code wont run") //在此行执行之前Go就会报死
	}()
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	select {
	case <-ch:
		fmt.Println("正常退出")
	case <-timeout:
		fmt.Println("超时间退出")
	}
}

/*package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1                                // 1流入信道，堵塞当前线, 没人取走数据信道不会打开
	fmt.Println("This line code wont run") //在此行执行之前Go就会报死锁
}
*/
/*package main

import "fmt"

//import "runtime"
import "sync"

var wg sync.WaitGroup

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		wg.Add(1)
		go func(who string) {
			fmt.Printf("Hello %s\n", who)
			wg.Done()
		}(name)
	}
	wg.Wait()

}
*/
/*package main

import "fmt"
import "github.com/sbinet/go-python"

func init() {
   err := python.Initialize()
   if err != nil {
          panic(err.Error())
   }
}

func main() {
 	 gostr := "foo"
	 pystr := python.PyString_FromString(gostr)
	 str := python.PyString_AsString(pystr)
	 fmt.Println("hello [", str, "]")
}*/
/*package main

import "fmt"
import "time"

func main() {
	t0 := time.Now().UnixNano()
	fmt.Println("t0=", t0)
	for i := 0; i < 1000; i++ {
		fmt.Printf("")
	}
	t1 := time.Now().UnixNano()
	fmt.Println("\nt1=", t1)
	fmt.Println((t1 - t0))
}*/

/*package main

import (
	"fmt"
	"math/rand"
	"time"
)

//本方法参考群里面使用字典的方法生成"符合条件的数组"，但原程序有缺陷，生成的数组并不是可以直接填空，
//只是每行每列的和符合那6个数字而已，故本方法又添加了排序的部分程序
func main() {
	a := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().Unix())
	done := true //添加信号
	i := 0
	t0 := time.Now()
	for done {
		//使用洗牌算法生成九个互不相同的随机数
		rm := rand.Intn(9)
		a[i], a[rm] = a[rm], a[i]
		//定义字典，字典的键为对应位置数组元素的和，字典的值为true
		mp := make(map[int]bool)
		mp[a[0]+a[1]+a[2]] = true
		mp[a[3]+a[4]+a[5]] = true
		mp[a[6]+a[7]+a[8]] = true
		mp[a[0]+a[3]+a[6]] = true
		mp[a[1]+a[4]+a[7]] = true
		mp[a[2]+a[5]+a[8]] = true
		//假如随机排列的数组恰好可以使字典的键符合那六个和，说明找到了
		if mp[14] && mp[15] && mp[16] && mp[8] && mp[18] && mp[19] {
			m(a)         //使用m()函数排序
			done = false //退出循环
		} else if i < 8 {
			i++
		} else {
			i = 0
		}
	}
	t0s := time.Now().Sub(t0)
	fmt.Println("Time consuming:", t0s)
}
func m(a [9]int) {
	t1 := time.Now() //计算排序时间的起点

	b := 0 //求和变量
	k := 0 //累计变量

	matrix := [3][3]int{} //定义一个3*3数组，并将数组a[9]的元素依次赋值
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = a[k]
			k++
		}
	}

	//如果要填入，每行的和应该是8，18，19这些数字，而生成的matrix数组不一定是这样
	//以k作为是否需要转置的信号
	k = 0
	x := [3]int{} //以相邻的三个数为一组，将数组a分为三组，每组的和分别赋给数组x
	y := [3]int{} //以每相隔两个数的三个数为一组（如0，3，6），将数组a分为三组，每组的和分别赋给数组y
	for i := 0; i < 3; i++ {
		b = a[i*3] + a[1+i*3] + a[2+i*3]
		x[i] = b
		if b == 15 {
			k = 1
		}
	}
	for i := 0; i < 3; i++ {
		b = a[i] + a[3+i] + a[6+i]
		y[i] = b
	}
	switch k {
	case 1: //需要转置
		//因为转置以后行列变换，所以数组x,y也需要交换
		x, y = turna(x, y)
		matrix = turn(matrix)
		//转置后，根据数组x确定的三个数大小关系从小到大先排横排，再排竖排
		matrix = sortx(matrix, x)
		matrix = sorty(matrix, y)
		fmt.Println(matrix)
		t1s := time.Now().Sub(t1)
		fmt.Println("Sort time consuming:", t1s)
	case 0:
		matrix = sortx(matrix, x)
		matrix = sorty(matrix, y)
		fmt.Println(matrix)
		t1s := time.Now().Sub(t1)
		fmt.Println("Sort time consuming:", t1s)
	}
}
func sortx(a [3][3]int, x [3]int) [3][3]int {
	//数组x[3]三个数比较大小并按从小到大排横排，排好之后每行的和应依次是8，18，19
	if x[0] < x[1] {
		if x[1] < x[2] {
			return a
		} else if x[0] < x[2] {
			for i := 0; i < 3; i++ {
				a[1][i], a[2][i] = a[2][i], a[1][i]
			}
			return a
		} else {
			for i := 0; i < 3; i++ {
				a[0][i], a[2][i] = a[2][i], a[0][i]
			}
			for i := 0; i < 3; i++ {
				a[2][i], a[1][i] = a[1][i], a[2][i]
			}
			return a
		}
	} else if x[0] < x[2] {
		for i := 0; i < 3; i++ {
			a[0][i], a[1][i] = a[1][i], a[0][i]
		}
		return a
	} else if x[1] < x[2] {
		for i := 0; i < 3; i++ {
			a[1][i], a[0][i] = a[0][i], a[1][i]
		}
		for i := 0; i < 3; i++ {
			a[2][i], a[1][i] = a[1][i], a[2][i]
		}
		return a
	} else {
		for i := 0; i < 3; i++ {
			a[0][i], a[2][i] = a[2][i], a[0][i]
		}
		return a
	}
}
func sorty(a [3][3]int, y [3]int) [3][3]int {
	//数组y[3]三个数比较大小并按从小到大排竖列，排好之后每列的和应依次是14，15，16
	if y[0] < y[1] {
		if y[1] < y[2] {
			return a
		} else if y[0] < y[2] {
			for i := 0; i < 3; i++ {
				a[i][1], a[i][2] = a[i][2], a[i][1]
			}
			return a
		} else {
			for i := 0; i < 3; i++ {
				a[i][0], a[i][2] = a[i][2], a[i][0]
			}
			for i := 0; i < 3; i++ {
				a[i][2], a[i][1] = a[i][1], a[i][2]
			}
			return a
		}
	} else if y[0] < y[2] {
		for i := 0; i < 3; i++ {
			a[i][0], a[i][1] = a[i][1], a[i][0]
		}
		return a
	} else if y[1] < y[2] {
		for i := 0; i < 3; i++ {
			a[i][1], a[i][0] = a[i][0], a[i][1]
		}
		for i := 0; i < 3; i++ {
			a[i][2], a[i][1] = a[i][1], a[i][2]
		}
		return a
	} else {
		for i := 0; i < 3; i++ {
			a[i][0], a[i][2] = a[i][2], a[i][0]
		}
		return a
	}
}

//将数组x，y交换（需要转置的时候）
func turna(x [3]int, y [3]int) ([3]int, [3]int) {
	return y, x
}

//转置
func turn(a [3][3]int) [3][3]int {
	b := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b[j][i] = a[i][j]
		}
	}
	return b
}*/

/*package main

import "fmt"

func main() {
	a := [2]int{10, 2}
	func(x *[2]int) {
		x[1] += 100
		fmt.Printf("%v", *x)
	}(&a)
}
*/
/*package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := [7]int{}
	b := 0
	for i := 0; i < 7; i++ {
		a[i] = i + 1
	}
	rand.Seed(time.Now().Unix())
	for k := 0; k < 100000; k++ {
		for j := 0; j < 5; j++ {
			m := j + rand.Intn(7-j)
			fmt.Println("指数分布：", rand.ExpFloat64())
			a[j], a[m] = swap(a[j], a[m])
		}
		if (a[0] != 1) && (a[1] != 2) && (a[2] != 3) && (a[3] != 4) && (a[4] != 5) {
			b += 1
		}
	}
	fmt.Println(b)
}
func swap(a, b int) (int, int) {
	return b, a
}
*/
/*package main

import "fmt"

func main() {
	a := [10]int{}
	for i := 1; i < 10; i++ {
		a[i] = i
	}
	for i := 1; i < 6; i++ {
		for j := 1; j < 6; j++ {

			for z := 1; z < 6; z++ {

				if (a[i]+a[j]+a[z]) == 8 && z != j && z != i && i != j {
					a[1], a[i] = swap(a[1], a[i])
					a[2], a[j] = swap(a[2], a[j])
					a[3], a[z] = swap(a[3], a[z])
					fmt.Println(a[1:4], "i=", i, "j=", j, "z=", z)
				}
			}
		}
	}

}*/

/*package main

import "fmt"

func main() {
	a := [10]int{}
	for i := 1; i < 10; i++ {
		a[i] = i
	}
	for i := 1; i < 6; i++ {
		for j := 1; j < 6; j++ {

			for z := 1; z < 6; z++ {

				if (a[i]+a[j]+a[z]) == 8 && z != j && z != i && i != j {
					a[1], a[i] = swap(a[1], a[i])
					a[2], a[j] = swap(a[2], a[j])
					a[3], a[z] = swap(a[3], a[z])

					for q := 4; q < 10; q++ {

						for w := 4; w < 10; w++ {

							for e := 4; e < 10; e++ {

								if (a[q]+a[e]+a[w]) == 18 && e != w && e != q && w != q {
									a[4], a[q] = swap(a[4], a[q])
									a[5], a[w] = swap(a[5], a[w])
									a[6], a[e] = swap(a[6], a[e])

									for r := 7; r < 10; r++ {

										for y := 7; y < 10; y++ {

											for p := 7; p < 10; p++ {

												if a[r]+a[y]+a[p] == 19 && a[1]+a[4]+a[r] == 14 && a[2]+a[5]+a[y] == 15 && a[3]+a[6]+a[p] == 16 && p != r && p != y && r != y {

													fmt.Println(a[1], a[2], a[3])
													fmt.Println(a[4], a[5], a[6])
													fmt.Println(a[r], a[y], a[p])
													fmt.Println("")
												}
											}

										}

									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func swap(a, b int) (int, int) {
	return b, a
}*/

/*// shuiji数独 project main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	a := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().Unix())
	f := true
	i := 0
	t0 := time.Now()

	for f {
		rm := rand.Intn(9)
		a[i], a[rm] = a[rm], a[i]
		mp := make(map[int]bool)
		mp[a[0]+a[1]+a[2]] = true
		mp[a[3]+a[4]+a[5]] = true
		mp[a[6]+a[7]+a[8]] = true

		mp[a[0]+a[3]+a[6]] = true
		mp[a[1]+a[4]+a[7]] = true
		mp[a[2]+a[5]+a[8]] = true
		if mp[14] && mp[15] && mp[16] && mp[8] && mp[18] && mp[19] {
			fmt.Println(a)
			f = false
		} else if i < 8 {
			i++
		} else {
			i = 0
		}
	}
	t1 := time.Now().Sub(t0)
	fmt.Println("time consumed:", t1)

}*/

/*
// shuiji数独 project main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	a := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().Unix())
	f := true
	i := 0
	r1, r2, r3, c1, c2, c3 := 0, 0, 0, 0, 0, 0
	t0 := time.Now()
	for f {
		rm := rand.Intn(9)
		a[i], a[rm] = a[rm], a[i]

		r1 = a[0] + a[1] + a[2]
		r2 = a[3] + a[4] + a[5]
		r3 = a[6] + a[7] + a[8]

		c1 = a[0] + a[3] + a[6]
		c2 = a[1] + a[4] + a[7]
		c3 = a[2] + a[5] + a[8]
		if r1 == 14 && r2 == 15 && r3 == 16 && c1 == 8 && c2 == 18 && c3 == 19 {
			fmt.Println(a)
			f = false
		} else if i < 8 {
			i++
		} else {
			i = 0
		}
	}
	t1 := time.Now().Sub(t0)
	fmt.Println("time consumed:", t1)
}
*/
/*package main

import "fmt"

func main() {
	var x, y int = 3, 4
	swap(&x, &y)
	fmt.Println(x, y)
}
func swap(a *int, b *int) {
	var temp *int
	temp = a
	a = b
	b = temp
}*/

/*package main

import "fmt"

func main() {
	const (
		_          = iota
		KB float64 = 1 << (iota * 10)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(MB)
}
*/
/*package main

import "fmt"

func consumer(data chan int, done chan bool) {
	for x := range data {
		fmt.Println(x)
	}
	done <- true
}
func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
	}
	close(data)
}
func main() {
	done := make(chan bool)
	data := make(chan int)
	go producer(data)
	go consumer(data, done)
	<-done
}*/

/*package main

import "fmt"
import "time"

func main() {
	go taskone()
	go tasktwo()
	time.Sleep(time.Second * 6)
}
func taskone() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello")
		time.Sleep(time.Second)
	}
}
func tasktwo() {
	for i := 0; i < 5; i++ {
		fmt.Println("世界")
		time.Sleep(time.Second)
	}
}*/

/*package main

import "fmt"

type user struct {
	name string
	age  byte
}

func (u user) Output() string {
	return fmt.Sprintf("%+v", u)
}

type leader struct {
	user
	title string
}

func main() {
	var l leader
	l.name = "Tom"
	l.age = 30
	l.title = "sales manager"
	fmt.Println(l.Output())
}*/

/*package main

import "fmt"

type Printer interface {
	Print()
}

type user struct {
	name string
	age  byte
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

type goods struct {
	name  string
	price int
	num   int
}

func (g goods) Print() {
	fmt.Printf("%v\n", g)
}
func main() {
	u1 := user{"zs", 18}
	g1 := goods{"Tv", 3000, 10}
	u1.Print()
	g1.Print()
	var p Printer
	p = u1
	p.Print()
	p = g1
	p.Print()

}*/

/*package main

import "fmt"
import "math"

type Point struct {
	x float64
	y float64
}

func (p *Point) Length() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
func (p *Point) Scale(x float64) {
	p.x *= x
	p.y *= x
}
func main() {
	var p1 Point = Point{3, 4}
	fmt.Println(p1.x, ":", p1.y)
	fmt.Println(p1.Length())
	p1.Scale(2)
	fmt.Println(p1.Length())
}
*/
/*package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("程序异常退出:", err)
		} else {
			fmt.Println("程序正常退出")
		}
	}()
	f(101) //传入101
}
func f(a int) {
	if a > 100 {
		panic("参数值超出范围")
	} else {
		fmt.Println("函数f成功调用")
	}
}*/

/*package main

import "fmt"

func main() {
	var a = 'a'
	fmt.Println("input a:")
	//fmt.Scanf("%c", &a)
	fmt.Scanln(&a)
	fmt.Printf("varlue a=%v\n", a)
	fmt.Printf("chararcter a=%c", a)
}*/

/*package main

import "fmt"

func main() {
	var ch1 rune //rune是 unicode类型，相当于 int32
	ch1 = '中'    //unicode字符在内存中存的是字符对应的数字码，32位
	ch2 := 22269
	fmt.Println("input ch1:")
	fmt.Scanf("%c", &ch1)
	fmt.Println(ch1)                        //以内存中的数字形式输出
	fmt.Printf("ch1=%c,ch2=%v\n", ch1, ch2) //同上，输出数字形式
	fmt.Printf("%c%c\n", ch1, ch2)          //以unicode字符形式输出
	//下面是将ch1和ch2转换为字符串并连接一起输出，这里+为字符串连接
	str := string(ch1) + string(ch2)
	for i, v := range str {
		fmt.Println(i, v)
	}

	//fmt.Println(string(ch1) + string(ch2))
}*/
