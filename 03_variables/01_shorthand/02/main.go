package main

import "fmt"

func main() {

	a, b, c, d, e, f, g := 10, "golang", 4.17, true, "Hello", `Do you like my hat?`, 'M'
	fmt.Printf("%T \n%T \n%T \n%T \n%T \n%T \n%T \n", a, b, c, d, e, f, g)
}
