package main
import (
	"fmt"
	"os"
)
func main() {
	s, sep := "", "" //:=는 할당 연산자
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}