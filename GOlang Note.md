# GOlang Note

[TOC]

## 一、go基础

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

- 采用utf8编码
- package声明所属包名称
- import导入标准库和第三方包
- 入口main函数没有参数和返回值，如果需要向main传入参数，可以在os.Args变量中保存
- **左大括号**必须和函数声明或控制结构放在同一行
- 最后的**分号** 可以省略，出现分号的典型位置是for循环及其类似位置 



### 2.变量

- 用var定义变量：默认初值为0值，避免意外

- 支持类型推断：根据初始值类型进行推断

  ```go
  var sum int
  var b int=1
  label :="string"
  /// var a,b *int///都是指针
  ```

### 3.if语句

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

### 4.switch

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

//for range可以同时返回集合中数据的索引和值
x:=[]int{100,101,102}
for i,v:=range x{
    fmt.Println(i,",",v)
}
//range右侧必须是array，slice，string，map，指向array指针，channel之一

```

### 6.函数







