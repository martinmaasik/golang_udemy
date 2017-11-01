package main

import "fmt"

func main() {
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%b \t %d \t %x \n", i, i, i)
	}
}

// \t lisab tabulaatori tyhiku, samal ajal kui \n viskab uuele reale
