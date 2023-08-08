package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	//整数转二进制

	result := ""
	//每次使用n/2当做循环的条件
	for ; n > 0; n /= 2 {
		//n/2的余数作为二进制的数字
		lsb := n % 2
		//将每一位余数加到前一位的前面。就是整数转二进制
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)

	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func forver() {
	for {
		fmt.Println("abc")
	}

}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(8892711),
		convertToBin(0))
	printFile("abc.txt")
	//forver()
}
