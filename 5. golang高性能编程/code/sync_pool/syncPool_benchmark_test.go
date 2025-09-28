package main

import (
	"runtime"
	"sync"
	"testing"
)

type Data struct {
	Values [1024]int
}

func BenchmarkWithoutPooling(b *testing.B) {
	var dummy *Data
	for range b.N {
		data := &Data{}
		data.Values[0] = 42
		dummy = data
	}
	runtime.KeepAlive(dummy)
}

var dataPool = sync.Pool{
	New: func() any {
		return &Data{}
	},
}

func BenchmarkWithPooling(b *testing.B) {
	var dummy *Data
	for range b.N {
		obj := dataPool.Get().(*Data)
		obj.Values[0] = 42
		dummy = obj
		dataPool.Put(obj)
	}
	runtime.KeepAlive(dummy)
}
