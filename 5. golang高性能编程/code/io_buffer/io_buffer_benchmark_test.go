package iobuffer

import (
	"bufio"
	"io"
	"os"
	"testing"
)

const N = 10000

func writeNotBuffered(w io.Writer, count int) {
	for i := 0; i < count; i++ {
		w.Write([]byte(":val\n"))
	}
}

func writeBuffered(w io.Writer, count int) {
	buf := bufio.NewWriterSize(w, 16*1024)
	for i := 0; i < count; i++ {
		buf.Write([]byte(":val\n"))
	}
	buf.Flush()
}

func BenchmarkWriteNotBuffered(b *testing.B) {
	for b.Loop() {
		f, _ := os.CreateTemp("", "nobuf")
		writeNotBuffered(f, N)
		f.Close()
		os.Remove(f.Name())
	}
}

func BenchmarkWriteBuffered(b *testing.B) {
	for b.Loop() {
		f, _ := os.CreateTemp("", "buf")
		writeBuffered(f, N)
		f.Close()
		os.Remove(f.Name())
	}
}
