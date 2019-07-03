# GOlang Note

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

- 每个包给他的声明提供独立的命名空间
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

### 2.switch

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

### 5.循环

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
- q:= [...]int{1,2,3} 省略号表示数组长度由初始化数组的元素个数金额U盾ing
- 当传入函数时，接受的是一个**副本，而不是原始的数组**

### 2.slice

- 相同类型元素的可变长度的序列
- []T
- len返回长度，cap返回容量
- s[i:j]截取一个子slice，包括i不包括j
- slice包含了指向数组元素的指针，即将slice传给函数时，传的是引用
- 无法直接比较，可以用bytes.Equal或者自己的函数来比较
- append(slice_name,T) 函数：追加到slice后面



### 6.函数

- 函数定义时，返回值类型放在参数表后，左大括号之前
- 允许多返回值

```go
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

### 6.1 匿名函数

- 匿名函数可被当成一种数据类型来使用

- 匿名函数不能作为顶级函数，而是放在其他函数的函数体中，就是说**必须有外层函数**

  ```go	
  f:=func(a,b int) int{
      return a+b
  }
  sum:=f(2,3)
  ```

### 6.3闭包

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
  ```

  - 闭包返回值是匿名函数
  - 匿名函数嵌套在闭包内部
  - 匿名函数引用了自身的外部变量 

  ### 6.4回调

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

  ### 6.5 defer语句

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

  





