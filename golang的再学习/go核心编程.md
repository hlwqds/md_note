# go 核心编程

## 1 基础知识

### 1.1 语言简介 

#### 1.1.1 go语言的诞生背景

go语言的诞生主要基于如下原因:

1. 摩尔定律接近失效后多核服务器已经成为主流，当前的编程语言对并发的支持不是很好，不能很好地发挥多核cpu的威力
2. 程序规模越来越大，编译速度越来越慢，如何快速地编译程序是程序员的迫切需求
3. 现有的编程语言设计越来越复杂

#### 1.1.2 语言特性

##### 语言组织

- 标识符和关键字
- 变量和常量
- 运算符
- 表达式
- 简单语句
- 控制结构

##### 类型系统

- 动静特性
- 类型强弱
- 基本数据类型
- 自定义数据类型

##### 抽象特性

- 函数：是否支持函数、匿名函数、高阶函数、闭包等
- 面向对象
- 多态
- 接口

##### 元编程特性

- 泛型
- 反射

##### 运行和跨平台语言特性

- 编译模式：是编译成可执行文件，还是编译成中间代码，还是解释器解释执行
- 运行模式
- 内存管理
- 并发支持
- 交叉编译
- 跨平台支持

##### 语言软实力特性

- 库
- 框架
- 语言自身兼容性
- 语言影响力

#### 1.1.3 go语言的特性

### 1.2 初识go语言

```go
//定义一个包，包名为main,main是可执行程序的包名，所有的Go源程序文件头部必须有一个包申明语句，Go通过包来管理命名空间
package main
//引用一个外部包fmt，可以是标准库的包，也可以是第三方或者自定义的包。fmt是标准输入\输出包，即format的缩写
import "fmt"

//使用ｆｕｎｃ关键字申明定义一个函数，函数名为main，main代表Go程序入口函数
func main(){
    //调用fmt包里的Printf函数，实参是一个字符串字面量
	fmt.Printf("Hello, world!\n")
}
```

#### Go源代码的特征解读

- 源程序以.go为后缀
- 源程序默认为UTF-8编码
- 标识符区分大小写
- 语句结尾的分号可以省略
- 函数以func开头，函数开头的"{"必须在函数头所在的行尾，不能单独起一行
- 调用包里的方法通过"."访问符
- main函数所在的包名必须是main

#### 编译运行

```bash
#编译
go build hello.go

#运行
./hello
```

### 1.3 Go词法单元

在介绍Go语言具体语法之前，先介绍现代高级语言的源程序内部的几个概念：token、关键字、标识符、操作符、分隔符和字面量

#### 1.3.1 token

token是构成源程序的基本不可再分割的单元。

Go的token分隔符有两类：一类是操作符，还有一类自身没有特殊含义，仅用来分隔其他token，被称为纯分隔符

- 操作符：操作符就是一个天然的分隔符，同时其自身也是一个token
- 纯分隔符：其本身不具备任何语法含义，只作为其他token的分割功能。包括空格、制表符、换行符和回车符

#### 1.3.2 标识符

编程语言的标识符用来标识变量、类型、常量等语法对象的符号名称，其在语法分析时作为一个token存在。编程语言的标识符总体上可以分为两类：语言设计这预留的标识符和编程时自定义的标识符。

Go语言预申明的标识符包括关键字、内置数据类型标识符、常量值标识符、内置函数和空白标识符。

编程语言中的关键字是指语言设计者保留的有特定语法含义的标识符，这些关键字有自己独特的用途和含义，它们一般用来控制程序结构，每个关键字都代表不同语义的语法糖

- 引导程序整体结构的８个关键字

```go
package //定义包名的关键字
import //导入包名的关键字
const //常量申明关键字
var //变量申明关键字variable
func //函数定义关键字
defer //延迟执行关键字
go //并发语法糖关键字
return //函数返回关键字
```

- 申明复合数据结构的４个关键字

```go
struct //定义结构类型关键字
interface //定义接口类型关键字
map	//申明或创建map类型关键字
chan //申明或创建通道类型关键字
```

- 控制程序结构的13个关键字

```go
if else //if else 语句关键字
for range break continue //for 循环使用的关键字
switch select type case default fallthrough //switch和select语句使用的关键字
goto //goto 跳转语句关键字
```

##### 内置数据类型标识符

丰富的内置类型支持是高级语言的基本特性，基本类型也是构造用户自定义类型的基础，为了标识每种内置数据类型，Go定义了一套预申明标识符，这些标识符用在变量或者常量申明中

```go
数值
	整型
		byte int int8 int16 int32 int64
		uint uint8 uint16 uint32 uint64 uintptr
	浮点型
		float32 float64
	复数型
		complex64 complex128
	字符和字符串型
		string rune
	接口型
		error
	布尔型
		bool
```

Go是一种强类型静态编译型语言。在定义变量和常量是需要显示支出数据类型，当然Go也支持自动类型推导，在申明初始化内置类型变量是，Go可以自动地进行类型推导，但是在定义新类型或函数时，必须显式地带上类型标识符。



##### 内置函数

```go
make new len cap append copy delete panic recover close complex real image print println
```

内置函数也是高级语言的一种语法糖，由于其是语言内置的，不需要通过import引入，内置函数具有全局可见性。



##### 常量

```go
true false	//布尔类型的两个常量值
iota		//用在连续的枚举类型的申明中
nil			//指针/引用型的变量值的默认指是nil
```

##### 空白标识符

```go
_
```

空白标识符具有特殊的含义，用来申明一个匿名的变量，该变量在赋值表达式的左端，空白标识符引用通常被用作占位，比如忽略函数多个返回值中的一个和强制编译器做类型检查

#### 1.3.3 操作符和分隔符

操作符就是语言使用的符号合集，包括运算符、显示的分隔符，以及其他语法辅助符号

Go语言一共有47个操作符：

- 算术运算符

```go
+ - * / %
```

- 位运算符

```go
& | ^ &^ >> <<
```

- 赋值和赋值复核运算符

```go
:= = += -= *= /= %= &= |= ^= &^= >>= <<=
```

- 比较运算符

```go
> >= < <= == ~=
```

- 括号

```go
() {} []
```

- 逻辑运算符

```go
&& || !
```

- 自增自减运算符

```go
++ --
```

- 其他运算符

```
: , ; . ... <-
```

#### 1.3.4 字面常量

编程语言源程序中表示固定值的符号叫做字面常量，简称字面量。一般使用裸字符序列来表示不同类型的值。字面量可以被编程语言直接转换为某个类型的值。Go的字面量可以出现在两个地方：一是用于常量和变量的初始化，二是用在表达式里或作为函数调用实参。变量初始化语句中如果没有显示指定变量类型，则Go编译器会结合字面量的值自动进行类型推断。Go中的字面量只能表达基本类型的值，不支持用户定义的字面量。

字面量有如下几类：

- 整形字面量

```go
42
0600
0x241
```

整型字面量使用特定字符序列来表示具体的整型数值，长用于整型变量或者常量的初始化

- 浮点字面量类型

