package main

import "math"

type Rot struct {
	Operacao
}

func (m Rot) Calculate(x, y float64) float64 {
	return math.Pow(x, 1/y)
}
