package sort

import (
	"sort"
	"strconv"
	"testing"

	"github.com/seifer/go-dsa/sort/heapsort"
	"github.com/seifer/go-dsa/sort/insertionsort"
	"github.com/seifer/go-dsa/sort/mergesort"
	"github.com/seifer/go-dsa/sort/quicksort"
	"github.com/seifer/go-dsa/sort/selectionsort"
	"github.com/seifer/go-dsa/sort/shellsort"
	"github.com/seifer/go-dsa/sort/timsort"
)

func BenchmarkString1KStdSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0x2cc)
		}
		b.StartTimer()
		sort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString1KSelectionSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0x2cc)
		}
		b.StartTimer()
		selectionsort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString1KInsertionSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0x2cc)
		}
		b.StartTimer()
		insertionsort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString1KQuickSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0x2cc)
		}
		b.StartTimer()
		quicksort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString1KShellSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0x2cc)
		}
		b.StartTimer()
		shellsort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString1KMergeSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0x2cc)
		}
		b.StartTimer()
		mergesort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString1KHeapSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0x2cc)
		}
		b.StartTimer()
		heapsort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString1KTimSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0x2cc)
		}
		b.StartTimer()
		timsort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString64KStdSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0xcccc)
		}
		b.StartTimer()
		sort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString64KSelectionSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0xcccc)
		}
		b.StartTimer()
		selectionsort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString64KInsertionSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0xcccc)
		}
		b.StartTimer()
		insertionsort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString64KQuickSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0xcccc)
		}
		b.StartTimer()
		quicksort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString64KShellSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0xcccc)
		}
		b.StartTimer()
		shellsort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString64KMergeSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0xcccc)
		}
		b.StartTimer()
		mergesort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString64KHeapSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0xcccc)
		}
		b.StartTimer()
		heapsort.Strings(data)
		b.StopTimer()
	}
}

func BenchmarkString64KTimSort(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]string, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = strconv.Itoa(i ^ 0xcccc)
		}
		b.StartTimer()
		timsort.Strings(data)
		b.StopTimer()
	}
}
