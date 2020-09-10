package biseccion

//import importa la librebria de matematicas y para formato de texto basico
import (
	"fmt"
	"math"
)

//funcion exponente eleva x a un exponente fijo
func exponente(x float64) float64 {
	return math.Pow(x, 3)
}

//funcion seno calcula el seno de x
func seno(x float64) float64 {
	return math.Sin(x)
}

//funcion FA calcula la funcion dada remplazando x con el limite inferior
func FA(a float64) float64 {
	parteA := exponente(a)
	parteB := seno(a)
	return parteA+parteB
}

//funcion FB calcula la funcion dada remplazando x con el limite superior
func FB(b float64) float64 {
	parteA := exponente(b)
	parteB := seno(b)
	return parteA+parteB
}

//funcion calculaC calcula el valor de la variable c
func calcularC(a float64, b float64) float64 {
	return (a+b)/2
}

//funcion FC calcula la funcion dada remplazando x con la variable c
func FC(a float64, b float64) float64 {
	c:= calcularC(a, b)
	parteA := exponente(c)
	parteB := seno(c)
	return parteA+parteB
}

//funcion FAFB calcula la multiplicacion de F(A)*F(B)
func FAFB(a float64, b float64) float64 {
	return FA(a)*FB(b)
}

//funcion FAFC calcula la multiplicacion de F(A)*F(C)
func FAFC(a float64, b float64) float64 {
	return FA(a)*FC(a, b)
}

//funcion FBFC calcula la multiplicacion de F(B)*F(C)
func FBFC(a float64, b float64) float64 {
	return FB(b)*FC(a ,b)
}

//funcion Raiz aproxima el valor numérico de las raíces de una función polinomial
func Raiz(a float64, b float64)  {
	fmt.Println("Limite inferior del intervalo", a)
	fmt.Println("Limite superior del intervalo", b)
	if FAFB(a, b) >=0{
		fmt.Println("No hay raiz en el intervalo",a, b)
		return
	}
	var forzar float64
	if forzar > (10^-12) {
		forzar=FC(a, b)
		forzar= math.Abs(forzar)
		if FAFC(a, b) < 0{
			b = calcularC(a, b)
			return
		}
		if FBFC(a, b) < 0 {
			a = calcularC(a, b)
			return
		}
		fmt.Println("la raiz de x =", calcularC(a, b))
		fmt.Println("f(c) = ", forzar)
	}
}