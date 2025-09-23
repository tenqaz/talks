package main

import "testing"

func BenchmarkBoxedNotInterface(b *testing.B) {
	jobs := make([]int, 0, 1000)
	for range b.N {
		jobs = jobs[:0]
		for j := 0; j < 1000; j++ {
			jobs = append(jobs, j)
		}
	}
}

func BenchmarkBoxedWithInterface(b *testing.B) {
	jobs := make([]interface{}, 0, 1000)
	for range b.N {
		jobs = jobs[:0]
		for j := 0; j < 1000; j++ {
			jobs = append(jobs, j)
		}
	}
}