```go
0.2
1E6
```

浮点型字面量是用特定字符序列来表示一个浮点数值。它支持两种格式：一种是标准的数学记录法，一种是科学计数法

- 复数类型字面量

```go
0i
011i
```

- 字符型字面量

Go的源码采用的是utf-8的编码方式，UTF-8的字符占用的字节数可以有1~4个字节，Rune字符常量也有多种表现形式，但是使用''将其括住

```go
'a'
'本'
```

- 字符串型字面量

```go
"\n"
```

字符串字面量的基本表现形式就是使用“”将字符序列包括在内

### 1.4 变量和常量

高级语言通过一个标识符来绑定一块特定的内存，后续对特定内存的操作都可以使用标识符来代替。这类绑定某个存储单元的标识符又可以分为两类，一类称之为变量，一类称之为常量。

把对地址的操作和引用变为对变量的操作是编程领域的巨大进步，它一方面简化了编写，记住标识符比记住某个地址更容易，另一方面极大地提升了程序的可读性。

#### 1.4.1 变量

变量：使用一个名称来绑定一块内存地址，该内存地址中存放的数据类型由定义变量时制定的类型型决定，该内存地址里面存放的内容可以改变。

Go的基本类型变量申明有两种：

1. 显示的完整申明

```go
var varName dataType [ = value]
```

说明：

- 关键字var用于变量申明
- varName是变量申明标识符
- dataType是1.3节介绍的基本类型
- value是变量的初识值，初始值可以使字面量，也可以是变量，也可以是表达式；如果不指定初始值，则Go默认将变量初始化为类型的零值
- Go的变量申明后就会立即为期分配空间

2. 短类型申明

```go
varName := value
```

- := 申明只能出现在函数内（包括方法）
- 此时Go编译器自动进行数据类型推断

Go支持多个类型变量同时申明并赋值

```go
a, b := 12,"dwww"
```

变量具有如下属性：

- 变量名
- 变量值
- 变量存储和生存期

Go语言提供自动内存管理，通常程序员不需要特别关注变量的生存期和存放位置。编译器使用栈逃逸技术能够自动为变量分配空间：可能在栈上，可能在堆上。

- 类型信息
- 可见性和作用域

#### 1.4.2 常量

常量使用一个名称来绑定一块内存地址，该内存地址中存放的数据类型由定义常量时制定的类型决定，而且该内存地址里存放的内容不可以改变。Go中常量分为布尔型，字符串型和数值型常量。常量存储在程序的只读段里。

预申明标识符iota用在常量申明中，其初始值为0.一组多个常量同时申明时，其值逐行增加，iota可以看做自增的枚举变量，专门用来初始化常量。

### 1.5 基本数据类型

Go是一种强类型的静态编译语言，类型是高级语言的基础，有了类型，高级语言才能对不同类型抽象出不同的运算，编程这才能在更高的抽象层次上操纵数据，而不用关注具体存储和运算细节。

### 1.6 复合数据类型

#### 1.6.1 指针

Go语言支持指针，指针的申明类型为*T，Go同样支持多级指针**T。通过在变量名前加&来获取变量的地址。指针的特点如下：

（1）在赋值语句中*T出现在=的左边表示指针申明， *T出现在=右边表示指针指向的值

```go
var a = 10
p := &a
```

（2）结构体指针访问结构体字段时仍然使用.操作符。

（3）Go不支持指针的运算

（4）函数中允许返回局部变量的地址

Go编译器使用栈逃逸的机制，将这种局部变量的空间分配在堆上

#### 1.6.2 数组

数组的类型名是[n]elementType，其中n是数组长度，elementType是数组元素类型。数组一般在创建时通过字面量初始化。

##### 数组初始化

```go
a := [3]int{1,2,3}	//指定长度和初始化字面量
b := [...]int{1,2,3}	//不指定长度，由后面的初始化列表数量来确定其长度
c := [3]int{1:1, 3:3}	//指定长度，并通过索引值进行初始化，没有初始化的元素使用类型默认值
d := [...]int{1:1, 3:3}	//不指定长度，通过索引值进行初始化，长度由最后一个索引确定
```

#### 数组的特点

（1）数组创建完长度就固定了，不可追加长度

（2）数组是值类型的，数组赋值或作为函数参数都是值拷贝

（3）数组长度是数组类型的组成部分。[10]int和[20]int表示不同的类型

（4）可以根据数组创建切片

#### 数组的相关操作

（1）数组元素方位

```go
a := [...]int{1, 2, 3}
b := a[0]
for i,v : range a{
    
}
```

（2）数组长度

```go
a := [...]int{1, 2, 3}
alength := len(a)

for i:=0; i < alength; i++ {
    
}
```

#### 1.6.3 切片

Go语言的数组的定长性和值拷贝限制了其使用场景，Go提供了另一种数据类型slice，这是一种变长数组，其数据结构中有指向数组的指针

```go
type slice struct{
    array unsafe.Pointer
    len int
    cap int
}
```

Go为切片维护三个元素——指向底层数组的指针、切片的元素数量和底层数组的容量

（1）切片的创建

只能通过数组或者通过make函数进行创建

```go
var array = [...]int{0, 1, 2, 3, 4, 5, 6}
s1 := array[0:4]
s2 := array[:4]
s3 := array[2:]
fmt.Printf("%v\n", s1)
fmt.Printf("%v\n", s2)
fmt.Printf("%v\n", s3)
```

```go
//len = 10, cap = 10
a := make([]int ,10)

//len = 10, cap = 15
b := make([]int, 10, 15)
```

（2）切片支持的操作

- 内置函数len()返回切片长度
- 内置函数cap()返回切片底层数组容量
- 内置函数append()对切片追加元素
- 内置函数copy()用于复制一个切片

```go
a := [...]int{0, 1, 2, 3, 4, 5, 6}
b := make([]int, 2, 4)
c := a[0:3]

fmt.Println(len(b))
fmt.Println(cap(b))
b.append(b, 1)
fmt.Println(b)
fmt.Println(len(b))


```

#### 1.6.4 map

Go语言内置的字典类型叫map。map的类型格式是map[K]T，其中K是任意可以进行比较的类型，T是值类型。map也是一种引用类型

（1）map的创建

- 使用字面量进行创建

```go
ma := map[string]int{"a":2, "ad":3}
fmt.Println(ma["a"])
fmt.Println(ma["ad"])
```

- 使用内置的make函数创建

```go
make(map[K]T)	//map的容量使用默认值
make(map[K]T, len)	//map的容量使用给定的len值

mp1 = make(map[int]string)
mp2 = make(map[int]string, 10)
mp1[1] = "tom"
mp2[1] = "pony"
fmt.Println(mp1[1])
fmt.Println(mp2[1])
```

（2）map支持的操作

- map的单个键值的访问格式为mapName[key]
- 可以使用range遍历一个map类型变量，但是不能保证每次迭代的顺序
- 删除map中的某个键值，使用以下语法：delete(mapName, key)。delete是内置函数，用来删除map中的某个键值对
- 可以使用内置的len()函数返回map中的键值对数量

