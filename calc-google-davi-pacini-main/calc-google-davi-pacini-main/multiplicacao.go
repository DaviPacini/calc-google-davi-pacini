package main

type Mult struct {
	Operacao
}

func (m Mult) Calculate(x, y float64) float64 {
	return x * y
}
