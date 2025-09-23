package main

import "testing"

type Worker interface {
	Work()
}

type LargeJob struct {
	payload [4096]byte
}

func (LargeJob) Work() {}

func BenchmarkBoxedLargeSlice(b *testing.B) {
	jobs := make([]Worker, 0, 1000)
	for range b.N {
		jobs = jobs[:0]
		for j := 0; j < 1000; j++ {
			var job LargeJob
			jobs = append(jobs, job)
		}
	}
}

func BenchmarkPointerLargeSlice(b *testing.B) {
	jobs := make([]Worker, 0, 1000)
	for range b.N {
		jobs := jobs[:0]
		for j := 0; j < 1000; j++ {
			job := &LargeJob{}
			jobs = append(jobs, job)
		}
	}
}

var sink Worker

func call(w Worker) {
	sink = w
}

func BenchmarkCallWithValue(b *testing.B) {
	for range b.N {
		var j LargeJob
		call(j)
	}
}

func BenchmarkCallWithPointer(b *testing.B) {
	for range b.N {
		j := &LargeJob{}
		call(j)
	}
}