```go
mp := make(map[int]string)
mp[0] = "tom"
mp[1] = "pony"
mp[2] = "jaky"
mp[3] = "andes"
delete(mp, 3)

fmt.Println(mp[1])
fmt.Println(len(mp))

for k, v := range(mp){
    fmt.Printf("key=",k,"value=",v)
}
```

##### 注意

- Go内置的map不是并发安全的，并发安全的map可以使用标准包sync中的map
- 不要直接修改map value中某个元素的值，如果想要修改map的某个键值，则必须整体赋值

```go
type User struct{
    name string
    age	int
}

ma = make(map[int]User)
andes := User{
    name: "huanglin",
    age: 24,
}

ma[1] = andes
//ma[1].age = 19 //ERROR,不能通过map引用直接修改
andes.age = 19
ma[1] = andes
fmt.Printf("%v\n", ma)
```

#### 1.6.5 struct

Go中的struct类型和C类似，中文翻译为结构，由多个不同类型的元素组合而成。这里面有两层含义：第一，struct结构中的类型可以使任意类型；第二，struct的存储空间是连续的，其字段按照申明时的顺序存放

struct有两种形式，一种是struct类型字面量，一种是使用type申明的自定义struct类型

（1）struct类型字面量

```go
struct {
    FieldName FieldType
}
```

（2）自定义struct类型

自定义struct类型申明格式如下

```go
type TypeName struct{
    FieldName FieldType
}
```

实际使用struct字面量的场景并不多，更多的时候是通过type自定义一个新的类型来实现的。type是自定义类型的关键字，不但支持struct类型的创建，还支持任意其他自定义类型的创建

（3）struct类型的初始化

```go
type Person struct{
    Name string
    Age int
}

type Student struct{
    *Person
    Number int
}

p := &Person{
    Name: "huanglin",
    age: 24,
}

s := Student{
    Person: p,
    Number: 110,
}
```

### 1.7 控制结构

## 2 函数

几乎所有的高级语言都支持函数或者类似函数的编成结构。函数如此普遍和重要的一个原因是现代计算机进程执行模型大部分是基于堆栈的，而编译器不需要对函数做过多的转换就能让其在栈上运行，另一方面函数对代码的抽象程度适中，就像胶水，很容易将编程语言的不同层级的抽象体黏合起来。

Go不是一门纯函数的编程语言，但函数在Go中是第一公民，表现在：

- 函数是一种类型，可以像其他变量类型一样使用，可以作为其他函数的参数或返回值，也可以直接调用执行
- 函数支持多值返回
- 支持闭包
- 函数支持可变参数

### 2.1 基本概念

#### 2.1.1 函数定义

函数是Go程序源代码的基本构造单位，一个函数的定义包括以下几个部分：函数申明关键字func、函数名、参数列表、返回值列表和函数体。函数名遵循标识符的命名规则，首字母的大小决定该函数在其他包内的可见性：大写时其他包可见，小写时只有相同的包才可以访问：函数的参数和返回值需要使用()包裹

```go
func funcName(param-list)(result-list){
    function-body
}
```

##### 函数的特点

（1）函数可以没有输入参数也可以没有返回值

```go
func A(){
    //do something
}

func A ()(int){
    //do something
    ...
    return 1
}
```

（2）多个相同类型的参数可以使用简写模式。

```go
func add(a, b int) (int) {
    return a + b
}
```

（3）支持有名的返回值，参数名就相当于函数体内最外层的局部变量，命名返回值变量会被初始化为类型零值，最后的return可以不带参数名直接返回

```go
import "fmt"

func a() (int) {
	fmt.Println("Hello, world")
	return 1
}

func b(a, b int) (sum int){
	sum = a + b
	return
}

func main() (){
	ret := a()
	fmt.Println(ret)

	ret = b(1, 2)
	fmt.Println(ret)
}
```

（4）不支持默认值参数

（5）不支持函数重载

（6）不支持函数嵌套，严格地说是不支持命名函数地嵌套定义，但是支持嵌套匿名函数

```go
func add(a, b int) (sum int){
    anonymous := func(x, y int) (int){
        return x + y
    }
    return anonymous(a, b)
}
```

#### 2.1.2 多值返回

Go函数支持多值返回，定义多值返回的返回参数列表时要使用"()"包裹，支持命令参数的返回

```go
func swap(a, b int) (int,ret int){
    return a + b, ret
}
```

习惯用法：

如果多值返回值有错误类型，则一般将错误类型作为最后一个返回值

#### 2.1.3 实参到形参的传递

Go函数实参到形参的传递永远是值拷贝，又是函数调用后实参指向的值发生了变化，那是因为参数传递的是指针值的拷贝

```go
package main

import (
	"fmt"
)

func chvalue(a int) (int){
	a = a + 1
	return a
}

func chpointer(a *int) (){
	*a = *a + 1
	return
}

func main() (){
	var a int = 10
	chvalue(a)
	fmt.Println(a)

	chpointer(&a)
	fmt.Println(a)
}
```

#### 2.1.4 不定参数

Go函数支持不定数目的形式参数，不定参数申明使用param ...type的语法格式。

函数的不定参数有如下特点：

（1）所有补丁参数类型必须是相同的

（2）不定参数必须是函数的最后一个参数

（3）不定参数名在函数体内相当于切片，对切片的操作同样适合对不定参数的操作

```go
func sumList(arr ...int) (sum int){
	for _, v := range arr{
		sum += v
	}

	return sum
}
```

（4）切片可以作为参数传递给不定参数，切片名后要加上"..."。例如：

```go
	arr := [...]{1,2,3,4,5}
	slice := []int{1, 2, 3, 4}
	ret = sumList(slice...)
	fmt.Println(ret)
	ret = sumList(arr...)
	fmt.Println(ret)
```

（5）形参为不定参数的函数和形参为切片的函数类型不同

```go
func sumList(arr ...int) (sum int){
	for _, v := range arr{
		sum += v
	}

	return sum
}

func sumLista(arr []int) (sum int){
	for _, v := range arr{
		sum += v
	}

	return sum
}
fmt.Printf("%T\n", sumList);
fmt.Printf("%T\n", sumLista);
```

### 2.2 函数签名和匿名函数

#### 2.2.1 函数签名

函数类型又叫函数签名，一个函数的类型就是函数定义首行去掉函数名、参数名和{}，可以使用fmt.Printf的%T格式化参数打印函数的类型

```go
package main
import (
	"fmt"
)

func add(a, b int) (int){
	return a + b
}

func main() (){
	fmt.Printf("%T\n", add)
}
```

两个函数类型相同的条件是：拥有相同的形参列表和返回值列表（列表元素的次序、个数和类型都相同），形参名可以不同。

可以使用type定义函数类型，函数类型变量可以作为函数的参数或者返回值。

```go
package main
import (
	"fmt"
)

func add(a, b int) (int){
	return a + b
}

func sub(a, b int) (int){
	return a - b
}

type op func(int, int) (int)

func do(f op,a int,b int) (int){
	return f(a, b)
}

func main() (){
	fmt.Printf("%T\n", add)
	fmt.Println(do(add, 1, 2))
	fmt.Println(do(sub, 1, 2))
}
```

