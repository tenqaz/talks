package main

func allocate() *int {
	x := 42
	return &x // x escapes to the heap
}

func main() {
	allocate()
}
