package main

import "fmt"

func main() {
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, i)
		fmt.Printf("Len: %d, Cap: %d\n", len(s), cap(s))
	}
}