函数类型和map、slice、chan一样，实际函数类型变量和函数名都可被当做指针变量，该指针指向函数代码的开始位置。通常说函数类型变量是一种引用类型，未初始化的函数类型的变量的默认值是nil

Go中函数是第一公民。有名函数的函数名可以看做函数类型的常量，可以直接使用函数名调用函数，也可以直接赋值给函数类型变量，后续通过该变量来调用该函数

```go
package main

func sum(a, b int) (int){
    return a+b
}

func main() (){
    sum(3, 5)
    f := sum
    f(3, 4)
}
```

#### 2.2.2 匿名函数

Go提供两种函数：有名函数和匿名函数。匿名函数可以看做函数字面量，所有直接使用函数类型变量的地方都可以直接使用匿名函数代替。

匿名函数可以直接赋值给函数变量，可以当做实参，也可以作为返回值，还可以直接被调用

```go
package main

import (
	"fmt"
)

var sum = func(a, b int) int{
	return a + b
}

type Aa func(int, int) (int)

func doinput(f Aa, a, b int) (int){
	return f(a, b)
}

//匿名函数作为返回值
func wrap(op string) (Aa){
	switch op{
	case "add":
		return func(a, b int) (int){
			return a + b
		}
	case "sub":
		return func(a, b int) (int){
			return a - b
		}
	default:
		return nil
	}
}

func main() (){
	//直接调用匿名函数
	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	sum(1, 2)

	//匿名函数作为实参
	doinput(func(x, y int) (int){
		return x + y
	}, 3, 4)

	opFunc := wrap("add")
	result := opFunc(2, 3)

	fmt.Println(result)
}
```

### 2.3 defer

Go函数里提供了的defer关键字，可以注册多个延迟调用，这些调用以先进后（FILO）的顺序在函数返回前被执行。这有点类似于Java语言中异常处理中的finaly子句。defer常用于保证一些资源最终一定能够得到回收和释放。

```go
package main

func main() (){
	//先进后出
	defer func() (){
		println("first")		
	}()

	defer func() (){
		println("second")
	}()

	println("func body")
}
```

defer后面必须是函数或者方法的调用，不能是语句，否则会报expression in  defer must be function call错误

defer函数的实参在注册时通过值拷贝传递进去。下面的示例代码中，实参a的值在defer注册时通过值拷贝传递进去，后续语句a++并不会影响defer语句最后的输出结果

```go
func f() int{
    a := 0
    defer func(i int) (){
        println("defer i=", i) 
    }(a)
    
    a++
    return a
}
```

defer语句必须先注册后才能执行，如果defer位于return之后，则defer因为没有注册，不会执行

主动调用os.Exit(int)退出进程时，defer将不再被执行

```go
package main
import (
    "os"
)

func main() (){
    defer func() (){
        println("defer after os.Exit()")
    }()
    println("func body")
    
    os.Exit(1)
}
```

defer的好处是可以在一定程度上避免资源泄露，特别是在有很多return语句，有多个资源需要关闭的场景中，很容易漏掉资源的关闭操作

### 2.4 闭包（附带数据的行为）

闭包是由函数及其相关引用环境组合而成的实体，一般通过在匿名函数中引用外部函数的局部变量或包全局变量构成

闭包 = 函数+引用环境

闭包对闭包外的环境引入是直接引用，编译器检测到闭包，会将闭包引用的外部变量分配到堆上

如果函数返回的闭包引用了该函数的局部变量：

1. 多次调用该函数，返回的多个闭包所引用的外部变量是多个副本，原因是每次调用函数都会为局部变量分配内存
2. 用一个闭包函数多次，如果该闭包修改了其引用的外部变量，则每一次调用该闭包对该外部变量都有影响，因为闭包函数共享外部引用

```go
package main
func fa(a int) (func(i int) (int)){
	return func(i int) (int){
		println(&a, a)
		a = a + i
		return a
	}
}

func main() (){
	f := fa(1)
	//g引用的外部的闭包环境包括本次函数调用的形参a的值1
	g := fa(1)
	//g引用的外部的闭包环境包括本次函数调用的形参a的值1
	//此时f,g引用的闭包环境中的a的值并不是同一个， 而是两次函数调用产生的副本

	println(f(1))
	println(f(1))

	println(g(1))
	println(g(1))
}
```

如果函数返回的闭包引用了全局变量，则操作的是全局变量所在内存空间中的实际值。一般不推荐闭包和全局变量结合使用。

### 2.5 panic和recover

本节主要介绍panic和recover两个内置

#### 2.5.1 基本概念

panic和recover的函数签名如下：

```go
panic(i interface{})
recover () interface{}
```

引发panic有两种情况，一种是程序主动调用panic函数，另一种是程序产生运行时错误由运行时检测并且抛出。

发生panic后，程序会从调用panic的函数位置或发生panic的地方立即返回，逐层向上执行函数的defer语句，然后逐层打印函数调用堆栈，直到被recover捕获或运行到最外层函数而退出

panic的参数是一个空接口类型interface{}，所以任意类型的变量都可以传递给panic。调用panic的方法非常简单：panic(xxx)

panic不但可以在函数正常流程中抛出，在defer逻辑里也可以再次调用panic或抛出panic。defer里面的panic能够被后续执行的defer捕获。

recover()用来捕获panic，阻止panic继续向上传递。recover()和defer一起使用，但是recover()只有在defer后面的函数体内被直接调用才能捕获panic终止异常，否则返回nil，异常继续向外传递。

```go
//这个会捕获失败
defer recover()

//这个会捕获失败
defer fmt.Println(recover())

//这个嵌套两层也会捕获失败
defer func(){
    func(){
        println("defer inneer")
        recover()
    }()
}()

//如下的场景会被捕捉成功
defer func() {
    println("defer inner")
    recover()
}()

func except() {
    recover()
}()

func test(){
    defer except()
    panic("test panic")
}
```

可以有多个panic被抛出，连续多个panic的场景只能出现在延迟调用里面，否则不会出现多个panic被抛出的场景。但只有最后一次panic能被捕获。

```go
package main

import (
	"fmt"
)

func main() (){
	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
		panic("first defer panic")
	}()

	defer func() (){
		panic("second defer panic")
	}()

	println("func body")
}
```

包中init的函数引发的panic只能在init函数中捕获，在main中无法捕获，原因是init函数先于main执行。函数并不能捕获内部新启动的goroutine所抛出的panic

```go
package main

import (
	"time"
	"fmt"
)

func da() (){
	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	panic("panic da")
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
}

func db() (){
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
}

func do() (){
	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	go da()
	go db()
	time.Sleep(3 * time.Second)
}

func main() (){
	do()
}
```

#### 2.5.2 使用场景

什么情况下主动调用panic函数抛出panic

一般有两种情况：

（1）程序遇到了无法执行下去的错误，主动调用panic函数结束程序运行

