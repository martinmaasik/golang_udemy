package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10) //1*2 astmel 10
	MB = 1 << (iota * 10) //1*2 astmel 20
	GB = 1 << (iota * 10) //1*2 astmel 30
	TB = 1 << (iota * 10) //1*2 astmel 40
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}

//n << x is "n times 2, x times". And y >> z is "y divided by 2, z times"
