package main

import (
	"fmt"
)

type Currency int
const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}

	fmt.Println(a == b, a == c, b == c)

	//d := [3]int{1, 2}
	//fmt.Println(a == d)

	symbol := [...]string{USD: "$", RMB: "¥", EUR: "€", GBP: "£"}

	fmt.Println(RMB, symbol[RMB])
	fmt.Println(symbol)

}
