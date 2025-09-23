package main

import (
	"fmt"
	"unsafe"
)

type PoorlyAligned struct {
	flag  bool
	count int64
	id    byte
}

type WellAligned struct {
	count int64
	flag  bool
	id    byte
}

func main() {
	fmt.Println(unsafe.Sizeof(PoorlyAligned{}))
	fmt.Println(unsafe.Sizeof(WellAligned{}))
}
