GOlang Note

[TOC]

## 一、程序结构

### 1.基础格式

```go	
package main
import(
		"fmt"
)

func main(){
    fmt.Println("Hello world")
}
```

- 采用utf-8编码
- package声明所属包名称
- import导入标准库和第三方包
- 入口main函数没有参数和返回值，如果需要向main传入参数，可以在os.Args变量中保存
- **左大括号**必须和函数声明或控制结构放在同一行
- 最后的**分号** 可以省略，出现分号的典型位置是for循环及其类似位置 
- %T 类型，%v输出数据，%p地址 



### 2.变量

- 用var定义变量：默认初值为0值，避免意外

- 支持类型推断：根据初始值类型进行推断

- 支持声明列表

- 声明 var name type =expression

- 短声明 name :=expression

  - 短声明不需要声明所有的左边的变量
  - 短变量声明最少声明一个新变量
  
  ```go
  var sum int
  var b int=1
  label :="string"
  var a,b int=1,2
  /// var a,b *int///都是指针
  ```



### 3.指针

- 指针的值是变量的地址
- 指针类型的零值是nil
- 函数返回局部变量的地址是十分安全的，调用返回后仍然存在
- *p++ 递增p所指向的值

### 4.new函数

- new(T) 创建一个未命名的T类型变量，初始化T类型的零值，并返回其地址（地址类型为*T）
- 每一次调用new返回一个具有唯一地址的不同变量

### 5. 赋值与多重赋值

- 只有i++没有++i
- x，y=y，x（交换x,y的值）

### 6. 类型声明

- type name underlying-type
- 类似 using underlying-type=name

### 7.包和文件

- 
- 标识符可见：导出的标识符以大写字母开头
- 包的导入：import " package name

### 8.复数

- var x complex128 =complex(1,2)
- real imag 分别提取实部和虚部
- x:= 1+2i

### 9.字符串

- len函数返回字符串的字节数（而非文本符号数）
- s[i]取第i个字符
- s[i:j] 截取i到j（不包括j）的子串，默认值为0
- 字符串内部数据**不允许修改**，两个字符串安全的在底层共用一段内存
- 重要的包
  - strings：搜索、替换、比较、修整、切分、连接(相当于string.h)
  - bytes：操作字节slice
  - strconv：转换bool值、整数、浮点数
  - unicode：相当于ctype.h

### 10.常量

- const name=
- iota（enum）

## 二、语句

### 1.if语句

**表达式无需括起来**

```go
result:=someFunc()
if result>0{
    //do somethins
}else{
    //do something
}
```

- **大括号不可以省略，else不能单独成行，必须和if的右大括号在一起**
- **result全局有效**

```go
if result:=someFunc();result>0{
   //do something
}else{
    //do something
}
```

### 2 .come-ok模式

```go
//value, ok := element.(T)
// element必须是接口类型的变量，T是普通类型。如果断言失败，ok为false，否则ok为true并且value为变量的值。
package main

import (
    "fmt"
)

type Html []interface{}

func main() {
    html := make(Html, 5)
    html[0] = "div"
    html[1] = "span"
    html[2] = []byte("script")
    html[3] = "style"
    html[4] = "head"
    for index, element := range html {
        if value, ok := element.(string); ok {
            fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
        } else if value, ok := element.([]byte); ok {
            fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
        }
    }
}
```



### 3.switch

- 允许有相同的条件，多个条件之间用 , 分隔

- 无需break即可自动结束整个分支过程

- 分支条件不限于整数和字符，任何有效的表达式都可

- 分支值省略默认为真

  ```go
  switch result:=calculate();{
      case result <0:
      //do somethings
      case result >0:
      //do something
      default:
      //do something
  }
  ```

### 3.循环

**只有for循环**,且循环条件不用括号

```go
x:=0
for :x<5{
    fmt.Println(x)
    x++
}//相当于while

for x:=0;x<5;x++{
    fmt.Println(x)
}//相当于for循环

x:=4
for{
    fmt.Println(x)
   x--
    if x<0{
        break
    }
}//相当于do while

//for range可以同时返回集合中数据的索引和值，类似范围for循环
x:=[]int{100,101,102}
for i,v:=range x{
    fmt.Println(i,",",v)
}
//range右侧必须是array，slice，string，map，指向array指针，channel之一

```

## 三、复合数据类型

### 1.数组

- 具有固定长度且拥有0个或者多个相同数据类型元素的序列

- len返回数组中的元素个数

- var q[3] int

- var a=[3] int {1,2,3}

- q:= [...]int{1,2,3} 省略号表示数组长度由初始化数组的元素个数由初始化列表决定,仅能用于第一维数组

