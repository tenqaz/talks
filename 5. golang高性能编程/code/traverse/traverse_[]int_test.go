package demo1

import "testing"

func BenchmarkSliceIntIndexLoop(b *testing.B) {
	slice := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(slice); j++ {
			_ = slice[j]
		}
	}
}

func BenchmarkSliceIntRangeValue(b *testing.B) {
	slice := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		for _, v := range slice {
			_ = v
		}
	}
}
