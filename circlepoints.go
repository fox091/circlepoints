package circlepoints

import (
	"math"
	"math/rand"
)

type Point struct {
	X, Y float64
}

type PointGenerator func() Point

type GenerationMethod int

const (
	Rejection GenerationMethod = iota
	SquareRoot
	Triangle
	Max
)

func GeneratePoints(numberToGenerate int, generationMethod GenerationMethod) []Point {
	var generator PointGenerator
	switch generationMethod {
	case Rejection:
		generator = GeneratePointRejection
	case SquareRoot:
		generator = GeneratePointSqrt
	case Triangle:
		generator = GeneratePointTriangle
	case Max:
		generator = GeneratePointMax
	}

	points := make([]Point, numberToGenerate)

	for i := 0; i < numberToGenerate; i++ {
		points[i] = generator()
	}

	return points
}

// GeneratePointRejection generates a point
func GeneratePointRejection() Point {
	for {
		x := (rand.Float64() * 2) - 1
		y := (rand.Float64() * 2) - 1
		if math.Sqrt(x*x+y*y) > 1 {
			continue
		}
		return Point{
			X: x,
			Y: y,
		}
	}
}

func GeneratePointSqrt() Point {
	radiusLength := math.Sqrt(rand.Float64())
	radianAngle := rand.Float64() * (2 * math.Pi)
	return Point{
		X: radiusLength * math.Sin(radianAngle),
		Y: radiusLength * math.Cos(radianAngle),
	}
}

func GeneratePointTriangle() Point {
	radiusLength := rand.Float64() + rand.Float64()
	if radiusLength > 1 {
		radiusLength = 2 - radiusLength
	}
	radianAngle := rand.Float64() * (2 * math.Pi)
	return Point{
		X: radiusLength * math.Sin(radianAngle),
		Y: radiusLength * math.Cos(radianAngle),
	}
}

func GeneratePointMax() Point {
	r, r2 := rand.Float64(), rand.Float64()
	if r2 > r {
		r = r2
	}
	radianAngle := rand.Float64() * (2 * math.Pi)
	return Point{
		X: r * math.Sin(radianAngle),
		Y: r * math.Cos(radianAngle),
	}
}
