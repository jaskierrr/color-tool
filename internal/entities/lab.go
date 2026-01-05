package entities

// LabPoint is pixel in LAB color space
type LabPoint struct {
	L, A, B float64
}

func ToLabPoint(l, a, b float64) LabPoint {
	return LabPoint{L: l, A: a, B: b}
}
