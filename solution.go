package square


import (
	"math"
)

type sidesNumber int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum sidesNumber) float64 {

	var result float64 = 0.0

	if sidesNum == SidesCircle {
		result = math.Pi * sideLen * sideLen
	} else if sidesNum == SidesTriangle {
		result = sideLen * math.Sqrt(sideLen) / 4
	} else if sidesNum == SidesSquare {
		result = sideLen * sideLen
	}

	return result
}

