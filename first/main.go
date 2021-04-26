package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU() + 2)
	fmt.Println(0x96)
	var testString string
	fmt.Println(testString)

	test1 := true
	fmt.Printf("%T\n", test1) 
}

