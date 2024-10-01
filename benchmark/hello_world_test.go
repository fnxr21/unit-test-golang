package unittest

import (
	"testing"
)

// ===================== table benchmark

// go test -v -run=NotMathUnitTest -bench=BenchmarkTable
func BenchmarkTable(b *testing.B) {

	Benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Ivan",
			request: "Ivan",
		},
		{
			name:    "Nur",
			request: "Nur",
		},
		{
			name:    "Ivan Nur",
			request: "Ivan Nur",
		},
	}

	for _, bencBenchmark := range Benchmarks {
		b.Run(bencBenchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bencBenchmark.request)
			}
		})
	}

}

// ===================== benchmark sub
// go test -v -run=NotMathUnitTest -bench=BenchmarkSub/Ivan
func BenchmarkSub(b *testing.B) {
	b.Run("Ivan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fandi")
		}
	})
	b.Run("Nur", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nur")
		}
	})
}

// ===================== benchmark
//ini kena semua dengan testing dan benchmark
//go test -v -bench=.

// ini kena  benchmark saja di
// go test -v -run=NotMathUnitTest -bench=.
// spesifik
// go test -v -run=NotMathUnitTest -bench=BenchmarkHelloWorldNur

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fandi")
	}
}

func BenchmarkHelloWorldNur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nur")
	}
}
