package main

import "testing"

func BenchmarkPoorlyAligned(b *testing.B) {
	for range b.N {
		var items = make([]PoorlyAligned, 10_000_000)
		for j := range items {
			items[j].count = int64(j)
		}
	}
}

func BenchmarkWellAligned(b *testing.B) {
	for range b.N {
		var items = make([]WellAligned, 10_000_000)
		for j := range items {
			items[j].count = int64(j)
		}
	}
}
