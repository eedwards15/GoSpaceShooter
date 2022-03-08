package helpers

import "math"

func DistanceBetween(x1, y1, x2, y2 float64) float64 {
	y := math.Pow(y2-y1, 2.0)
	x := math.Pow(x2-x1, 2)

	results := math.Sqrt(y + x)
	return results
}
