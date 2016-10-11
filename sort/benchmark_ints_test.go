package sort

import (
	"sort"
	"testing"

	"github.com/seifer/go-dsa/sort/insertionsort"
	"github.com/seifer/go-dsa/sort/mergesort"
	"github.com/seifer/go-dsa/sort/quicksort"
	"github.com/seifer/go-dsa/sort/selectionsort"
	"github.com/seifer/go-dsa/sort/shellsort"
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

func BenchmarkInt1KInsertionSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		insertionsort.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkInt1KQuickSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		quicksort.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkInt1KShellSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		shellsort.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkInt1KMergeSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		mergesort.Ints(data)
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

func BenchmarkInt64KInsertionSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		insertionsort.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkInt64KQuickSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		quicksort.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkInt64KShellSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		shellsort.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkInt64KMergeSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		mergesort.Ints(data)
		b.StopTimer()
	}
}