- 当传入函数时，接受的是一个**副本，而不是原始的数组

- 多维数组: n维数组由n-1维数组构成

  ```go 
  
  ```

- 数组和指针

  - 数组名不代表指针

  ```go
  x,y:=10,20
  a:=[2][2]int{{1,2},{3,4}}
  a:=[...]*int{&x,&y}
  p=&a
  fmt.Printf("%T,%v\n",a,a)
  fmt.Printf("%T,%p,%v,%v,%p\n",p,p,p,*p,&p)
  ```

- 允许数组的整体copy

- var x *[2] int x是指向数组的指针

- var x[2] *int x是指针数组

### 2.slice

- 相同类型元素的可变长度的序列

- []T

- len返回长度，cap返回容量

- s[i:j]截取一个子slice，包括i不包括j

- slice包含了指向数组元素的指针，即将slice传给函数时，传的是引用

- 无法直接比较，可以用bytes.Equal或者自己的函数来比较

- 允许和nil比较

- 检测是否为空应该用len(s)==0

- 建立制定长度的容量的slice ：make([]T,len,cap)

  ```go
  
  ```
  
  
  
- append(slice_name,T) 函数：追加到slice后面

  ```go
  //append函数实现,类似vector方式
  package main
  
  import "fmt"
  
  func appendslice(x []int, y ...int) []int {
  	var z []int
  	zlen := len(x) + len(y)
  	if zlen <= cap(x) {
  
  		z = x[:zlen]
  	} else {
  		
  		zcap := zlen
  		if zcap < 2*len(x) {
  			zcap = 2 * len(x)
  		}
  		z = make([]int, zlen, zcap)
  		copy(z, x)
  	}
  	copy(z[len(x):], y)
  	return z
  }
  
  func appendInt(x []int, y int) []int {
  	var z []int
  	zlen := len(x) + 1
  	if zlen <= cap(x) {
  
  		z = x[:zlen]
  	} else {
  		zcap := zlen
  		if zcap < 2*len(x) {
  			zcap = 2 * len(x)
  		}
  		z = make([]int, zlen, zcap)
  		copy(z, x) 
  	}
  	z[len(x)] = y
  	return z
  }
  
  func main() {
  	var x, y []int
  	for i := 0; i < 10; i++ {
  		y = appendInt(x, i)
  		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
  		x = y
  	}
  }
  
  x:=make([]int,5,5)
  	for i,v:=range x{
  		fmt.Println(i,v)
  	}
  0 0
  1 0
  2 0
  3 0
  4 0
  
  x:=make([]int,0,5)
  	for i,v:=range x{
  		fmt.Println(i,v)
  	}
  无任何输出
  
  ```



### 3 .map

- key-value对

- 散列表的引用

- key的类型必须可以通过 ==来比较

- make(map[key_type]value_type)

- 空map：map[string]int{}

- delete函数移除元素:delete(map_name,key_name),删除失败返回0

- 可以使用范围for遍历map，但是顺序是不定的

- 如果需要顺序需要sort

  ```go
  import"sort"
  var names[] string
  for name:= range ages{
      names =append(name,name)
  }
  sort.Strings(names)
  for_,name:= range names{
      fmt.Printf("%s\t%d\n",name,ages[name])
  }
  ```

  - map值类型本身可以是复合数据类型



### 4.方法
- 基本格式

  小写为私有，大写为公有

  ```go
  type student struct{
      stuName string
      age byte
  }
  
  stu1:=student{"小明"，16}
  stu1:=[2]student{{"xi",16},{"he",18}}
  stu1=student{age=16,stuName="小明"}
  ```

- 结构体嵌套和匿名成员

  - 允许结构体嵌套
  
  - 匿名成员：只需要指定类型名，必须是一个命名类型或者指向命名类型的指针
  
    ```go
    type Point struct{
        x,y,int
    }
    type Circle struct{
        Point
        Radius int
    }
    type Wheel struct{
        Circle
        Spokes int
    }
    var w Wheel
    {
        w.x=8//相当于w.Circle.Point.x=8
        w.y=8
        w.Radius=5
        w.Spokes=20
    }
    ```
  
    
  
- 方法

  ```go
  //私有方法
  type Point struct{
      float64 x,y
  }
  func (p Point) dist() float64{    //函数名和函数定义之间为所有者
  	return math.Sqrt(p.x*p.x + p.y*p.y)
  }
  
  //公有方法
  func (p Point) Dist() float64{    //函数名和函数定义之间为所有者
  	return math.Sqrt(p.x*p.x + p.y*p.y)
  }
  ```

- 通常用指针的方式使用

- 可比较性：其中所有的成员都是可比较的就是可比较的

