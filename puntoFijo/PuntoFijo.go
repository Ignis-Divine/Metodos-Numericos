package puntoFijo

import (
	"fmt"
	"math"
)

func G1X(x float64) float64 {
	return (math.Cos(x)) / (3 * (math.Pow(math.Sin(x), 2/3)))
}

func converge(x float64) bool {
	if x < 1 {
		return true
	}
	return false
}

//funcion calculaC calcula el valor de la variable c
func calcularC(a float64, b float64) float64 {
	return (a + b) / 2
}

func Iterar(a float64, b float64) {
	fmt.Println("Limite inferior del intervalo", a)
	fmt.Println("Limite superior del intervalo", b)
	hayRaiz := converge(calcularC(a, b))
	fmt.Println(hayRaiz)
	if hayRaiz {
		x := calcularC(a, b)
		fmt.Println("x = ", x)
		xa := x + 100
		fmt.Println("xa = ", xa)
		for math.Abs(xa-x) > 1*10^13 {
			xa = x
			x = G1X(x)
		}
		fmt.Println("la raiz esta en", x)
	}

}