（2）在调试程序时，通过主动调用panic实现快速退出，panic打印出的堆栈信息可以更快地定位错误

为了保证程序的健壮性，需要主动在程序的分支流程上调用recover函数拦截运行时错误

Go提供了两种处理错误的方式，一种是借助panic和recover的抛出捕获机制，另一种是使用error错误类型

### 2.6 错误处理

Go的错误处理涉及接口的相关知识

#### 2.6.1 error

Go语言内置错误接口类型error。任何类型只要实现Error() string方法，都可以传递error接口类型变量。Go语言典型的错误处理方式是将error作为函数最后一个返回值。在调用函数时，通过检测其返回值的error值是否为nil来进行错误处理

```go
type error interfaces {
    ERROR() string
}
```

错误处理的最实践：

- 在多个返回值的函数中，error通常作为函数的最后一个返回值
- 如果一个函数返回error类型变量，则先用if语句处理error!=nil的异常场景，正常逻辑放在if语句块的后面，保持代码平坦
- defer语句应该放到err判断的后面，不然有可能产生panic
- 在错误逐级向上传递的过程中，错误信息应该不断地丰富和完善，而不是简单地抛出下层调用的错误。

#### 2.6.2 错误和异常

异常和错误在现代编程语言中是一对使用混乱的词语，下面将错误和异常做一个区分。

广义上的错误：发生非期望的行为

狭义上的错误：发生非期望的已知行为，这里的已知是指错误的类型是预料并定义好的

异常：发生非期望的未知行为。这里的未知是指错误的类型不在预先定义的范围内

（1）程序发生错误导致程序不能容错继续执行，此时程序应该主动调用panic或由运行时抛出panic

（2）程序虽然发生错误，但是程序能够容错继续执行，此时应该使用错误返回值的方式处理错误，或者在可能产生运行是错误的非关键分支调用recover捕获panic

## 3 类型系统

Go语言的类型系统可以分为命名系统、非命名类型、底层类型、动态类型和静态类型

### 3.1 类型介绍

#### 3.1.1 命名类型和未命名类型

##### 命名类型

类型可以通过标识符来表示，这种类型称为命名类型。Go语言的基本类型中有20个预申明简单类型都是命名类型，Go语言还有一种命名类型，用户自定义类型

##### 未命名类型

一个类型由预申明类型、关键字和操作符组合而成，这个类型称为未命名类型。未命名类型又称为类型字面量

Go语言中的基本类型的复合类型：数组、切片、字典、通道、指针、函数字面量、结构和接口都属于类型字面量，也都是未命名类型

注意：前面所说的结构和接口是未命名类型，这里的结构和接口没有使用type定义

```go
package main

import (
	"fmt"
)

type Person struct{
	name string
	age int
}

func main() (){
	a := struct{
		name string
		age int
	}{"huanglin1", 18}

	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)

	b := Person{"huanglin2", 24}
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)
}
```

#### 3.1.2 底层类型

所有类型都有一个underlying type（底层类型）。底层类型的规则如下：

（1）预申明类型(Pre-declared type)和类型字面量(type literals)的底层类型是它们自身。

（2）自定义类型type newtype oldtype 中newtype的底层类型是逐层递归向下查找的

#### 3.1.3 类型相同和类型赋值

##### 类型相同

Go是强类型的语言，编译器在编译时会进行严格的类型校验。两个命名类型是否相同，参考如下：

（1）两个命名类型相同的条件是两个类型申明的语句完全相同

（2）命名类型和未命名类型永远不相同

（3）两个未命名类型相同的条件是他们的类型申明的字面量的结构相同，并且内部元素的类型相同

（4）通过类型别名语句申明的两个类型相同

##### 类型可直接进行赋值

不同类型的变量之间一般是不能直接相互赋值的，除非满足一定的条件。下面探讨类型可赋值的条件。

1. T1和T2的类型相同
2. T1和T2具有相同的底层类型，并且T1和T2中至少有一个是未命名类型
3. T2是接口类型，T1是具体类型，T1的方法集是T2方法集的超集
4. T1和T2都是通道类型，它们拥有相同的元素类型，并且T1和T2中至少有一个未命名类型
5. a是预申明标识符nil，T2是pointer、function、slice、map、channel、interface类型中的一个
6. a是一个字面常量值，可以用来表示类型T的值

```go
package main

import (
	"fmt"
)

type Map map[string] string

func (m Map) Print() {
	for _, v := range m{
		fmt.Println(v)
	}
}

type iMap Map

func (m iMap) Print() {
	for _, v := range m{
		fmt.Println(v)
	}
}

type slice []int
func (s slice) Print() {
	for _, v := range s{
		fmt.Println(v)
	}
}

func main() (){
	mp := make(map[string]string, 10)
	mp["hi"] = "huanglin"

	var mb Map = mp
	var ma iMap = mp

	ma.Print()
	mb.Print()

	var i interface{
		Print()
	} = ma
	i.Print()
	//它居然也能打印，数据是哪里来的
	
    //im与ma虽然拥有相同的底层类型，但是二者中没有一个是未命名类型，不能直接赋值，可以进行强制类型转换
    //var im iMap = ma
    var im iMap = (iMap) (ma)
    
	s1 := []int{1,2,3}
	var s2 slice
	s2 = s1
	s2.Print()
}
```

#### 3.1.4 类型强制转换

由于Go是强类型的语言，如果不满足自动转换的条件，则必须进行强制类型转换。任意两个不相干的类型如果进行强制转换，则必须符合一定的规则。强制类型的语法格式：var a T = (T) (b)，使用括号将类型和要转换的变量或表达式的值括起来。

非常量类型的变量x可以强制转化并传递给类型T，需要满足如下任一条件：

1. x可以直接赋值给T类型变量
2. x的类型和T具有相同的底层类型
3. x的类型和T都是未命名的指针类型，并且指针指向的类型具有相同的底层类型
4. x和T的类型都是整型或者都是浮点型
5. x和T的类型都是复数类型
6. x是整数值或[]byte类型的值，T是string类型
7. x是字符串，T是[]byte或[]rune

字符串和字符切片之间的转换最常见，实例如下：

```go
	s := "hello,世界"
	var a []byte
	a = ([]byte) (s)
	var b string
	b = (string) (a)
	var c []rune
	c = ([]rune) (s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
```

注意：

1. 数值类型和string类型之间的相互转化可能会造成值部分丢失；其他的转换仅是类型的转换，不会造成值的改变。string和数字之间的转换可使用标准库strconv
2. Go语言没有语言机制支持指针和interger之间的直接转换，可以使用标准包中的unsafe包进行处理



### 3.2 类型方法

为类型增加方法是Go语言实现面向对象编程的基础

#### 3.2.1 自定义类型

##### 自定义struct类型

struct类型是Go语言自定义类型的普遍形式，是Go语言类型扩展的基础，也是Go语言面向对象承载的基础

```go
//使用type自定义的结构类型属于命名类型
type xxx struct{
    name string
    age int
}

//结构字面量属于未命名类型
struct {
    name string
    age int
}
var s = struct{}{}
```

