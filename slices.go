package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:6]
	fmt.Println("arr 2-6", s)
	fmt.Println("arr ")
}