- 接口：类似cpp的匿名函数

  ```go
  type Printer interface {
  Print()//公共行为
  /*无需显示给出实现类型*/
  }
  
  ```

  

### 6.函数

- 函数定义时，返回值类型放在参数表后，左大括号之前
- 允许多返回值

```go
//format:
func name(parameter_list)(result_list){
    body
}
func divide(a,b int)(int, int){
    quotient:=a/b
    remainder:=a%b
    return quotient,remainder
}

func divide(a,b,int)(quotient ,reminder int){
    quotient=a/b
    remainder=a%b
    return
}
```

#### 6.1 匿名函数

- 匿名函数可被当成一种数据类型来使用

- 匿名函数不能作为顶级函数，而是放在其他函数的函数体中，就是说**必须有外层函数**

  ```go	
  f:=func(a,b int) int{
      return a+b
  }
  sum:=f(2,3)
  ```

#### 6.2闭包

- 内部函数通过某种方式使其可见范围超出了其定义的范围

- go为声明匿名函数提供了简单的语法，像许多动态语言一样，这些匿名函数在他们被定义的范围内创建了词法闭

  ```go
  package main
  import "fmt"
  fun makeAddr(x int) (func(int)int){
      return func(y int) int{ return x+y}
  }
  func main(){
      add 5:=makeAdder(5)
      add36 :=makeAdder(36)
  }
  //另一种用法
    fun makeAddr(x int) (func(int)int){
        return func(y int) int{ return x+y}（5）
    }
  ```

- 闭包返回值是匿名函数
  
  - 匿名函数嵌套在闭包内部
  
- 匿名函数引用了自身的外部变量 
  
  

#### 6.3回调

函数可以作为其他函数的参数进行传递，然后在其他函数内调用执行，一般称之为回调

```go
package main
import "fmt"
func main(){
    callback(1,Add)
}
func Add(a,b int){
    fmt.Printf("The sum of %D and %d is:%d\n",a,b,a+b)
}

func callback(y int,f func(int,int)){
  f(y,2)//<=> Add(1,2)
}
```

#### 6.4 defer语句

延迟调用语句，无论函数执行是否出错，都确保结束前被调用	

```go
package main
import "main"
func main(){
    defer fmt.Println("The fist")
    fmt.Println("The second")
}
// The second
// The first

//用于数据清理工作，保证代码结构清晰，避免遗忘
scrFile,err:= os.Open(scrName)
defer srcFile.Close()
```



## 7文件操作

### 7.1 用户输入的读取

#### 7.1.1 fmt方法

- `Scanln` 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。

- `Scanf`与其类似，`Scanf` 的第一个参数用作格式字符串，用来决定如何读取。

- `Sscan` 和以 `Sscan` 开头的函数则是从字符串读取，除此之外，与 `Scanf` 相同。

  ```go
  // 从控制台读取输入:
  package main
  import "fmt"
  
  var (
     firstName, lastName, s string
     i int
     f float64
     input = "56.12 / 5212 / Go"
     format = "%f / %d / %s"
  )
  
  func main() {
     fmt.Println("Please enter your full name: ")
     fmt.Scanln(&firstName, &lastName)
     // fmt.Scanf("%s %s", &firstName, &lastName)
     fmt.Printf("Hi %s %s!\n", firstName, lastName) 
     fmt.Sscanf(input, format, &f, &i, &s)
     fmt.Println("From the string we read: ", f, i, s)
      
  }
  ```

  

#### 7.1.2 bufio方法

- `inputReader` 是一个指向 `bufio.Reader` 的指针。

  `inputReader := bufio.NewReader(os.Stdin)` 这行代码，将会创建一个读取器，类似文件指针

​	返回的读取器对象提供一个方法 `ReadString(delim byte)`，该方法从输入中读取内容，直到碰到 `delim` 指定的字符，	然后将读取到的内容连同 `delim` 字符一起放到缓冲区。

- `ReadString` 返回读取到的字符串，如果碰到错误则返回 `nil`。它一直读到文件结束，则返回读取到的字符串和 `io.EOF`。如果读取过程中没有碰到 `delim` 字符，将返回错误 `err != nil`。

- 屏幕是标准输出 `os.Stdout`；`os.Stderr` 用于显示错误信息，大多数情况下等同于 `os.Stdout`。

```go
package main
import (
    "fmt"
    "bufio"
    "os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
    inputReader = bufio.NewReader(os.Stdin)
    fmt.Println("Please enter some input: ")
    input, err = inputReader.ReadString('\n')
    if err == nil {
        fmt.Printf("The input was: %s\n", input)
    }
}
```



