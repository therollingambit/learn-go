package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("css/main.css")
	fmt.Println(dir)
	fmt.Println(file)
}