##### struct初始化

```go
type Person struct{
    name string
    age int
}
```

1. 按照字段顺序进行初始化

```go
//注意又是三种写法
a := Person{"huanglin", 18}
b := Preson{
    "huanglin",
    18,
}
c := Person{
    "huanglin",
    18}
```

这不是一种推荐的方法，一旦结构增加字段，则不得不修改顺序初始化语句

2. 指定字段名进行初始化

```go
a := Person{name:"huanglin", age:18}

b := Person{
    name:"huanglin",
    age:18
}

c := Person{
    name:"haunglin",
    age:18}
```

注意：如果上述两种结构的初始化语句的}独占一行，则最后一个字段末尾一定要带上逗号

3. 使用new创建内置函数，字段默认初始化为其类型的零值，返回值是指向结构的指针

```go
p := new(Person)
//此时name为"",age是0
```

这种方法不常用，一般使用struct都不会将所有字段初始化为零值。

4. 一次初始化一个字段

```go
p := Person{}
p.name = "huanglin"
p.age = 11
```

这种方法不常用，这是一种结构化的编程思维，没有封装，违背了struct本身抽象封装的理念

5. 使用构造函数进行初始化

这是一种推荐的方法，当结构发生变化时，构造函数可以屏蔽细节

##### 结构字段的特点

结构的字段可以是任意的类型，基本类型、接口类型、指针类型、函数类型都可以作为struct的字段。结构字段的类型名必须唯一，struct字段类型可以是普通类型，也可以是指针。另外，结构支持内嵌自身的指针，这也是实现树形和链表等复杂数据结构的基础

```go
//标准库 container/list

type Element struct{
    next, prev *Element
    list *list
    Value interface{}
}
```

##### 匿名字段

在定义struct的过程中，如果字段只给出字段类型，没有给出字段名，则称这样的字段为匿名字段。被 匿名嵌入的字段必须是命名类型或命名类型的指针，未命名类型不能作为匿名字段使用。匿名字段的字段名默认就是类型名，如果匿名字段是指针类型，则默认的字段名就是指针指向的类型名。

##### 自定义接口类型

前面介绍了Go语言的类型系统和自定义类型，仅适用类型对数据进行抽象和封装还是不够的，本接介绍Go语言的类型方法。Go语言的类型方法是一种对类型行为的封装。Go语言的方法非常纯粹，可以看做特殊类型的函数，其显式地将对象实例或指针作为函数的第一个参数，并且参数名可以自己制定，而不强制要求一定是this或self。这个对象实例或指针称为方法的接收者

为命名类型定义方法的语法格式如下：

```go
func (t TypeName)MethodName(Paramlist)(Returnlist){
    //method body
}

func (t *TypeName)MethodName(Paramlist)(Returnlist){
    //method body
}
```

说明：

- t是接收者，可以自由指定名称
- TypeName为命名类型的类型名
- MethodName为方法名
- Paramlist为形参列表
- Returnlist为返回值列表

```go
type SliceInt []int

func (s SliceInt) Sum() int{
    sum := 0
    for _, i := range s{
        sum += i
    }
    
    return sum
}
```

类型方法有如下特点：

1. 可以为命名类型增加方法（除了接口），非命名类型不能自定义方法
2. 为类型增加方法有一个限制，那就是方法和类型的定义必须在同一个包里
3. 方法的命名空间的可见性和变量一样，大写开头的方法可以在包外被访问，否则只能在包内访问
4. 使用type定义的自定义类型是一个新类型，新类型不能调用原有类型的方法，但是底层类型支持的运算可以被新类型继承

### 3.3 方法调用

#### 3.3.1 一般调用

类型方法的一般调用方式：

TypeInstanceName.MethodName(ParamList)

- TypeInstanceName:类型实例名或者指向实例的指针变量名
- MethodName:类型方法名
- ParamList:方法实参

```go
type T struct{
    a int
}

func (t *T) Get() (int){
    return t.a
}

func (t *T) Set(i int) (){
    t.a = i
}

t := T{}

t.Set(2)

t.Get()
```

#### 3.3.2 方法值（method value）

变量x的静态类型是T，M是类型T的一个方法，x.T被称为方法值。x.T是一个函数类型的变量，可以赋值给其他变量，并像普通的函数名一样使用。

```go
f := x.M
f(args...)
```

等价于

```go
x.M(args)
```

方法值其实是一个带有包的函数变量，其底层实现原理和带有闭包的匿名函数类似，接受值被隐式地绑定到方法值的闭包环境中。后续调用不需要显示地传递接收者

#### 3.3.3 方法表达式(method expression)

方法表达式相当于提供一种语法，将类型方法调用显式地转换为函数调用，接收者必须显示地传递进去

表达式T.Get()和T.Set()被称为方法表达式

#### 3.3.4 方法集(method set)

命名类型接收者有两种类型，一个是值类型，一个是指针类型。

无论接收者是什么类型，方法和函数的形参传递都是值拷贝。

在直接使用类型实例调用类型的方法时，无论值类型还是指针类型，都可以调用类型的方法，原因是编译器在编译期间能够识别出这种调用关系，做了自动的转换

#### 3.3.5 值调用和表达式调用的方法集

1. 通过类型字面量显示的进行值调用或表达式调用，不会进行类型自动转换
2. 表达式调用不会进行类型自动转换

### 3.4 组合和方法集

结构类型为Go提供了强大的类型扩展，主要体现在两个方面：第一，struct可以嵌入任意其他类型的字段；第二，struct可以嵌套自身的指针类型的字段。这两个特性决定了struct有着强大的表达力，几乎可以表示任意的数据字段。同时，结合结构类型的方法，”数据+方法“可以灵活地表达程序逻辑。

#### 3.4.1 组合

从前面讨论的命名类型的方法可知，使用type定义的新类型不会继承原有类型的方法，有一个特例就是命名结构类型，命名结构类型可以嵌套其他的命名类型的字段，外层的结构可以调用嵌入字段类型的方法，这种调用既可以是显式的调用，也可以是隐式的调用（你不早说）。

这就是Go的继承或者组合

struct中的组合非常灵活，可以表现为水平的字段扩展，由于struct可以嵌套其他struct字段，所以组合也可以分层次扩展。struct类型中的字段称为内嵌字段

##### 内嵌字段的初始化和访问

struct的字段访问使用点操作符"."，struct的字段可以嵌套很多层，只要内嵌的字段是唯一的即可，不需要使用全路径进行访问。

```go
package main

type X struct{
	a int
}

type Y struct{
	X
	b int
}

type Z struct{
	Y
	c int
}

func main() (){
	x := X{a:1}
	y := Y{
		X: x,
		b: 2,
	}
	z := Z{
		Y: y,
		c: 3,
	}
	println(z.a, z.Y.a, z.Y.X.a)

	z = Z{}
	z.a = 2
	println(z.a, z.Y.a, z.Y.X.a)
}
```

