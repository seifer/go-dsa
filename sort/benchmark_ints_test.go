package sort

import (
	"sort"
	"testing"

	"github.com/seifer/go-dsa/sort/selectionsort"
)

func BenchmarkInt1KStdSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		sort.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkInt1KSelectionSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		selectionsort.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkInt64KStdSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		sort.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkInt64KSelectionSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		selectionsort.Ints(data)
		b.StopTimer()
	}
}
