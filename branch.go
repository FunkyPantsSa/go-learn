package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {

	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score %d ", score))

	case score < 60:
		g = "F"
	case score < 70:
		g = "E"
	case score < 80:
		g = "D"
	case score < 90:
		g = "C"
	case score < 100:
		g = "B"
	case score == 100:
		g = "A"
		//default:
		//
		//	panic(fmt.Printf("wrong score %d", score))
	}
	return g

}

func main() {
	const (
		filename = "abc.txt"
	)

	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("%s\n", contents)
	}
	// fmt.Println(contents)  err

	fmt.Println(
		grade(67),
		grade(101),
	)

}
