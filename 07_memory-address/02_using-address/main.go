package main

import "fmt"

const meetridJardideks float64 = 1.09361

func main() {
	var meetrid float64
	fmt.Print("Sisesta ujutud meetrite arv: ")
	fmt.Scan(&meetrid)
	jardid := meetrid * meetridJardideks
	fmt.Println(meetrid, " meetrit on", jardid, " jardi.")
}
