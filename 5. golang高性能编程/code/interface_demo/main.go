package main

import "fmt"

// func basic_demo() {
// 	var demo interface{}
// 	demo = 1 // 发生boxing，整型1逃逸到堆
// 	fmt.Println(demo)
// }

// func basic_demo2() {
// 	var demo int
// 	demo = 1
// 	fmt.Println(demo)
// }

type Person struct {
	Name string
}

func struct_demo() {
	var demo interface{}
	demo = Person{}
	fmt.Println(demo)
}

func point_demo() {
	var demo interface{}
	demo = &Person{Name: "zhangsan"}
	demo = &demo
	_ = demo
	// fmt.Println(demo)
}

func main() {
	// basic_demo()
	// basic_demo2()
	// struct_demo()
	// point_demo()
}
