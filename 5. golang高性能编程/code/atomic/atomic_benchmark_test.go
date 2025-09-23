package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicIncrement(b *testing.B) {
	var counter atomic.Int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Add(1)
		}
	})
}

func BenchmarkMutexIncrement(b *testing.B) {
	var (
		counter int64
		mu      sync.Mutex
	)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	})
}
