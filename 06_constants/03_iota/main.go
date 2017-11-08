package main

import "fmt"

const (
	a = iota
	b = iota
	c = 2 * iota
	d = iota / 3
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// iota t6stab yhe yhiku haaval kogu aeg
