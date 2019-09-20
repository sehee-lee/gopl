package main

import (
	"fmt"
)

func main() {
	s := "/User/chad.lee"

	fmt.Printf("original string : %s\n", s)

	s = basename1(s)
	fmt.Printf("after basename1 : %s\n", s)

	s = basename2(s)
	fmt.Printf("after basename2 : %s\n", s)
}

func basename1(s string) string {
	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:] // 이때, s[i+1:] len보다 큰 index의 공간은 null이 됨. 즉, chad.lee\0에서 lee\0이 됨. 그냥(s[4]) 접근 불가
			break
		}
	}
	return s
}

func basename2(s string) string {
	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i] // 이때, s[i+1:] len보다 큰 index의 공간은 null이 됨. 즉, chad.lee\0에서 lee\0이 됨. 그냥(s[4]) 접근 불가
			break
		}
	}
	return s
}