在struct的多层嵌套中，不同的嵌套层次可以有相同的字段，此时最好使用完全路径进行访问和初始化。在实际数据结构的定义中应该尽量避开使用相同的字段，以免在使用中出现歧义

##### 内嵌字段的方法调用

struct类型方法调用也使用点操作符，不同嵌套层次的字段可以有相同的方法，外层变量调用内嵌字段的方法时也可以像嵌套字段的访问一样使用简化模式。方法名相同时，外层会覆盖内层的方法。因为Go编译器优先从外往内逐层查找方法。

```go
package main

import (
	"fmt"
)

type X struct{
	a int
}

type Y struct{
	X
	b int
}

type Z struct{
	Y
	c int
}

func (x X) Print() (){
	fmt.Printf("In X, a=%d\n", x.a)
}

func (x X) PrintX() (){
	fmt.Printf("In X, a=%d\n", x.a)
}

func (y Y) Print() (){
	fmt.Printf("In Y, b=%d\n", y.b)
}

func (z Z) Print() (){
	fmt.Printf("In Z, c=%d\n", z.c)

	z.Y.Print()
	z.Y.X.PrintX()
}

func main() (){
	x := X{a: 1}
	y := Y{
		X: x,
		b: 2,
	}
	z := Z{
		Y: y,
		c: 3,
	}

	z.Print()
	z.PrintX()
	z.Y.PrintX()
}
```

不推荐在多层的struct类型中内嵌多个同名的字段，但是不反对struct定义和内嵌字段同名的方法，因为这提供了一种编程技术，使得struct能够重写内嵌字段的方法

#### 3.4.2 组合的方法集

### 3.5 函数类型

有名函数：像func FunctionName()语法格式定义的函数称为有名函数

匿名函数：定义时没有指定函数名的是匿名函数

#### 函数字面量类型

函数字面量类型的语法表达格式是func (InputTypeList) (OutputTypeList)，可以看出有名函数和匿名函数都属于函数字面量类型。

#### 函数命名类型

type NewFuncType FuncLiteral

#### 函数签名

所谓函数	签名就是有名函数或匿名函数的字面量类型

#### 函数申明

## 4 接口

接口是一个编程规约，也是一组方法签名的集合。Go的接口是非侵入式的设计，也就是说，一个具体类型实现接口不需要在语法上显式地申明，只要具体类型的方法集是接口方法集的超集，就代表该类型实现了接口，编译器会在编译时进行方法集的校验。接口是没有具体实现逻辑的也不能定义字段。

### 变量和实例

### 空接口

最常使用的接口字面量类型就是空接口interface{},由于空接口的方法集为空，所以任意类型都被认为实现了空接口，任意类型的实例都可以赋值或者传递给空接口，包括非命名类型的实例。

### 4.1 基本概念

#### 4.4.1 接口申明

Go语言的接口分为接口字面量类型和接口命名类型，接口的申明使用interface关键字

接口字面量类型的申明语法如下：

```go
interface {
    MethodSignature1
    MethodSignature2
}
```

接口命名类型使用type关键字申明：

```go
type InterfaceName interface {
    MethodSignature1
    MethodSignature2
}
```

使用接口字面量的场景很少，一般只有空接口interface{}类型变量的申明是才会使用。

接口定义大括号内可以是方法申明的集合，也可以嵌入另一个接口类型匿名字段，也可以是两者的混合。接口支持嵌入匿名接口字段，就是一个接口定义里可以包括其他接口，Go编译器会自动进行展开处理，有点类似C语言中宏的概念。

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

//下面三种申明是等价的
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Write(p []byte) (n int, err error)
}
```

##### 申明新接口类型的特点

1. 接口的命名一般都已er结尾
2. 接口定义的内部方法申明不需要func引导
3. 在接口定义中，只有方法申明，没有方法实现

#### 4.1.2 接口初始化

单纯地申明一个接口变量没有任何意义，接口只有被初始化为具体的类型时才有意义。接口作为一个胶水层或者抽象层，起到抽象和适配的作用。没有初始化的接口变量，其默认值是nil。

接口绑定具体类型的实例的过程称为接口初始化。接口变量支持两种直接初始化方法。

##### 实例赋值接口

如果具体类型实例的方法集是某个接口的方法集的超集，则称该具体类型实现了接口，可以将该具体类型的实例直接赋值给接口类型的变量，此时编译器会进行静态的类型检查。接口被初始化后，调用接口的方法就相当于调用接口绑定的具体类型的方法。这就是接口调用的语义。

##### 接口变量赋值接口变量

已经初始化的接口类型变量a直接赋值给另一种接口变量b，要求b的方法集是a的方法集的子集。此时Go编译器会在编译时进行方法集静态检查。这个过程也是接口初始化的一种方式

#### 4.1.3 接口方法调用

接口方法调用和普通的函数调用是有区别的。接口方法调用的最终地址是在运行期决定的，将具体类型变量赋值给接口后，会使用具体类型的方法指针初始化接口变量，当调用接口变量的方法时，实际上是间接地调用实例的方法。接口方法调用不是一种直接的调用，有一定的运行时开销。

直接调用未初始化的接口变量的方法会产生panic。

```go
package main

type Printer interface {
	Print() ()
}

type S struct{}

func (s S) Print() (){
	println("print")
}

func main() (){
	var i Printer
	//i.Print()
	//为初始化的接口调用会产生panic

	i = S{}
	i.Print()
}
```

#### 4.1.4 接口的动态类型和静态类型

#####  动态类型

接口绑定的具体实例的类型称为接口的动态类型。接口可以绑定不同类型的实例，所以接口的动态类型是随着其绑定的不同类型实例而发生变化的。

##### 静态类型

接口被定义时，其类型就已经被确定，这个类型被称为接口的静态类型。接口的静态类型在其定义时就被确定，静态类型的本质是及接口的方法签名集。两个接口如果方法签名集合相同，顺序可以不同，则这两个接口在语义上完全等价，他们之间不需要进行类型强制转换就可以相互赋值。

### 4.2 接口运算

接口是一个抽象的类型，接口像一层胶水，可以灵活地解耦软件的每一个层次，基于接口编程是Go语言的基本思想。

有时我们需要知道已经初始化的接口变量绑定具体实例是什么类型，以及这个具体实例是否还实现了其他接口，这就要用到接口类型断言和接口类型查询

#### 4.2.1 类型断言

接口类型断言的语法形式如下：

```go
i.(TypeName)
```

i必须是接口变量，如果是具体类型变量，则编译器将会报错。

TypeName可以是接口类型名，也可以是具体类型名。

##### 接口查询的两层语义

1. 如果TyepName是一个具体类型名，则类型断言用于判断接口变量i绑定的实例类型是否就是具体类型TypeName
2. 如果TypeName是一个接口类型名，则类型断言用于判断接口变量i绑定的实例是否同时实现了TypeName接口

##### 接口断言的两种语法表现

直接赋值模式：

o := i.(TypeName)

语义分析：

1. TypeName是具体类型名，此时如果接口i绑定的实例类型就是具体类型TypeName，则变量o的类型就是TypeName，o的值就是接口i绑定的实例的副本
2. TypeName是接口类型名，此时如果接口绑定的实例满足接口TyeName，则变量o的类型就是接口类型TypeName，o底层绑定的就是接口i绑定的实例的副本
3. 如果上述两种情况都不满足，则程序抛出panic

```go
package main
import (
	"fmt"
)

