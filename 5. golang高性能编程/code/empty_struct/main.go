package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a, b := struct{}{}, struct{}{}
	fmt.Printf("address: %p sizeof: %d\n", &a, unsafe.Sizeof(a))
	fmt.Printf("address: %p sizeof: %d\n", &b, unsafe.Sizeof(b))
}
