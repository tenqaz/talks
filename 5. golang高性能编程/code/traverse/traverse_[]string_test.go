package demo1

import "testing"

func BenchmarkSliceStringIndexLoop(b *testing.B) {
	slice := make([]string, 1000)
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(slice); j++ {
			_ = slice[j]
		}
	}
}

func BenchmarkSliceStringRangeValue(b *testing.B) {
	slice := make([]string, 1000)
	for i := 0; i < b.N; i++ {
		for _, v := range slice {
			_ = v
		}
	}
}
