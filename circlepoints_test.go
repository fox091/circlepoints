package circlepoints

import "testing"

func BenchmarkGeneratePointSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePointSqrt()
	}
}

func BenchmarkGeneratePointRejection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePointRejection()
	}
}

func BenchmarkGeneratePointTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePointTriangle()
	}
}

func BenchmarkGeneratePointMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePointMax()
	}
}

// TODO: Tests
