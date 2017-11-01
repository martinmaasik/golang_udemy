package main

import "fmt"

func main() {
	for i := 65; i < 123; i++ {
		fmt.Printf("%b \t %d \t %x \t %q \n", i, i, i, i)
	}
}

// %d on siis decimal
// %b binaarne
// %x on hexadecimal
// ja #q on see utf8 mingi rahvusvaheliselt tunnustatud huinjaa, kus m2rgid 65 kuni 123 on inglise t2hestik
