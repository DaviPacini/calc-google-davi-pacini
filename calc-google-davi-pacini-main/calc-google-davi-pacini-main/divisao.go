package main

type Div struct {
	Operacao
}

func (m Div) Calculate(x, y float64) float64 {
	return x / y
}
