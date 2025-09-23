package demo1

import "testing"

type Item struct {
	ID   int
	Data [4096]byte // 增加数据量以放大性能差异
}

func BenchmarkStructIndex(b *testing.B) {
	var slice [1000]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for j := 0; j < len(slice); j++ {
			tmp = slice[j].ID
		}
		_ = tmp
	}
}

func BenchmarkStructRangeValue(b *testing.B) {
	var slice [1000]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, v := range slice {
			tmp = v.ID
		}
		_ = tmp
	}
}

func BenchmarkStructRangeIndexValue(b *testing.B) {
	var slice [1000]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for j := range slice {
			tmp = slice[j].ID
		}
		_ = tmp
	}
}
