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

// GeneratePointRejection generates a point by picking a random point in cartesian coordinates.
// If the point generated falls outside of the bounds of the circle, then it generates a new one.
// It will continue attempting until a working point is found.
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

// TODO: Doc comment
func GeneratePointSqrt() Point {
	radiusLength := math.Sqrt(rand.Float64())
	radianAngle := rand.Float64() * (2 * math.Pi)
	sin, cos := math.Sincos(radianAngle)
	return Point{
		X: radiusLength * sin,
		Y: radiusLength * cos,
	}
}

// TODO: Doc comment
func GeneratePointTriangle() Point {
	radiusLength := rand.Float64() + rand.Float64()
	if radiusLength > 1 {
		radiusLength = 2 - radiusLength
	}
	radianAngle := rand.Float64() * (2 * math.Pi)
	sin, cos := math.Sincos(radianAngle)
	return Point{
		X: radiusLength * sin,
		Y: radiusLength * cos,
	}
}

// TODO: Doc comment
func GeneratePointMax() Point {
	r, r2 := rand.Float64(), rand.Float64()
	if r2 > r {
		r = r2
	}
	radianAngle := rand.Float64() * (2 * math.Pi)
	sin, cos := math.Sincos(radianAngle)
	return Point{
		X: r * sin,
		Y: r * cos,
	}
}
