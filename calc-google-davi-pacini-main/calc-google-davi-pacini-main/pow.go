package main

import "math"

type Pow struct {
	Operacao
}

func (m Pow) Calculate(y, x float64) float64 {
	return math.Pow(x, y)
}
