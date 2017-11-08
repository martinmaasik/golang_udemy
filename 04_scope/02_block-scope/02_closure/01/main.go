package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		x := 43
		fmt.Println(x)
	}
	fmt.Println(x)

}

//tegin seda n2idet veits paremaks 0, n2itamaks, et curlyde
//sees on x teistsugune kui v2ljas
