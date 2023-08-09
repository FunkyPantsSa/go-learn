package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//函数和多值返回

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		//div函数返回两个值，只取q
		q, _ := div(a, b)
		return q, nil
	default:
		//panic("unknown operation" + op)
		return 0, fmt.Errorf(
			"wrong operation %s", op)

	}

}

func div(a, b int) (q, r int) {
	//可以给返回值取名字，对于调用者没区别，只有提示作用。
	//q := a / b
	//r := a % b
	return a / b, a % b
}

// 把函数当做传入的参数，函数式编程
func apply(op func(float64, float64) float64, a, b float64) float64 {
	//获取函数名的指针
	p := reflect.ValueOf(op).Pointer()
	//获取函数名
	opName := runtime.FuncForPC(p)
	fmt.Printf("Calling function %s with arges"+"(%f,%f)", opName, a, b)

	return op(a, b)

}

func sum(number ...int) int {
	k := 0
	for i := range number {
		k += number[i]

	}
	return k
}

func swap(a, b *int) {
	*a, *b = *b, *a

}
func swap2(a, b int) (int, int) {
	return b, a

}
func main() {

	fmt.Println(eval(3, 4, "+"))
	fmt.Println(div(6, 2))
	q, r := div(5, 7)
	fmt.Println(q, r)
	//if result, err := eval(3, 4, "!"); err != nil {
	//	fmt.Println("error:", err)
	//
	//} else {
	//	fmt.Println(result)
	//}

	fmt.Println(apply(math.Pow, 3, 4))
	//函数简写
	fmt.Println(apply(func(a float64, b float64) float64 {
		return math.Pow(a, b)

	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	//变量交换
	a, b := 3, 4
	//指针交换
	swap(&a, &b)
	fmt.Println(a, b)
	//返回交换
	fmt.Println(swap2(3, 4))

}

//返回值写在最后
//没有默认参数，可选参数，只有可变参数列表