### 7.2文件读写

- 文件使用指向 `os.File` 类型的指针来表示的，也叫做文件句柄。

#### 7.2.1按行读取

- ReadString('\n') 或者 Readbytes('\n')或者ReadLine()

  ```go
  package main
  
  import (
      "bufio"
      "fmt"
      "io"
      "os"
  )
  
  func main() {
      inputFile, inputError := os.Open("input.dat")
      if inputError != nil {
          fmt.Printf("An error occurred on opening the inputfile\n" +
              "Does the file exist?\n" +
              "Have you got acces to it?\n")
          return 
      }
      defer inputFile.Close()
  
      inputReader := bufio.NewReader(inputFile)
      for {
          inputString, readerError := inputReader.ReadString('\n')
          fmt.Printf("The input was: %s", inputString)
          if readerError == io.EOF {
              return
          }      
      }
  }
  ```


  #### 7.2.2 整体读入字符串

  - 使用io/ioutil包里的ioutil.ReadFile()

  - 返回值[]byte,err

    

    ```go
    package main
    import (
        "fmt"
        "io/ioutil"
        "os"
    )
    
    func main() {
        inputFile := "hehe.txt"
        outputFile := "hehehe.txt"
        buf, err := ioutil.ReadFile(inputFile)
        if err != nil {
            fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
        }
        fmt.Printf("%s\n", string(buf))
        err = ioutil.WriteFile(outputFile, buf, 0644) 
        if err != nil {
            panic(err.Error())
        }
    }
    ```



#### 7.2.3 带缓存的读入

- bufio.Reader 的 Read()

  ```go
  buf := make([]byte, num)//num表示读取到的字节数
  ...
  n, err := inputReader.Read(buf)
  if (n == 0) { break}
  ```

#### 7.2.4 按列读取

- fmt的FSCan开头的函数

  ```go
  package main
  import (
      "fmt"
      "os"
  )
  
  func main() {
      file, err := os.Open("yyy.txt")
      if err != nil {
          panic(err)
      }
      defer file.Close()
  
      var col1, col2, col3 []string
      for {
          var v1, v2, v3 string
          _, err := fmt.Fscanln(file, &v1, &v2, &v3)
          if err != nil {
              break
          }
          col1 = append(col1, v1)
          col2 = append(col2, v2)
          col3 = append(col3, v3)
      }
  
      fmt.Println(col1)
      fmt.Println(col2)
      fmt.Println(col3)
  }
  ```

#### 7.2.5 压缩文件

- import"compress"

- 支持 bzip2,flate,gzip,lzw,zlib

  ``` go
  package main
  
  import (
      "fmt"
      "bufio"
      "os"
      "compress/gzip"
  )
  
  func main() {
      fName := "hehe.gz"
      var r *bufio.Reader
      fi, err := os.Open(fName)
      if err != nil {
          fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
              err)
          os.Exit(1)
      }
      fz, err := gzip.NewReader(fi)
      if err != nil {
          r = bufio.NewReader(fi)
      } else {
          r = bufio.NewReader(fz)
      }
  
      for {
          line, err := r.ReadString('\n')
          if err != nil {
              fmt.Println("Done reading file")
              os.Exit(0)
          }
          fmt.Println(line)
      }
  }
  ```

#### 7.2.6 带缓冲区的写文件

- OpenFile(filename,flag,权限)

- `os.O_RDONLY`：只读

- `os.O_WRONLY`：只写

- `os.O_CREATE`：创建：如果指定文件不存在，就创建该文件。

- `os.O_TRUNC`：截断：如果指定文件已存在，就将该文件的长度截为0。

  ```go
  package main
  
  import (
      "os"
      "bufio"
      "fmt"
  )
  
  func main () {
      outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
      if outputError != nil {
          fmt.Printf("An error occurred with file opening or creation\n")
          return  
      }
      defer outputFile.Close()
  
      outputWriter := bufio.NewWriter(outputFile)//创建一个写入缓冲区
      outputString := "hello world!\n"
  
      for i:=0; i<10; i++ {
          outputWriter.WriteString(outputString)
      }
      outputWriter.Flush()//将缓冲区内容完全写入文件
  }
  ```

#### 7.2.7 不带缓冲区写文件

- os.Stdout.WriteString ()  输出到屏幕

- f.WriteString() 直接写入文件

  ```go
  package main
  
  import "os"
  
  func main() {
      os.Stdout.WriteString("hello, world\n")
      f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
      defer f.Close()
      f.WriteString("hello, world in a file\n")
  }
  ```

### 1.如何分工

- 模块化编程，提高清晰度，促进团队合作水平，提高开发效率
  - 建立项目文件