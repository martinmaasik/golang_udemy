package main

import "fmt"

func main() {
	x := 15
	kahandus := func() int {
		x = x / 2
		return x * 3 / 5
	}
	fmt.Println(kahandus())
	fmt.Println(kahandus())
}

// proovisin siin ++ ja ka --,
// lopuks panin, et jagaks kahega
// pane tahele et vastus tuleb t2isarv (ja on ka t2isarv systeemisiseselt)
// compiler j2tab j22gi lihtsalt k6rvale
