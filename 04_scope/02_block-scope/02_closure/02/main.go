package main

import "fmt"

var x = 0

func lisandumine() int {
	x++
	return x
}

func main() {
	fmt.Println(lisandumine())
	fmt.Println(lisandumine())
}
