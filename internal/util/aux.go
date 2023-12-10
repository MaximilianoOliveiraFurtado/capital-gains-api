package util

import "math"

func RoundTo2Decimals(value float64) float64 {
	return math.Round(value*100) / 100
}
