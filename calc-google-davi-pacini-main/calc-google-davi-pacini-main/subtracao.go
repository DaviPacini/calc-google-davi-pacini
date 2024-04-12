package main

type Sub struct {
	Operacao
}

func (m Sub) Calculate(x, y float64) float64 {
	return x - y
}