type Inter interface {
	Ping()
	Pong()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (s St) Ping(){
	println("ping")
}

func (s St) Pong(){
	println("pong")
}

func main() (){
	st := &St{Name: "huanglin"}
	var i interface{} = st
	o := i.(Inter)
	o.Ping()
	o.Pong()

//	p := i.(Anter)
//	p.String()

	s := i.(*St)
	s.Ping()
	s.Pong()
	fmt.Printf("%s\n", s.Name)
}
```

comma表达式模式如下：

```go
if o, ok := i.(TypeName); ok {
    
}
```

语义分析：

1. TypeName是具体类型名，此时如果接口i绑定的实例类型就是具体类型TypeName，则变量o的类型就是TypeName，o的值就是接口i绑定的实例的副本，ok为true
2. TypeName是接口类型名，此时如果接口绑定的实例满足接口TyeName，则变量o的类型就是接口类型TypeName，o底层绑定的就是接口i绑定的实例的副本，ok为true
3. 如果上述两个都不满足，则ok为false。o是TypeName的零值。

```go
package main

import (
	"fmt"
)

type Inner interface {
	Ping()
	Pong()
}

type Outer interface {
	Inner
	Print()
}

type St struct {
	Name string
}

func (s St)Ping(){
	println("Ping")
}

func (s St)Pong(){
	println("Pong")
}

func main() (){
	var i interface{} = St{Name: "huanglin"}
	
	if o, ok := i.(Inner); ok{
		o.Ping()
		o.Pong()
	}

	if o, ok := i.(Outer); ok{
		o.Ping()
		o.Pong()
		o.Print()
	}

	if o, ok := i.(St); ok{
		o.Ping()
		o.Pong()
		fmt.Printf("%s\n", o.Name)
	}
}
```

#### 4.2.2 类型查询

接口类型查询的语法格式如下：

```go
switch v := i.(type){
    case type1:
    	xxxx
    case type2:
    	xxxx
    default:
    	xxxx
}
```

##### 语义分析

接口查询有两层语义一个是查询一个接口变量底绑定的底层变量的具体类型是什么，二是查询接口变量绑定的具体类型是否还是实现了其他接口

1. i必须是接口类型

类型查询一定是对一个接口变量进行操作，i必须是接口变量。如果i是未初始化接口变量，则v的值是nil

2. case字句后面可以跟非接口类型名，亦可以跟接口类型名，匹配是按照case字句的顺序进行的

case后面跟着多个类型，使用逗号分隔。接口变量绑定的具体实例只要和其中一个类型匹配没直接使用o赋值给v

#### 4.2.3 接口有点和使用形式

##### 接口优点

1. 解耦复杂系统进行垂直和水平分割是常用的设计手段，在层与层之间使用接口进行抽象和解耦是一种好的编程策略。Go的非侵入式的接口使层与层之间的代码更加干净，具体类型和实现的接口之间不需要显示申明，增加了接口使用的自由度。
2. 实现泛型，由于现阶段Go语言还不支持泛型，使用空接口作为函数或方法参数能够用在需要泛型的场景中。

##### 接口使用形式

1. 作为结构内嵌字段
2. 作为函数或者方法的形参
3. 作为函数或方法的返回值
4. 作为其他接口定义的嵌入参数

### 4.3 空接口

#### 4.3.1 基本概念

没有任何方法的接口，我们称之为空接口。空接口表示为interface{}。系统中任何类型都符合空接口的要求。Go语言中的空接口有点向C语言中的void *，只不过void *是指针，而Go语言的空接口内部封装了指针而已。

#### 4.3.2 空接口的用途

##### 空接口和泛型

Go语言没有泛型，如果一个函数需要接收任意类型的参数，则参数类型可以使用空接口类型，这是一种弥补的一种手段

```go
//典型的就是fmt标准包里的print函数
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
```

##### 空接口和反射

空接口是反射实现的基础，反射库就是将相关具体的类型转换并赋值给空接口后才去处理

#### 4.3.3 空接口和nil

空接口并不是真的为空，接口有类型和值两个概念。

空接口有两个字段，一个是实例类型，另一个是指向绑定的实例的指针，只有两个都为nil时，空接口才为nil

## 5 并发

### 5.1 并发基础

#### 5.1.1 并发和并行

并发和并行是不同的两个概念：

- 并行意味着程序在任意时刻都是同时运行的
- 并发意味着程序在单位时间内是同时运行的

并行就是在任一粒度的时间内都具备同时执行的能力

并发是在规定的时间内多个请求都得到执行和处理，强调的是给外界的感觉，实际上内部可能是分时操作的。并发重在避免阻塞，使程序不会因为一个阻塞而停止处理。

#### 5.1.2 goroutine

操作系统可以进行线程和进程的调度，本身具备并发处理能力，但进程切换的代价还是过于高昂，进程切换需要保存现场，耗费较多的时间。如果应用程序能在用户层在构筑一级调度，将并发的粒度进一步降低，是不是可以更大限度的提升程序的运行效率呢？Go语言的额并发就是基于这一思想实现的

通过go关键字启动一个goroutine go例程。go关键字后必须跟一个函数

- 通过go+匿名函数形式启动goroutine

```go
package main

import (
	"time"
	"runtime"
)

func main() (){
	go func(){
		sum := 0
		for i := 0; i < 10000; i++{
			sum += i
		}

		println(sum)
		time.Sleep(1 * time.Second)

	}()

	println("NumGoroutine=", runtime.NumGoroutine())

	time.Sleep(5 * time.Second)
}
```

- 通过go+有名函数形式启动goroutine

```go
package main

import (
	"runtime"
	"time"
)

func sum() (){
	sum := 0
	for i := 0; i < 10000; i++{
		sum += i
	}

	println(sum)
	time.Sleep(1 * time.Second)
}

func main() (){
	go sum()

	println("NumGoroutine=", runtime.NumGoroutine())

	time.Sleep(5 * time.Second)
}
```

goroutine有如下特性：

- go的执行是非阻塞的，不会等待
- go后面的函数的返回值会被忽略
- 调度器不能保证多个goroutine的执行次序
- 没有父子goroutine的概念，所有goroutine是平等地被调度和执行的
- Go程序在执行时会单独为main函数创建一个goroutine，遇到其他go关键字时再去创建其他的goroutine
- Go没有暴露go routine id给用户，所以不能在一个goroutine里面显式地操作另一个goroutine,不过runtime包提供了一些函数访问和设置goroutine的相关信息

1. func GOMAXPROCS

   func GOMAXPROCS(n int) int用来设置或查询并发执行的goroutine数目，n大于1表示设置GOMAXPROCS值，否则表示查询当前的值

2. func Goexit

   func Goexit()是结束当前goroutine





### tmp

```

```

