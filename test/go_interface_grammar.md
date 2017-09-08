

 - [接口 interface](#57f2ffaf14788e0050594f8ce0c6a134)
    - [6.1 接口定义](#6a72925f32afa3ff12617d249a0f86ff)
       - [接口是方法签名的集合](#692de9b0a54655802f83e4974bfe57ce)
        - [类型方法集中 拥有与接口对应的全部方法，就实现了该接口](#e66632238ab3e4115c7db3d54c0afd3c)
        - [接口命名习惯以 er 结尾，结构体](#e56b1ae641e9775362a66915aa5534df)
        - [接口中可以嵌入其他接口](#bb05456de1d5da54d0e099786a833846)
        - [空接口 interface{} 没有任何方法签名, 任何类型都实现了空接口](#529226f855a6548337ed66083f4c5ca8)
        - [匿名接口可用作变量类型, 或结构成员](#f17c0e3cb5964cfec0b082d50d8e36d2)
     - [6.2 执行机制](#d24b5171452124598b04254b5cb0572d)
        - [接口对象由接口表(interface table)指针 和数据指针组成](#b2d36ea79da4d01531b631b81ecfd3ec)
         - [数据指针持有的是目标对象的只读复制品, 复制完整对象 或指针](#e54c78ab99f2ef0f785e9d5061e99317)
         - [接口类型返回临时对象，只有指针才能修改其状态 (参考map 取值)](#11be92b9ff3c1d0e298e8a1434ce4351)
         - [只有 tab 和 data 都为 nil 时,接⼝才等于 nil](#0dad826b0a889e64861135de276fc1b9)
     - [6.3 接⼝转换](#da9e2eae4d985d6ab55ece99faf75d43)
        - [利用类型推断, 判断接口对象是否是 某个具体的接口或类型](#8397c802bcfa8ebabd50167063cd6766)
         - [用 switch 做批量 类型判断,不支持 fallthrough](#7ef91bde805ff5a99acc4869a3e892b1)
         - [超集接口对象 可以转为 子集接口， 反之出错](#e22517a0a99fccfc40fa94916db115ff)
     - [6.4 接口技巧](#d14e02b0e2f5b36b976f6adcdba535ee)
        - [让编译器检查，确保某个类型实现 接口](#e51931e1b204267a72cf1f21b8185b9c)
         - [让函数直接"实现"接口](#271a34f0fb6d790c4f6cb6f932c41bbe)

# 接口 interface

<h2 id="6a72925f32afa3ff12617d249a0f86ff"></h2>

## 6.1 接口定义

<h2 id="692de9b0a54655802f83e4974bfe57ce"></h2>

##### 接口是方法签名的集合

##### 类型方法集中 拥有与接口对应的全部方法，就实现了该接口

    所谓对应方法，是指 方法名，参数类型，返回值类型 都相同。

<h2 id="e56b1ae641e9775362a66915aa5534df"></h2>

##### 接口命名习惯以 er 结尾，结构体

<h2 id="bb05456de1d5da54d0e099786a833846"></h2>

##### 接口中可以嵌入其他接口

```go
type Stringer interface {
    String() string
}
type Printer interface {
    Stringer            // 借口嵌入
    Print() 
}


func (self *User) String() string {
    return fmt.Sprintf("user %d, %s", self.id, self.name)
}
func (self *User) Print() {
    fmt.Println(self.String())
}
func main() {
    var t Printer = &User{1, "Tom"}  // *User 实现了 Printer接口
    t.Print()       // user 1, Tom
}
```

<h2 id="529226f855a6548337ed66083f4c5ca8"></h2>

##### 空接口 interface{} 没有任何方法签名, 任何类型都实现了空接口 

    interface{} 作用类似其他语言中的根对象 object


```go
func Print(v interface{}) {
    fmt.Printf("%T: %v\n", v, v)
}
func main() {
    Print(1)                // int: 1
    Print("Hello, World!")  // string: Hello, World!
}
```

<h2 id="f17c0e3cb5964cfec0b082d50d8e36d2"></h2>

##### 匿名接口可用作变量类型, 或结构成员

```go
type Tester struct {
    s interface {
        String() string
    }
}
```

<h2 id="d24b5171452124598b04254b5cb0572d"></h2>

## 6.2 执行机制

<h2 id="b2d36ea79da4d01531b631b81ecfd3ec"></h2>

##### 接口对象由接口表(interface table)指针 和数据指针组成

```go
struct Iface
{
    Itab* tab;
    void*    data;
};
```

---
<h2 id="e54c78ab99f2ef0f785e9d5061e99317"></h2>

##### 数据指针持有的是目标对象的只读复制品, 复制完整对象 或指针

```go
type User struct {
    id   int
    name string 
}

func main() {
    u := User{1, "Tom"}
    var i interface{} = u   // 复制 结构体
    u.id = 2
    u.name = "Jack"
    fmt.Printf("%v\n", u)           // {2 Jack}
    fmt.Printf("%v\n", i.(User))    // {1 Tom}
}
```

---
<h2 id="11be92b9ff3c1d0e298e8a1434ce4351"></h2>

##### 接口类型返回临时对象，只有指针才能修改其状态 (参考map 取值)

```go
type User struct {
    id   int
    name string 
}
func main() {
    u := User{1, "Tom"}
    var vi, pi interface{} = u, &u
    // vi.(User).name = "Jack"   // cannot assign to vi.(User).name
    pi.(*User).name = "Jack"
    fmt.Printf("%v\n", vi.(User))   // {1 Tom}
    fmt.Printf("%v\n", pi.(*User))  // &{1 Jack}
}
```

---
<h2 id="0dad826b0a889e64861135de276fc1b9"></h2>

##### 只有 tab 和 data 都为 nil 时,接⼝才等于 nil

```go
var a interface{} = nil         // tab = nil, data = nil
var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil
```

<h2 id="da9e2eae4d985d6ab55ece99faf75d43"></h2>

## 6.3 接⼝转换

<h2 id="8397c802bcfa8ebabd50167063cd6766"></h2>

##### 利用类型推断, 判断接口对象是否是 某个具体的接口或类型

```go
type User struct {
    id   int
    name string
}
func (self *User) String() string {
    return fmt.Sprintf("%d, %s", self.id, self.name)
}
func main() {
    var o interface{} = &User{1, "Tom"}  
    if i, ok := o.(fmt.Stringer); ok {  // ok-idiom
        fmt.Println(i)   // 1, Tom 
    }
    u := o.(*User)
    // u := o.(User)  // panic: *main.User, not main.User
    fmt.Println(u)    // 1, Tome
}
```

---
<h2 id="7ef91bde805ff5a99acc4869a3e892b1"></h2>

##### 用 switch 做批量 类型判断,不支持 fallthrough

```go
func main() {
    var o interface{} = &User{1, "Tom"}
    switch v := o.(type) {
    case nil:                   // o == nil
        fmt.Println("nil")
    case fmt.Stringer:          // interface
        fmt.Println(v)
    case func() string:         // func
        fmt.Println(v())
    case *User:                 // *struct
        fmt.Printf("%d, %s\n", v.id, v.name)    
    default:
        fmt.Println("unknown")
    } 
}
```

<h2 id="e22517a0a99fccfc40fa94916db115ff"></h2>

##### 超集接口对象 可以转为 子集接口， 反之出错

<h2 id="d14e02b0e2f5b36b976f6adcdba535ee"></h2>

## 6.4 接口技巧

<h2 id="e51931e1b204267a72cf1f21b8185b9c"></h2>

##### 让编译器检查，确保某个类型实现 接口

```go
// 确保 *Data 实现 fmt.Stringer 接口
var _ fmt.Stringer = (*Data)(nil)
```

---

<h2 id="271a34f0fb6d790c4f6cb6f932c41bbe"></h2>

##### 让函数直接"实现"接口

```go
type Tester interface {  // Tester接口，需要实现 Do()方法
    Do()
}
type FuncDo func()      // 定义一个方法类型 FuncDo
func (self FuncDo) Do() { self() }  // FuncDo 实现了 Tester 接口

func main() {
    // 匿名方法转为 FuncDo 类型，再赋值给 Tester接口对象
    var t Tester = FuncDo(func() { println("Hello, World!") })
    t.Do()
}
```









