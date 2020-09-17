package puntoFijo

//import importa todas las librerias que requiere el programa
import (
	"fmt"
	"math"
)

//funcion G1X reemplaza a x para el calculo de la funcion ya derivada.
func G1X(x float64) float64 {
	return (math.Cos(x)) / (3 * (math.Pow(math.Sin(x), 2/3)))
}

//funcion converge revisa si x es menor a 1 para saber si hay convergencia
//en la funcion.
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

//funcion Iterar aproxima el valor numérico de las raíces de una función polinomial
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
