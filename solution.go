package square

import (
	"math"
)

type SidesNum int

const (
	SidesCircle   SidesNum = 0
	SidesTriangle SidesNum = 3
	SidesSquare   SidesNum = 4
)

func CalcSquare(sideLen float64, sidesNum SidesNum) float64 {
	if sideLen <= 0.0 {
		return 0.0
	}

	switch sidesNum {
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	case SidesTriangle:
		return sideLen * sideLen * math.Sqrt(3.0) / 4.0
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0.0
	}
}
