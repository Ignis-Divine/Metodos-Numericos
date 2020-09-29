package newton

//import, importa las librerias necesarias para el programa
import (
	"fmt"
	"math"
)

// var x0, es la variable para el valor incial
var x0 float64

//funcion FX contiene la funcion original dada
func FX(x float64) float64 {
	return (math.Pow(x,3)+2*math.Pow(x,2))
}

//funcion F1X contiene la funcion en su primera derivada
func F1X(x float64) float64 {
	return (3*math.Pow(x, 2)+ 4*x)
}

//funcion F2X, contiene la segunda derivada de la funcion principal
func F2X(x float64) float64 {
	return (6*x+4)
}

//funcion ValidarConvergencia valida que haya convergencia entre las funciones
func ValidarConvergencia()  {
	var converge float64
	converge = 2
	sn := true
	for converge >=1 && sn == true{
		Repite:
		fmt.Println("Ingresa el valor inicial x0")
		fmt.Scanln(&x0)
		converge=(FX(x0)*F2X(x0))/math.Pow(F1X(x0),2)
		fmt.Println("converge =", converge)
		if converge >= 1{
			sn=false
			fmt.Println("No converge")
			goto Repite
		}
	}
}

//funcion Iterar rectifica donde se encuentran las raices
func Iterar()  {
	sn:= true
	if sn {
		var error float64
		error = 2
		xd := math.Pow(1*10, -12)
		for error >= xd{
			xa := x0
			x0 = x0 - (FX(x0)/F1X(x0))
			error = math.Abs(xa - x0)
			fmt.Println("error = ", error)
		}
		fmt.Println("la raiz es: ", x0)
	}
}