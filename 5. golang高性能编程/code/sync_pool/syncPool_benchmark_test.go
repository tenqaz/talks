package main

import (
	"sync"
	"testing"
)

type Item struct {
	A, B, C int
}

var itemPool = sync.Pool{
	New: func() interface{} {
		return &Item{}
	},
}

func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		item := &Item{A: 1, B: 2, C: 3}
		_ = item.A + item.B + item.C
	}
}

func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		item := itemPool.Get().(*Item)
		item.A, item.B, item.C = 1, 2, 3 // 重置
		_ = item.A + item.B + item.C
		itemPool.Put(item)
	}
}
