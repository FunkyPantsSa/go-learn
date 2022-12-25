package main

import (
	fm "fmt"
	"os"
)

var a string

func main() {
	fm.Println(os.Hostname())
	fm.Print(os.DirFS("C:/"))

}
