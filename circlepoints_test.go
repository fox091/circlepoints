package circlepoints

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func TestInnerCircleDistribution(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	area := math.Pi
	halfArea := area / 2
	halfAreaInnerCircleRadius := math.Sqrt(halfArea / math.Pi)
	count := 100000000

	generators := []struct {
		name      string
		generator PointGenerator
	}{
		{
			name:      "GeneratePointMax",
			generator: GeneratePointMax,
		},
		{
			name:      "GeneratePointRejection",
			generator: GeneratePointRejection,
		},
		{
			name:      "GeneratePointSqrt",
			generator: GeneratePointSqrt,
		},
		{
			name:      "GeneratePointTriangle",
			generator: GeneratePointTriangle,
		},
	}

	for _, g := range generators {
		t.Run(g.name, func(t *testing.T) {
			points := make([]Point, count)
			for i := 0; i < count; i++ {
				points[i] = g.generator()
			}
			countInsideInner := 0
			countOutsideInner := 0
			for _, point := range points {
				hypotenuse := math.Sqrt(point.X*point.X + point.Y*point.Y)
				if hypotenuse <= halfAreaInnerCircleRadius {
					countInsideInner++
				} else {
					countOutsideInner++
				}
			}
			assert.GreaterOrEqual(t, float64(countInsideInner)/float64(countOutsideInner), .99)
			assert.LessOrEqual(t, float64(countInsideInner)/float64(countOutsideInner), 1.01)
		})
	}
}

func TestSemiCircleDistribution(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	count := 100000000

	generators := []struct {
		name      string
		generator PointGenerator
	}{
		{
			name:      "GeneratePointMax",
			generator: GeneratePointMax,
		},
		{
			name:      "GeneratePointRejection",
			generator: GeneratePointRejection,
		},
		{
			name:      "GeneratePointSqrt",
			generator: GeneratePointSqrt,
		},
		{
			name:      "GeneratePointTriangle",
			generator: GeneratePointTriangle,
		},
	}

	for _, g := range generators {
		t.Run(g.name, func(t *testing.T) {
			points := make([]Point, count)
			for i := 0; i < count; i++ {
				points[i] = g.generator()
			}
			countLeft := 0
			countRight := 0
			countUp := 0
			countDown := 0
			for _, point := range points {
				if point.X < 0 {
					countLeft++
				}
				if point.X > 0 {
					countRight++
				}
				if point.Y < 0 {
					countDown++
				}
				if point.Y > 0 {
					countUp++
				}
			}
			assert.GreaterOrEqual(t, float64(countLeft)/float64(countRight), .99)
			assert.LessOrEqual(t, float64(countLeft)/float64(countRight), 1.01)
			assert.GreaterOrEqual(t, float64(countUp)/float64(countDown), .99)
			assert.LessOrEqual(t, float64(countUp)/float64(countDown), 1.01)
		})
	}
}
