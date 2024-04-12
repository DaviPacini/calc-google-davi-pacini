package main

type Soma struct {
	Operacao
}

func (m Soma) Calculate(x, y float64) float64 {
	return x + y
}
