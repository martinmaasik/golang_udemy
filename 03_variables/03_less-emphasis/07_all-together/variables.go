package main

import "fmt"

var a = "seda talletatakse muutujas a"
var b, c string = "muutujas b", "muutujas c"
var d string

func main() {

	d = "muutujas d"
	var e = 42 // see enam ei ole paketisisene
	f := 43
	g := "muutujas g"
	h, i := "muutujas h", "muutujas i"
	j, k, l, m := 44.7, true, false, 'm'
	n := "n"
	o := `o`

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)

	// kui deklaratsioon on paketi levelil, aga vaartuse andmine
	// funktsiooni levelil, siis scope jaab ikka paketi tasandile
}
