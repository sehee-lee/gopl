package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	s := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Printf("revers array from %d to ", a)
	reverse(a[:])
	fmt.Printf("%d\n", a)

	fmt.Printf("rotate two element from %d to ",s)
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)

	reverse(nil)

	fmt.Println(equal(nil, nil))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) || len(x) == 0 || len(y) == 0 {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}