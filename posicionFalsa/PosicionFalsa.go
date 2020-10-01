package posicionFalsa

import (
	"fmt"
	"math"
)

var a, b float64

func funcion(x float64) float64 {
	return (math.Pow(x,3)+2*math.Pow(x, 2)+10*x-20)
}

func pideDatos() (float64, float64) {
	fmt.Println("Ingresa el valor de a")
	fmt.Scanln(&a)
	fmt.Println("ingresa el valor de b")
	fmt.Scanln(&b)
	return a, b
}

func Iterar()  {
	var error float64
	REPITE:
	a, b = pideDatos()
	fa := funcion(a)
	fmt.Println("f(a): ",fa)
	fb := funcion(b)
	fmt.Println("f(b): ",fb)
	fafb := fa * fb
	fmt.Println("f(a)*f(b): ",fafb)
	if fafb <= 0{
		fmt.Println("converge")
	}else{
		fmt.Println("diverge")
		goto REPITE
	}
	c:= ((a*fb-b*fa)/(fb-fa))
	fmt.Println("c: ",c)
	fc:= funcion(c)
	fmt.Println("f(c): ",fc)
	error = 1
	c = ((a*fb-b*fa)/(fb-fa))+100
	for error > (math.Pow(1*10, -12)){
		ca := c
		c= ((a*fb-b*fa)/(fb-fa))
		if (fa*fc) > 0 {
			a = c
		}
		if (fb*fc) > 0{
			b = c
		}
		error = math.Abs(ca-c)
	}
	fmt.Println("la raiz es: ", c)
	fmt.Println("el error es: ", error)
}

