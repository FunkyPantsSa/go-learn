package cpuinfo

import (
	"fmt"
	"wmi"
)

type cpuInfo struct {
	Name          string
	NumberOfCores uint32
	ThreadCount   uint32
}

func main() {

	var cpuinfo []cpuInfo

	err := wmi.Query("Get-Process", &cpuinfo)
	if err != nil {
		return
	}
	fmt.Printf("Cpu info =", cpuinfo)
}
