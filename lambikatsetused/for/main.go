package main

import "fmt"

func main() {

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 11; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
