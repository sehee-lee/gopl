package main

import "fmt"

func main() {
	const (
		deadbeef = 0xdeadbeef
		a = uint32(deadbeef)
		b = float32(deadbeef)
		c = float64(deadbeef)
		f = uint(deadbeef)
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(f)


}