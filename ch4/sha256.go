package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	zero(&c2)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	c3 := sha512.Sum512([]byte("x"))
	c4 := sha512.Sum512([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c3, c4, c3 == c4, c3)

}

func zero (ptr *[32]byte) {
	/* same work
	for i := range ptr {
		ptr[i] = 0
	} */

	*ptr = [32]byte{}
}