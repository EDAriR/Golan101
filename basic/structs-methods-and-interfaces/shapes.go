package main

import "math"

/*
这种定义 interface 的方式与大部分其他编程语言不同。通常接口定义需要这样的代码 My type Foo implements interface Bar。
但是在我们的例子里，
Rectangle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
Circle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
string 没有这种方法，所以它不满足这个接口
等等
在 Go 语言中 interface resolution 是隐式的。如果传入的类型匹配接口需要的，则编译正确。
*/
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}
