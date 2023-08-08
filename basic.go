package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 33
var ss = "kkk"
var bb = true

var (
	v  = 1
	b  = 2
	kk = "str" +
		"ss" +
		"ss"
)

func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}

func varableInitialValue2() {
	var a, b = 4, 5
	var s = "qwer"
	fmt.Println(a, b, s)
}

func varabletypeDeduction() {
	var a, b, d, f = 1, 2, 3, true
	b = 6
	fmt.Println(a, b, d, f)

}

func variableStorter() {
	a, b, c := 1, "abv", true
	fmt.Println(a, b, c)
}

func eular() {
	//complex64
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

}

func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	// c = math.Sqrt(a*a + b*b)
	fmt.Println(c)
	//强制显示转换

}

func consts() {
	const (
		name = string("xiaoma")
		age  = int(18)
		a, b = 3, 4
		//可以做各个类型用
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(name, age, c)

}

func enums() {
	const (
		cpp = iota
		java
		python
		c
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println(cpp, java, python, c)

}
func main() {
	variableZeroValue()
	fmt.Println("hello word")
	variableStorter()
	varabletypeDeduction()
	varableInitialValue2()
	fmt.Println(aa, ss, kk, bb, v)
	eular()
	triangle()
	consts()
	enums()
}
