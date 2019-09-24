package main

import (
	"fmt"
)

type IntSlice struct {
	ptr 		*int
	len, cap	int
}

func main() {
	var runes []rune
	var x, y []int
	var z []int

	for _, r := range "Hello~!!!" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)

	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("y => %d cap=%d \t%v\n", i, cap(y), y)
		x = y
	}

	z = append(z, 1)
	z = append(z, 2, 3)
	z = append(z, 4, 5, 6)
	z = append(z, z...)
	fmt.Println(z)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}