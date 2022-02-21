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
		result = math.Pi * math.Pow(sideLen, 2)
	} else if sidesNum == SidesTriangle {
		result = sideLen * sideLen * math.Sqrt(3) / 4
	} else if sidesNum == SidesSquare {
		result = math.Pow(sideLen, 2)

	}

	return result
}

