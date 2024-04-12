package main

type Operacao struct {
	Op1      float64
	Op2      float64
	Operador string
}

type Mathematician interface {
	Calculate(x, y float64) float64
}
