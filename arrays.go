package main

import (
	"fmt"
)

func printArrys(arr [5]int) {
	for i := range arr {
		arr[0] = 100
		fmt.Println(arr[i])
	}

}
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr3)
	fmt.Println(arr2)
	fmt.Println(arr1)

	var grade [4][5]bool
	fmt.Println(grade)

	//常用遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//range遍历
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	//获取下标和遍历
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	//只要遍历不要下标
	for _, v := range arr3 {
		fmt.Println(v)
	}

	//遍历获得最大的数和下标
	maxi := -1
	maxValue := -1
	for i, v := range arr3 {
		if v > maxValue {
			maxi, maxValue = i, v
		}

	}
	fmt.Println(maxi, maxValue)

	//数组求和
	sum := 0
	for _, v := range arr3 {
		sum += v

	}

	//printArrys(arr2)    //err:Cannot use 'arr2' (type [3]int) as the type [5]int
	fmt.Println("printArrys(arr1)")
	//数组是值类型，不同长度的数组被认为是不一样的
	printArrys(arr1)
	//调用数组会拷贝数组，可以用指针修改
	fmt.Println("arr1")
	fmt.Println(arr1)
}
