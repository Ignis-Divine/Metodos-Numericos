package main

/*
Escrito por Juan Salomón Sains Hdez
Ing en sistemas computacionales.
Metodos numericos
*/

//import importa los paquetes que se requieren para la ejecucion del programa
import (
	"fmt"
	"github.com/Ignis-Divine/Metodos-Numericos/puntoFijo"
)

//metodo main llama a las funciones principales
func main() {
	fmt.Println("f(x)= x^3+sen(x)")
	//biseccion.Raiz(-2, 2)
	puntoFijo.Iterar(-2, 1)
}
