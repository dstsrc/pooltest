package pooltest

import (
	"fmt"
	"sync"
	"testing"
)

// go test . -bench=Benchmark_SC1 --race
func Benchmark_SC1(b *testing.B) {
	goroutines := []int{1, 10, 100}
	for _, j := range goroutines {
		b.Run(fmt.Sprintf("%d goroutines", j), func(b *testing.B) {
			p := New()
			benchSC(b, p.sc1, j)
		})
	}
}

// go test . -bench=Benchmark_SC2 --race
func Benchmark_SC2(b *testing.B) {
	goroutines := []int{1, 10, 100}
	for _, j := range goroutines {
		b.Run(fmt.Sprintf("%d goroutines", j), func(b *testing.B) {
			p := New()
			benchSC(b, p.sc2, j)
		})
	}
}

// go test . -bench=Benchmark_SC3 --race -benchmem
func Benchmark_SC3(b *testing.B) {
	goroutines := []int{1, 10, 100}
	for _, j := range goroutines {
		b.Run(fmt.Sprintf("%d goroutines", j), func(b *testing.B) {
			p := New()
			benchSC(b, p.sc3, j)
		})
	}
}

func benchSC(b *testing.B, scFunc func(), goroutines int) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg := &sync.WaitGroup{}
		wg.Add(goroutines)
		for j := 0; j < goroutines; j++ {
			go sc(scFunc, wg)
		}
		wg.Wait()
	}
}

func sc(scFunc func(), wg *sync.WaitGroup) {
	scFunc()
	wg.Done()
}
